# goram

Zero-dependency Telegram Bot API library for Go

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
    me, err := bot.GetMe(ctx)

    if err != nil {
        panic(err)
    }

    fmt.Println("me", me)

    message, err := bot.SendMessage(ctx, &goram.SendMessageRequest{
        ChatId: goram.ChatId{Username: "somecoolchannel"},
        Text:   "Hello world",
    })

    if err != nil {
        panic(err)
    }

    fmt.Println("sent", message.MessageId)
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
