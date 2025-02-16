package bot

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

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

type apiResponse[T any] struct {
	Ok          bool                      `json:"ok"`
	Description string                    `json:"description"`
	Result      *T                        `json:"result"`
	ErrorCode   int                       `json:"error_code"`
	Parameters  *types.ResponseParameters `json:"parameters"`
}

type APIError struct {
	Description string
	ErrorCode   int
	Parameters  *types.ResponseParameters
}

func (a *APIError) Error() string {
	stringErrorCode := strconv.Itoa(a.ErrorCode)
	return stringErrorCode + " " + a.Description
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

	if !res.Ok {
		return nil, &APIError{
			Description: res.Description,
			ErrorCode:   res.ErrorCode,
			Parameters:  res.Parameters,
		}
	}

	return res.Result, nil
}

func makeRequest[T any](
	ctx context.Context,
	client *http.Client,
	baseURL string,
	apiMethod string,
	data any,
) (*apiResponse[T], error) {
	url := baseURL + apiMethod

	var body io.Reader = http.NoBody

	if data != nil {
		buf := &bytes.Buffer{}

		if err := json.NewEncoder(buf).Encode(data); err != nil {
			return nil, err
		}

		body = buf
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, body)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	r := &apiResponse[T]{}
	return r, json.NewDecoder(res.Body).Decode(r)
}
