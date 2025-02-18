package bot

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"time"

	"github.com/TrixiS/goram/pkg/flood"
	"github.com/TrixiS/goram/pkg/types"
)

type Error struct {
	Method      string
	Description string
	ErrorCode   int
	Parameters  *types.ResponseParameters
}

func (a *Error) Error() string {
	return fmt.Sprintf("%s: %d %s", a.Method, a.ErrorCode, a.Description)
}

type apiResponse[R any] struct {
	OK          bool                      `json:"ok"`
	Description string                    `json:"description"`
	Result      R                         `json:"result"`
	ErrorCode   int                       `json:"error_code"`
	Parameters  *types.ResponseParameters `json:"parameters"`
}

func (a *apiResponse[R]) error(method string) *Error {
	return &Error{
		Method:      method,
		Description: a.Description,
		ErrorCode:   a.ErrorCode,
		Parameters:  a.Parameters,
	}
}

type apiRequest interface {
	WriteMultipart(*multipart.Writer)
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
		data.WriteMultipart(w)
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
			floodHandler.Enter(apiMethod, data)
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
		floodHandler.Handle(apiMethod, data, duration)

		if body != nil {
			body.Seek(0, io.SeekStart)
		}
	}
}
