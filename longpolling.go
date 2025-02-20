package goram

import (
	"context"
	"time"
)

// See goram.LongPollUpdates()
type LongPollUpdatesOptions struct {
	RequestOptions GetUpdatesRequest // Initial getUpdates request options
	Cap            uint              // Optional. Updates channel capacity
	RetryInterval  time.Duration     // Optional. Sleep for this duration if an error happens on getUpdates request. Default is 0
	MaxErrors      uint              // Optional. Exit from polling loop after MaxErrors errors in a row. The returned channel gets closed too. Default is unlimited
}

// Polls updates via calling Bot.GetUpdates() in a loop.
// Streams []Update instead of just Update because it enables better media group handling.
func LongPollUpdates(
	ctx context.Context,
	bot *Bot,
	options *LongPollUpdatesOptions,
) chan []Update {
	c := make(chan []Update, options.Cap)

	go func() {
		var errCount uint

		for {
			updates, err := bot.GetUpdates(ctx, &options.RequestOptions)

			if err != nil {
				if options.MaxErrors > 0 {
					errCount++

					if errCount >= options.MaxErrors {
						break
					}
				}

				time.Sleep(options.RetryInterval)
				continue
			}

			errCount = 0

			if len(updates) == 0 {
				continue
			}

			options.RequestOptions.Offset = updates[len(updates)-1].UpdateId + 1
			c <- updates
		}

		close(c)
	}()

	return c
}
