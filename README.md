# goram

Zero-dependency Telegram Bot API library for Go

```shell
go get -u github.com/TrixiS/goram@latest
```

See [Examples](https://github.com/TrixiS/goram/tree/master/examples) and [Reference](https://pkg.go.dev/github.com/TrixiS/goram)

## Features

- **No dependencies** - built purely around Go standart library
- **Simple** - library implementation is around 300 LOC. The rest is generated from <https://github.com/TBXark/telegram-bot-api-types>
- **`context.Context` support**
- Minimal runtime reflection overhead
- Type safety

## Calling methods

```Go
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
```

## Updates

```Go
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
```
