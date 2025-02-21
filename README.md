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
        FloodHandler: &flood.SleepHandler{ // Optional flood handler
            OnFlood: func(ctx context.Context, method string, request any, duration time.Duration) {
                fmt.Println("waiting for flood", method, duration)
            },
        },
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
        ChatId: goram.ChatId{ID: -100123123123123},
        Photo:  photoFile,
    })

    // downloading files by file ids
    downloadedFile, _ := os.OpenFile("./downloaded.jpeg", os.O_CREATE|os.O_WRONLY, 0o660)
    bot.DownloadFile(ctx, message.Photo[0].FileId, downloadedFile) // accepts io.Writer
}
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
    bot := goram.NewBot(goram.BotOptions{
        Token: "YOUR_TOKEN",
    })

    ctx := context.Background()

    updatesChan := goram.LongPollUpdates(ctx, bot, &goram.GetUpdatesRequest{
        Timeout:        10,
        AllowedUpdates: []string{"message"},
    })

    for updates := range updatesChan {
        for _, u := range updates {
            fmt.Println("new message", u.Message.Chat.Id, u.Message.Text)
        }
    }
}
```
