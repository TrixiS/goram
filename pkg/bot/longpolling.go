package bot

import (
	"context"

	"github.com/TrixiS/goram/pkg/types"
)

// Polls updates via calling Bot.GetUpdates() in a loop
//
// Request parameter is used for initial options of getUpdates requests.
// The returned channel never gets closed.
// This function does not panic and does not return errors encountered while making requests.
// If you need to handle/log those errors or set some retry policy, rewrite it yourself.
func LongPollUpdates(
	ctx context.Context,
	bot *Bot,
	request *types.GetUpdatesRequest,
) chan types.Update { // TODO: chan []types.Update for better media group handling?
	c := make(chan types.Update)

	go func() {
		for {
			updates, err := bot.GetUpdates(ctx, request)

			if err != nil || len(updates) == 0 {
				continue
			}

			request.Offset = updates[len(updates)-1].UpdateId + 1

			for _, u := range updates {
				c <- u
			}
		}
	}()

	return c
}
