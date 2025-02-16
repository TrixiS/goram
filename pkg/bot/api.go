package bot

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/TrixiS/goram/pkg/types"
)

type apiResponse[R any] struct {
	OK          bool                      `json:"ok"`
	Description string                    `json:"description"`
	Result      *R                        `json:"result"`
	ErrorCode   int                       `json:"error_code"`
	Parameters  *types.ResponseParameters `json:"parameters"`
}

type Error struct {
	Method      string
	Description string
	ErrorCode   int
	Parameters  *types.ResponseParameters
}

func (a *Error) Error() string {
	return fmt.Sprintf("%s: %d %s", a.Method, a.ErrorCode, a.Description)
}

func makeRequest[R any](
	ctx context.Context,
	client *http.Client,
	baseURL string,
	apiMethod string,
	data any,
) (*apiResponse[R], error) {
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

	r := &apiResponse[R]{}
	return r, json.NewDecoder(res.Body).Decode(r)
}
