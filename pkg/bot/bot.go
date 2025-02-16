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
		return nil, res.error("getMe")
	}

	return res.Result, nil
}

func (b *Bot) SendMessage(
	ctx context.Context,
	request *types.SendMessageRequest,
) (*types.Message, error) {
	res, err := makeRequest[types.Message](ctx, b.options.Client, b.baseURL, "sendMessage", request)

	if err != nil {
		return nil, err
	}

	if !res.OK {
		return nil, res.error("sendMessage")
	}

	return res.Result, nil
}
