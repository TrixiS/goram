# goram

Zero-dependency Telegram Bot API library for Go

```shell
go get github.com/TrixiS/goram@latest
```

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
)

func main() {
    bot := goram.NewBot(goram.BotOptions{
        Token: "YOUR_TOKEN",
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
        ChatId: goram.ChatId{Username: "somecoolchannel"},
        Text:   "Hello world",
    })

    fmt.Println("sent", message.MessageId)

    // sending files
    // you can use goram.NameReader struct to send buffered files
    photoFile, _ := os.Open("./photo.jpeg")
    message, _ = bot.SendPhoto(ctx, &goram.SendPhotoRequest{
        ChatId: goram.ChatId{Id: -100123123123123},
        Photo:  goram.InputFile{Reader: photoFile},
    })

    // downloading files by file ids
    downloadedFile, _ := os.OpenFile("./downloaded.jpeg", os.O_CREATE|os.O_WRONLY, 0o660)
    bot.DownloadFile(ctx, message.Photo[0].FileId, downloadedFile) // takes io.Writer
}
```

## Markups

```Go
import "github.com/TrixiS/goram/keyboards"

bot.SendMessage(context.Background(), &goram.SendMessageRequest{
    ChatId: goram.ChatId{Id: 578371487},
    Text:   "Hello world",
    ReplyMarkup: &goram.ReplyKeyboardMarkup{
        Keyboard: keyboards.NewBuilder[goram.KeyboardButton]().
            Add(goram.KeyboardButton{Text: "Hello"}).
            Add(goram.KeyboardButton{Text: "World"}).
            Adjust(1).
            Build(),
        ResizeKeyboard: true,
    },
})
```

## Updates

```Go
package main

import (
    "context"
    "fmt"

    "github.com/TrixiS/goram"
)

func main() {
    bot := goram.NewBot(goram.BotOptions{Token: "YOUR_TOKEN"})

    updatesChan := goram.LongPollUpdates(context.Background(), bot, &goram.LongPollUpdatesOptions{
        RequestOptions: goram.GetUpdatesRequest{
            Limit:          100,
            Timeout:        10,
            AllowedUpdates: []goram.UpdateType{goram.UpdateMessage, goram.UpdateEditedMessage},
        },
    })

    for updates := range updatesChan {
        for _, u := range updates {
            if u.Message != nil {
                fmt.Println("new message", u.Message.Chat.Id, u.Message.Text, u.Message.MessageId)
            } else if u.EditedMessage != nil {
                fmt.Println("edited message", u.EditedMessage.Chat.Id, u.EditedMessage.MessageId)
            }
        }
    }
}
```
