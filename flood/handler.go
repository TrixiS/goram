package flood

import (
	"context"
	"time"
)

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

// Dummy flood handler. It will just sleep for flood duration.
type SleepHandler struct {
	OnFlood func(
		ctx context.Context,
		method string,
		request any,
		duration time.Duration,
	) // Optional callback that will be called before sleeping
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

// TODO: conditional flood handler
