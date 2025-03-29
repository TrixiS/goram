package goram

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"time"

	"github.com/TrixiS/goram/flood"
)

// Represents Telegram API error
type APIError struct {
	Method      string // Called API method
	Description string
	ErrorCode   int
	Parameters  *ResponseParameters
}

func (a *APIError) Error() string {
	return fmt.Sprintf("%s: %d %s", a.Method, a.ErrorCode, a.Description)
}

type apiResponse[R any] struct {
	OK          bool                `json:"ok"`
	Description string              `json:"description"`
	Result      R                   `json:"result"`
	ErrorCode   int                 `json:"error_code"`
	Parameters  *ResponseParameters `json:"parameters"`
}

func (a *apiResponse[R]) error(method string) *APIError {
	return &APIError{
		Method:      method,
		Description: a.Description,
		ErrorCode:   a.ErrorCode,
		Parameters:  a.Parameters,
	}
}

type apiRequest interface {
	writeMultipart(*multipart.Writer)
}

func makeRequest[R any](
	ctx context.Context,
	client *http.Client,
	baseUrl string,
	apiMethod string,
	floodHandler flood.Handler,
	data apiRequest,
) (*apiResponse[R], error) {
	url := baseUrl + apiMethod
	contentType := "multipart/form-data"
	body := io.ReadSeeker(nil)

	if data != nil {
		buf := &bytes.Buffer{}
		w := multipart.NewWriter(buf)
		data.writeMultipart(w)
		w.Close()
		contentType = w.FormDataContentType()
		body = bytes.NewReader(buf.Bytes())
	}

	for {
		req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, body)

		if err != nil {
			return nil, err
		}

		req.Header.Set("Content-Type", contentType)

		if floodHandler != nil {
			floodHandler.Enter(ctx, apiMethod, data)
		}

		res, err := client.Do(req)

		if err != nil {
			return nil, err
		}

		response := &apiResponse[R]{}
		err = json.NewDecoder(res.Body).Decode(response)
		res.Body.Close()

		if err != nil {
			return nil, err
		}

		if response.OK {
			return response, nil
		}

		if response.ErrorCode != http.StatusTooManyRequests ||
			response.Parameters == nil ||
			floodHandler == nil {

			return response, response.error(apiMethod)
		}

		duration := time.Second * time.Duration(response.Parameters.RetryAfter)
		floodHandler.Handle(ctx, apiMethod, data, duration)

		if body != nil {
			body.Seek(0, io.SeekStart)
		}
	}
}

func makeVoidRequest(ctx context.Context,
	client *http.Client,
	baseUrl string,
	apiMethod string,
	floodHandler flood.Handler,
	data apiRequest,
) error {
	url := baseUrl + apiMethod

	buf := &bytes.Buffer{}
	w := multipart.NewWriter(buf)
	data.writeMultipart(w)
	w.Close()

	contentType := w.FormDataContentType()
	body := bytes.NewReader(buf.Bytes())

	for {
		req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, body)

		if err != nil {
			return err
		}

		req.Header.Set("Content-Type", contentType)

		if floodHandler != nil {
			floodHandler.Enter(ctx, apiMethod, data)
		}

		res, err := client.Do(req)

		if err != nil {
			return err
		}

		if res.StatusCode == http.StatusOK {
			return nil
		}

		response := &apiResponse[json.RawMessage]{}
		err = json.NewDecoder(res.Body).Decode(response)
		res.Body.Close()

		if err != nil {
			return err
		}

		if response.OK {
			return nil
		}

		if response.ErrorCode != http.StatusTooManyRequests ||
			response.Parameters == nil ||
			floodHandler == nil {

			return response.error(apiMethod)
		}

		duration := time.Second * time.Duration(response.Parameters.RetryAfter)
		floodHandler.Handle(ctx, apiMethod, data, duration)

		if body != nil {
			body.Seek(0, io.SeekStart)
		}
	}
}
