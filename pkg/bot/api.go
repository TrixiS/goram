package bot

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

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

type apiRequest interface {
	WriteMultipart(*multipart.Writer)
}

// TODO: handle flood (seek body buffer)
func makeRequest[R any](
	ctx context.Context,
	client *http.Client,
	baseURL string,
	apiMethod string,
	data apiRequest,
) (*apiResponse[R], error) {
	url := baseURL + apiMethod

	body := io.Reader(nil)
	contentType := "multipart/form-data"

	if data != nil {
		buf := &bytes.Buffer{}
		w := multipart.NewWriter(buf)
		data.WriteMultipart(w)
		w.Close()
		contentType = w.FormDataContentType()
		body = buf
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, body)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", contentType)
	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	response := &apiResponse[R]{}

	if err = json.NewDecoder(res.Body).Decode(response); err != nil {
		return nil, err
	}

	if !response.OK {
		return response, &Error{
			Method:      apiMethod,
			Description: response.Description,
			ErrorCode:   response.ErrorCode,
			Parameters:  response.Parameters,
		}
	}

	return response, nil
}
