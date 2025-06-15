package filters

import (
	"context"

	"github.com/TrixiS/goram"
	"github.com/TrixiS/goram/handlers"
)

// Passes if update message text equals to any of the provided strings
func Text(texts ...string) handlers.Filter[*goram.Message] {
	return func(
		ctx context.Context,
		bot *goram.Bot,
		message *goram.Message,
		data handlers.Data,
	) (bool, error) {
		for _, text := range texts {
			if message.Text == text {
				return true, nil
			}
		}

		return false, nil
	}
}

func HasText(
	ctx context.Context,
	bot *goram.Bot,
	message *goram.Message,
	data handlers.Data,
) (bool, error) {
	return len(message.Text) > 0, nil
}

func Or[T any](filters ...handlers.Filter[T]) handlers.Filter[T] {
	return func(
		ctx context.Context,
		bot *goram.Bot,
		update T,
		data handlers.Data,
	) (bool, error) {
		for _, f := range filters {
			ok, err := f(ctx, bot, update, data)

			if err != nil {
				return false, err
			}

			if ok {
				return true, nil
			}
		}

		return false, nil
	}
}

func And[T any](filters ...handlers.Filter[T]) handlers.Filter[T] {
	return func(
		ctx context.Context,
		bot *goram.Bot,
		update T,
		data handlers.Data,
	) (bool, error) {
		for _, f := range filters {
			ok, err := f(ctx, bot, update, data)

			if err != nil {
				return false, err
			}

			if !ok {
				return false, nil
			}
		}

		return true, nil
	}
}
