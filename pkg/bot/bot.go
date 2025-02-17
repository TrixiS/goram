package bot

import (
	"net/http"
)

const apiURL = "https://api.telegram.org/bot"

type BotOptions struct {
	Token  string
	Client *http.Client
}

// TODO: AfterRequest funcs
type Bot struct {
	options BotOptions
	baseURL string
}

func NewBot(options BotOptions) *Bot {
	baseURL := apiURL + options.Token + "/"

	if options.Client == nil {
		options.Client = http.DefaultClient
	}

	return &Bot{
		options: options,
		baseURL: baseURL,
	}
}
