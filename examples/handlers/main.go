package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	"github.com/TrixiS/goram"
	"github.com/TrixiS/goram/examples/routes"
	"github.com/TrixiS/goram/handlers"
)

var (
	token       = flag.String("token", "", "")
	adminUserID = flag.Int64("admin-id", 0, "")
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

	router := routes.CreateRouter(*adminUserID)

	for updates := range updatesChan {
		go func(updates []goram.Update) {
			found, err := router.FeedUpdates(ctx, bot, updates, handlers.Data{})

			if found && err != nil {
				fmt.Println("handler error", err)
			}
		}(updates)
	}
}
