package main

import (
	"context"
	"fmt"
	"os"

	"github.com/TrixiS/goram"
)

var token = os.Getenv("BOT_TOKEN")

func main() {
	bot := goram.NewBot(goram.BotOptions{Token: token})

	updatesChan := goram.LongPollUpdates(context.Background(), bot, &goram.LongPollUpdatesOptions{
		RequestOptions: goram.GetUpdatesRequest{
			Limit:          100,
			Timeout:        10,
			AllowedUpdates: []goram.UpdateType{goram.UpdateMessage, goram.UpdateEditedMessage},
		},
		Cap: 100,
	})

	for updates := range updatesChan {
		for _, u := range updates {
			if u.Message != nil {
				fmt.Println("new message", u.Message.Chat.ID, u.Message.Text, u.Message.MessageID)
			} else if u.EditedMessage != nil {
				fmt.Println("edited message", u.EditedMessage.Chat.ID, u.EditedMessage.MessageID)
			}
		}
	}
}
