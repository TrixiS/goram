package routes

import (
	"context"

	"github.com/TrixiS/goram"
	"github.com/TrixiS/goram/examples/markups"
	"github.com/TrixiS/goram/handlers"
)

func Start(
	ctx context.Context,
	bot *goram.Bot,
	message *goram.Message,
	data handlers.Data,
) error {
	return bot.SendMessageVoid(ctx, &goram.SendMessageRequest{
		ChatID:      message.ChatID(),
		Text:        "Hello world!",
		ReplyMarkup: markups.Start,
	})
}

func HelloQuery(
	ctx context.Context,
	bot *goram.Bot,
	query *goram.CallbackQuery,
	data handlers.Data,
) error {
	return bot.AnswerCallbackQueryVoid(ctx, &goram.AnswerCallbackQueryRequest{
		CallbackQueryID: query.ID,
		Text:            "Hello",
		ShowAlert:       true,
	})
}

func WorldQuery(
	ctx context.Context,
	bot *goram.Bot,
	query *goram.CallbackQuery,
	data handlers.Data,
) error {
	return bot.AnswerCallbackQueryVoid(ctx, &goram.AnswerCallbackQueryRequest{
		CallbackQueryID: query.ID,
		Text:            "World",
		ShowAlert:       true,
	})
}
