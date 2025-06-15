package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	"github.com/TrixiS/goram"
	"github.com/TrixiS/goram/cbdata"
	"github.com/TrixiS/goram/examples/handlers/markups"
	"github.com/TrixiS/goram/examples/handlers/routes"
	"github.com/TrixiS/goram/handlers"
)

var (
	token       = flag.String("token", "", "")
	adminUserId = flag.Int64("admin-id", 0, "")
)

func main() {
	flag.Parse()

	bot := goram.NewBot(goram.BotOptions{Token: *token})
	ctx := context.Background()

	updatesChan := goram.LongPollUpdates(ctx, bot, &goram.LongPollUpdatesOptions{
		RequestOptions: goram.GetUpdatesRequest{
			Timeout:        10,
			AllowedUpdates: []goram.UpdateType{goram.UpdateMessage, goram.UpdateCallbackQuery},
			Limit:          100,
		},
		Cap:           100,
		RetryInterval: time.Second * 1,
		MaxErrors:     3,
	})

	router := createRouter()

	for updates := range updatesChan {
		go func(updates []goram.Update) {
			found, err := router.FeedUpdates(ctx, bot, updates, handlers.Data{})

			if found && err != nil {
				fmt.Println("handler error", err)
			}
		}(updates)
	}
}

func createRouter() *handlers.Router {
	return handlers.
		NewRouter(handlers.RouterOptions{Name: "root"}).
		FilterMessage(
			func(ctx context.Context, bot *goram.Bot, message *goram.Message, data handlers.Data) (bool, error) {
				return message.From.ID == *adminUserId, nil
			},
		).
		FilterCallbackQuery(
			func(ctx context.Context, bot *goram.Bot, query *goram.CallbackQuery, data handlers.Data) (bool, error) {
				return query.From.ID == *adminUserId, nil
			},
		).
		Message(routes.Start, text("/start")).
		CallbackQuery(
			routes.HelloQuery,
			cbdata.FilterFunc(
				markups.Prefix,
				func(data markups.CbData) bool { return data.Type == markups.Hello },
			),
		).
		CallbackQuery(
			routes.WorldQuery,
			cbdata.FilterFunc(
				markups.Prefix,
				func(data markups.CbData) bool { return data.Type == markups.World },
			),
		)
}

func text(t string) handlers.Filter[*goram.Message] {
	return func(ctx context.Context, bot *goram.Bot, message *goram.Message, data handlers.Data) (bool, error) {
		return message.Text == t, nil
	}
}
