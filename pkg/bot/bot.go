package bot

import (
	"net/http"

	"github.com/TrixiS/goram/pkg/flood"
)

const apiUrl = "https://api.telegram.org/bot"

type BotOptions struct {
	Token        string        // Required
	Client       *http.Client  // Optional. If Client is nil, http.DefaultClient will be used
	FloodHandler flood.Handler // Optional. If FloodHandler is nil, 429 flood error will be propagated to the caller of a flooded method
	BaseUrl      string        // Optional. If BaseUrl is empty, https://api.telegram.org/bot will be used
}

// Holds all methods of Telegram Bot API.
//
// For example: Bot.SendMessage()
type Bot struct {
	options BotOptions
	baseUrl string
}

func NewBot(options BotOptions) *Bot {
	baseUrl := apiUrl

	if options.BaseUrl != "" {
		baseUrl = options.BaseUrl
	}

	baseUrl = baseUrl + options.Token + "/"

	if options.Client == nil {
		options.Client = http.DefaultClient
	}

	return &Bot{
		options: options,
		baseUrl: baseUrl,
	}
}
