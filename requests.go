package tgbotstgapi

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

var (
	ErrAPIError = errors.New("api error")
)

func request[T, E any](ctx context.Context, cl *http.Client, token string, method string, data *T) (*E, error) {
	body, err := requestRaw(ctx, cl, token, method, data)
	if err != nil {
		return nil, err
	}

	var res E
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}

	return &res, nil
}

func requestRaw[T any](ctx context.Context, cl *http.Client, token string, method string, data *T) ([]byte, error) {
	body, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("marshal body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url(token, method), bytes.NewReader(body))
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	resp, err := cl.Do(req)
	if err != nil {
		return nil, fmt.Errorf("make request: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%w: status %d", ErrAPIError, resp.StatusCode)
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body: %w", err)
	}
	defer resp.Body.Close()

	return respBody, nil
}

func url(token, method string) string {
	return fmt.Sprintf("https://api.telegram.org/bot%s/%s", token, method)
}
