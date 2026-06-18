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

const maxRetries = 5

func makeRequest[R any](
	ctx context.Context,
	client *http.Client,
	baseURL string,
	apiMethod string,
	floodHandler flood.Handler,
	data apiRequest,
) (*apiResponse[R], error) {
	url := baseURL + apiMethod
	body, contentType := prepareRequestBody(data)

	for attempt := 0; attempt < maxRetries; attempt++ {
		response, retryRequired, err := doRequestAttempt[R](
			ctx,
			client,
			url,
			contentType,
			body,
			apiMethod,
			floodHandler,
			data,
			attempt,
		)

		if err != nil || !retryRequired {
			return response, err
		}

		if body != nil {
			body.Seek(0, io.SeekStart)
		}
	}

	return nil, context.DeadlineExceeded
}

func makeVoidRequest(
	ctx context.Context,
	client *http.Client,
	baseURL string,
	apiMethod string,
	floodHandler flood.Handler,
	data apiRequest,
) error {
	url := baseURL + apiMethod
	body, contentType := prepareRequestBody(data)

	for attempt := 0; attempt < maxRetries; attempt++ {
		retryRequired, err := doVoidRequestAttempt(
			ctx,
			client,
			url,
			contentType,
			body,
			apiMethod,
			floodHandler,
			data,
			attempt,
		)

		if err != nil || !retryRequired {
			return err
		}

		body.Seek(0, io.SeekStart)
	}

	return context.DeadlineExceeded
}

func doRequestAttempt[R any](
	ctx context.Context,
	client *http.Client,
	url string,
	contentType string,
	body io.ReadSeeker,
	apiMethod string,
	floodHandler flood.Handler,
	data apiRequest,
	attempt int,
) (*apiResponse[R], bool, error) {
	res, err := executeHTTPRequest(
		ctx,
		client,
		url,
		contentType,
		body,
		apiMethod,
		floodHandler,
		data,
	)

	if err != nil {
		return nil, false, err
	}

	defer res.Body.Close()

	response := &apiResponse[R]{}

	if err := json.NewDecoder(res.Body).Decode(response); err != nil {
		return nil, false, err
	}

	if response.OK {
		return response, false, nil
	}

	if response.ErrorCode != http.StatusTooManyRequests ||
		response.Parameters == nil ||
		floodHandler == nil ||
		attempt == maxRetries-1 {

		return response, false, response.error(apiMethod)
	}

	duration := time.Second * time.Duration(response.Parameters.RetryAfter)
	floodHandler.Handle(ctx, apiMethod, data, duration)

	return nil, true, nil
}

func doVoidRequestAttempt(
	ctx context.Context,
	client *http.Client,
	url string,
	contentType string,
	body io.ReadSeeker,
	apiMethod string,
	floodHandler flood.Handler,
	data apiRequest,
	attempt int,
) (bool, error) {
	res, err := executeHTTPRequest(
		ctx,
		client,
		url,
		contentType,
		body,
		apiMethod,
		floodHandler,
		data,
	)

	if err != nil {
		return false, err
	}

	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		return false, nil
	}

	response := &apiResponse[json.RawMessage]{}

	if err := json.NewDecoder(res.Body).Decode(response); err != nil {
		return false, err
	}

	if response.OK {
		return false, nil
	}

	if response.ErrorCode != http.StatusTooManyRequests ||
		response.Parameters == nil ||
		floodHandler == nil ||
		attempt == maxRetries-1 {

		return false, response.error(apiMethod)
	}

	duration := time.Second * time.Duration(response.Parameters.RetryAfter)
	floodHandler.Handle(ctx, apiMethod, data, duration)

	return true, nil
}

func prepareRequestBody(data apiRequest) (io.ReadSeeker, string) {
	if data == nil {
		return nil, "multipart/form-data"
	}

	buf := bytes.Buffer{}
	w := multipart.NewWriter(&buf)
	data.writeMultipart(w)
	w.Close()

	return bytes.NewReader(buf.Bytes()), w.FormDataContentType()
}

func executeHTTPRequest(
	ctx context.Context,
	client *http.Client,
	url string,
	contentType string,
	body io.ReadSeeker,
	apiMethod string,
	floodHandler flood.Handler,
	data apiRequest,
) (*http.Response, error) {
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

	return res, nil
}
