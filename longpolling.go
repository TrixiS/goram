package goram

import (
	"context"
	"errors"
	"net/http"
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
//
// This function will panic on conflict (webhook set or another instance calling getUpdates).
func LongPollUpdates(
	ctx context.Context,
	bot *Bot,
	options *LongPollUpdatesOptions,
) chan []Update {
	c := make(chan []Update, options.Cap)

	go func() {
		errCount := uint(0)
		checkedForConflict := false

		for {
			updates, err := bot.GetUpdates(ctx, &options.RequestOptions)

			if err == nil {
				errCount = 0

				if len(updates) == 0 {
					continue
				}

				options.RequestOptions.Offset = updates[len(updates)-1].UpdateId + 1
				c <- updates
				continue
			}

			if errors.Is(err, context.Canceled) {
				break
			}

			if !checkedForConflict {
				if apiError, ok := err.(*APIError); ok &&
					apiError.ErrorCode == http.StatusConflict {

					panic(apiError)
				}

				checkedForConflict = true
			}

			if options.MaxErrors > 0 {
				errCount++

				if errCount >= options.MaxErrors {
					break
				}
			}

			time.Sleep(options.RetryInterval)
		}

		close(c)
	}()

	return c
}
