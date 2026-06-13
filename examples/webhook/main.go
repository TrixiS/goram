package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/TrixiS/goram"
	"github.com/TrixiS/goram/examples/routes"
	"github.com/TrixiS/goram/handlers"
)

var (
	token      = os.Getenv("BOT_TOKEN")
	url        = os.Getenv("URL")
	secret     = os.Getenv("SECRET")
	listenAddr = os.Getenv("LISTEN_ADDR")
)

func main() {
	ctx := context.Background()
	bot := goram.NewBot(goram.BotOptions{Token: token})

	err := bot.SetWebhookVoid(ctx, &goram.SetWebhookRequest{
		URL:                url,
		SecretToken:        secret,
		DropPendingUpdates: true,
	})

	if err != nil {
		panic(err)
	}

	router := routes.CreateRouter(0)

	http.HandleFunc("/updates", func(w http.ResponseWriter, r *http.Request) {
		update := goram.Update{}

		if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
			fmt.Println(err)
		} else {
			go router.FeedUpdates(ctx, bot, []goram.Update{update}, handlers.Data{})
		}

		r.Body.Close()
		w.WriteHeader(http.StatusOK)
	})

	http.ListenAndServe(listenAddr, nil)
}
