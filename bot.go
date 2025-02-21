package goram

import (
	"context"
	"io"
	"net/http"
	"strings"

	"github.com/TrixiS/goram/flood"
)

type BotOptions struct {
	Token        string        // Required
	Client       *http.Client  // Optional. If Client is nil, http.DefaultClient will be used
	FloodHandler flood.Handler // Optional. If FloodHandler is nil, 429 flood error will be propagated to the caller of a flooded method
	BaseUrl      string        // Optional. If BaseUrl is empty, https://api.telegram.org/ will be used
}

// Holds all methods of Telegram Bot API.
//
// For example: Bot.SendMessage()
//
// If a method call fails, it can return *goram.APIError as the second return value. The returned error can be checked whether it's an API error as easy as:
//
// apiError, ok := err.(*goram.APIError)
type Bot struct {
	options BotOptions
	baseUrl string
}

func NewBot(options BotOptions) *Bot {
	if options.BaseUrl == "" {
		options.BaseUrl = "https://api.telegram.org"
	}

	baseUrl := options.BaseUrl + "/bot" + options.Token + "/"

	if options.Client == nil {
		options.Client = http.DefaultClient
	}

	return &Bot{
		options: options,
		baseUrl: baseUrl,
	}
}

type ErrDownloadFile struct {
	Response *http.Response
	File     *File
}

func (e ErrDownloadFile) Error() string {
	return "downloadFile: " + e.Response.Status
}

// Downloads a file by file id using provided or default http client.
// Writes response to dst and returns amount of bytes written and an error.
// This function does not close or seek the provided writer.
//
// If download http request status != 200, returns ErrDownloadFile
func (b *Bot) DownloadFile(ctx context.Context, fileId string, dst io.Writer) (int64, error) {
	file, err := b.GetFile(ctx, &GetFileRequest{FileId: fileId})

	if err != nil {
		return 0, err
	}

	urlBuilder := strings.Builder{}
	urlBuilder.WriteString(b.options.BaseUrl)
	urlBuilder.WriteString("/file/bot")
	urlBuilder.WriteString(b.options.Token)
	urlBuilder.WriteRune('/')
	urlBuilder.WriteString(file.FilePath)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, urlBuilder.String(), nil)

	if err != nil {
		return 0, err
	}

	res, err := b.options.Client.Do(req)

	if err != nil {
		return 0, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return 0, ErrDownloadFile{Response: res, File: file}
	}

	return io.Copy(dst, res.Body)
}
