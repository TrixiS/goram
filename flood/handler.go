package flood

import (
	"context"
	"sync"
	"time"
)

type OnFloodFunc func(ctx context.Context, method string, request any, duration time.Duration)

// Interface for flood handlers. See `flood.CondHandler` or `flood.SleepHandler`. Or write one yourself.
type Handler interface {
	Enter(
		ctx context.Context,
		method string,
		request any,
	) // Gets called BEFORE making an api request
	Handle(
		ctx context.Context,
		method string,
		request any,
		duration time.Duration,
	) // Gets called when 429 error happens
}

// Dummy flood handler. It will just sleep for the flood duration.
//
// See `flood.CondHandler`.
type SleepHandler struct {
	OnFlood OnFloodFunc // Optional callback that will be called before sleeping
}

func (s *SleepHandler) Enter(context.Context, string, any) {}

func (s *SleepHandler) Handle(
	ctx context.Context,
	method string,
	request any,
	duration time.Duration,
) {
	if s.OnFlood != nil {
		s.OnFlood(ctx, method, request, duration)
	}

	time.Sleep(duration)
}

type methodState struct {
	mu    sync.Mutex
	cond  *sync.Cond
	flood bool
}

// This handler uses sync.Cond for flood handling.
// It has additional overhead of memory and locking.
// Flood is handled in separate for each api method.
type CondHandler struct {
	onFlood    OnFloodFunc
	stateMap   map[string]*methodState // key is api method name
	stateMapMu sync.RWMutex
}

func NewCondHandler(onFlood OnFloodFunc) *CondHandler {
	c := CondHandler{onFlood: onFlood, stateMap: make(map[string]*methodState)}
	return &c
}

func (c *CondHandler) Enter(ctx context.Context, method string, request any) {
	c.stateMapMu.RLock()
	state := c.stateMap[method]
	c.stateMapMu.RUnlock()

	if state == nil {
		return
	}

	state.mu.Lock()

	for state.flood {
		state.cond.Wait()
	}

	state.mu.Unlock()
}

func (c *CondHandler) Handle(
	ctx context.Context,
	method string,
	request any,
	duration time.Duration,
) {
	c.stateMapMu.Lock()
	state := c.stateMap[method]

	if state == nil {
		state = &methodState{}
		state.cond = sync.NewCond(&state.mu)
		c.stateMap[method] = state
	}

	c.stateMapMu.Unlock()

	state.mu.Lock()

	if state.flood {
		state.mu.Unlock()
		return
	}

	if c.onFlood != nil {
		c.onFlood(ctx, method, request, duration)
	}

	state.flood = true
	state.mu.Unlock()

	time.Sleep(duration)

	state.mu.Lock()
	state.flood = false
	state.mu.Unlock()

	state.cond.Broadcast()
}
