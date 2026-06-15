package goram

import (
	"context"
	"io"
	"net/http"
	"strings"

	"github.com/TrixiS/goram/flood"
)

const DefaultAPIBaseURL = "https://api.telegram.org"

type BotOptions struct {
	Token        string        // Required
	Client       *http.Client  // Optional. If Client is nil, http.DefaultClient will be used
	FloodHandler flood.Handler // Optional. If FloodHandler is nil, 429 flood error will be propagated to the caller of a flooded method
	BaseURL      string        // Optional. If BaseUrl is empty, goram.DefaultAPIBaseURL will be used
}

// Holds all methods of Telegram Bot API.
//
// For example: Bot.SendMessage()
//
// If a method call fails, it can return *goram.APIError as the second return value. The returned error can be checked whether it's an API error as easy as:
//
// apiError, ok := err.(*goram.APIError)
type Bot struct {
	Options BotOptions
	baseURL string
}

func NewBot(options BotOptions) *Bot {
	if options.BaseURL == "" {
		options.BaseURL = DefaultAPIBaseURL
	}

	baseURL := options.BaseURL + "/bot" + options.Token + "/"

	if options.Client == nil {
		options.Client = http.DefaultClient
	}

	return &Bot{
		Options: options,
		baseURL: baseURL,
	}
}

type ErrDownloadFile struct {
	Response *http.Response
	File     *File
}

func (e ErrDownloadFile) Error() string {
	return "downloadFile: " + e.Response.Status
}

// Creates this url https://api.telegram.org/file/bot<token>/<file_path>
// See bot.GetFile()
func MakeFileDownloadURL(baseURL string, token string, filePath string) string {
	urlBuilder := strings.Builder{}
	urlBuilder.WriteString(baseURL)
	urlBuilder.WriteString("/file/bot")
	urlBuilder.WriteString(token)
	urlBuilder.WriteByte('/')
	urlBuilder.WriteString(filePath)
	return urlBuilder.String()
}

// Downloads a file by file id using provided or default http client.
// Writes response to dst and returns amount of bytes written and an error.
// This function does not close or seek the provided writer.
//
// If download http response status != 200, returns goram.ErrDownloadFile
func (b *Bot) DownloadFile(ctx context.Context, fileID string, dst io.Writer) (int64, error) {
	file, err := b.GetFile(ctx, &GetFileRequest{FileID: fileID})

	if err != nil {
		return 0, err
	}

	r, err := b.OpenFile(ctx, file)

	if err != nil {
		return 0, err
	}

	defer r.Close()
	return io.Copy(dst, r)
}

// Opens a file got from bot.GetFile() for downloading.
// The caller is responsible for closing the returned io.ReaderCloser
//
// If download http response status != 200, returns goram.ErrDownloadFile
func (b *Bot) OpenFile(ctx context.Context, file *File) (io.ReadCloser, error) {
	downloadURL := MakeFileDownloadURL(b.Options.BaseURL, b.Options.Token, file.FilePath)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, downloadURL, nil)

	if err != nil {
		return nil, err
	}

	res, err := b.Options.Client.Do(req)

	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		res.Body.Close()
		return nil, ErrDownloadFile{Response: res, File: file}
	}

	return res.Body, nil
}
