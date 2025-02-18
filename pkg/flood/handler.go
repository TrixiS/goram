package flood

import (
	"time"
)

type Handler interface {
	Enter(
		method string,
		request any,
	) // Gets called BEFORE making an api request
	Handle(
		method string,
		request any,
		duration time.Duration,
	) // Gets called when 429 error happens
}

// Dummy flood handler. It will just sleep for flood duration.
type SimpleHandler struct {
	OnFlood func(method string, request any, duration time.Duration) // Optional callback that will be called before sleeping
}

func (s *SimpleHandler) Enter(string, any) {}

func (s *SimpleHandler) Handle(method string, request any, duration time.Duration) {
	if s.OnFlood != nil {
		s.OnFlood(method, request, duration)
	}

	time.Sleep(duration)
}

// TODO: conditional flood handler
