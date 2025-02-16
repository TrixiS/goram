package bot

import (
	"context"
	"net/http"

	"github.com/TrixiS/goram/pkg/types"
)

const apiURL = "https://api.telegram.org/bot"

type BotOptions struct {
	Token  string
	Client *http.Client
}

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

func (b *Bot) GetMe(ctx context.Context) (*types.User, error) {
	res, err := makeRequest[types.User](ctx, b.options.Client, b.baseURL, "getMe", nil)

	if err != nil {
		return nil, err
	}

	if !res.OK {
		return nil, &Error{
			Method:      "getMe",
			Description: res.Description,
			ErrorCode:   res.ErrorCode,
			Parameters:  res.Parameters,
		}
	}

	return res.Result, nil
}
