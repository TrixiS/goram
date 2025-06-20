package routes

import (
	"context"

	"github.com/TrixiS/goram"
	"github.com/TrixiS/goram/examples/handlers/markups"
	"github.com/TrixiS/goram/handlers"
)

func Start(
	ctx context.Context,
	bot *goram.Bot,
	message *goram.Message,
	data handlers.Data,
) error {
	_, err := bot.SendMessage(ctx, &goram.SendMessageRequest{
		ChatID:      message.ChatID(),
		Text:        "Hello world!",
		ReplyMarkup: markups.Start,
	})

	return err
}

func HelloQuery(
	ctx context.Context,
	bot *goram.Bot,
	query *goram.CallbackQuery,
	data handlers.Data,
) error {
	_, err := bot.AnswerCallbackQuery(ctx, &goram.AnswerCallbackQueryRequest{
		CallbackQueryID: query.ID,
		Text:            "Hello",
		ShowAlert:       true,
	})

	return err
}

func WorldQuery(
	ctx context.Context,
	bot *goram.Bot,
	query *goram.CallbackQuery,
	data handlers.Data,
) error {
	_, err := bot.AnswerCallbackQuery(ctx, &goram.AnswerCallbackQueryRequest{
		CallbackQueryID: query.ID,
		Text:            "World",
		ShowAlert:       true,
	})

	return err
}
