package handlers

import (
	"context"

	"github.com/TrixiS/goram"
)

// Arbitrary data that is being passed to filters and handlers.
// See cbdata.Filter for an example.
type Data map[string]any

// Update handler function.
type Func[U any] func(ctx context.Context, bot *goram.Bot, update U, data Data) error

// Update handler filter function.
type Filter[U any] func(ctx context.Context, bot *goram.Bot, update U, data Data) bool

type handler[U any] struct {
	cb      Func[U]
	filters []Filter[U]
}

type RouterOptions struct {
	Name string // Name of the router. Useful for debugging.
}

type Router struct {
	Options  RouterOptions
	handlers handlers
	children []*Router
}

// Creates a new router.
//
// Routers can have top-level filters. You can add them using .Filter* methods.
// Top level filters get executed first.
//
// Also, routers can have any amount of update handlers.
// You can add them using .Message, .CallbackQuery, etc. methods.
//
// You can group routers using .Group() method.
func NewRouter(options RouterOptions) *Router {
	return &Router{
		Options: options,
	}
}

func (r *Router) Group(options RouterOptions) *Router {
	child := NewRouter(options)
	r.children = append(r.children, child)
	return child
}

// It is expected that you pass updates from goram.LongPollUpdates (for example) to the root router
// using .FeedUpdates() method.
// Updates fed to a router get passed through top-level filters first,
// then to router handlers and then to every child top-level filters and handlers (recoursively).
//
// If top-level filter returns false, then the update gets passed to the next router (if any).
// If handler filter returns false, then the update gets passed to the next handler (if any).
func (r *Router) FeedUpdates(
	ctx context.Context,
	bot *goram.Bot,
	updates []goram.Update,
	data Data,
) (bool, error) {
	// TODO: support for media group handlers

	for _, u := range updates {
		found, err := r.feedUpdate(ctx, bot, &u, data)

		if found {
			return found, err
		}
	}

	return false, nil
}

func callHandlers[T any](
	ctx context.Context,
	bot *goram.Bot,
	handlers []handler[T],
	update T,
	data Data,
) (bool, error) {
handlersLoop:
	for _, handler := range handlers {
		for _, filter := range handler.filters {
			if !filter(ctx, bot, update, data) {
				continue handlersLoop
			}
		}

		err := handler.cb(ctx, bot, update, data)
		return true, err
	}

	return false, nil
}
