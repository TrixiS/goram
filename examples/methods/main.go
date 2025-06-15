package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/TrixiS/goram"
	"github.com/TrixiS/goram/flood"
	"github.com/TrixiS/goram/keyboards"
)

var token = os.Getenv("BOT_TOKEN")

var markup = &goram.ReplyKeyboardMarkup{
	Keyboard: keyboards.NewBuilder[goram.KeyboardButton]().
		Add(goram.KeyboardButton{Text: "Hello world"}).
		Build(),
	ResizeKeyboard: true,
	IsPersistent:   true,
}

func main() {
	bot := goram.NewBot(goram.BotOptions{
		Token: token,
		FloodHandler: flood.NewCondHandler( // optional flood handler
			func(ctx context.Context, method string, request any, duration time.Duration) {
				fmt.Println("waiting for flood", method, duration)
			},
		),
	})

	ctx := context.Background()

	me, _ := bot.GetMe(ctx) // errors aren't handled in this example, but you should do it
	fmt.Println("me", me)

	message, _ := bot.SendMessage(ctx, &goram.SendMessageRequest{
		ChatID:      goram.ChatID{Username: "somecoolchannel"},
		Text:        "Hello world",
		ReplyMarkup: markup,
	})

	fmt.Println("sent", message.MessageID)

	// sending files
	// you can use goram.NameReader struct to send buffered files
	photoFile, _ := os.Open("./photo.jpeg")
	message, _ = bot.SendPhoto(ctx, &goram.SendPhotoRequest{
		ChatID: goram.ChatID{ID: -100123123123123},
		Photo:  goram.InputFile{Reader: photoFile},
	})

	// downloading files by file ids
	downloadedFile, _ := os.OpenFile("./downloaded.jpeg", os.O_CREATE|os.O_WRONLY, 0o660)
	bot.DownloadFile(ctx, message.Photo[0].FileID, downloadedFile) // takes io.Writer
}
