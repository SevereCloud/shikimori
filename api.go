/*
Package shikimori

# Authentication

OAuth2 is used for authentication. [OAuth2 guide].

# Restrictions

API access is limited by `5rps` and `90rpm`.

# Requirements

Add your Oauth2 Application name to [API.UserAgent] requests header.
Don’t mimic a browser.
Your IP address may be banned if you use API without properly set [API.UserAgent] header.

# Pagination in API

When you request N elements from paginated API, you will get N+1 results if API has next page.

[OAuth2 guide]: https://shikimori.me/oauth
*/
package shikimori

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type UnknownError struct {
	Status int
	Body   string
}

func (e UnknownError) Error() string {
	return fmt.Sprintf("shikimori: %d %s", e.Status, e.Body)
}

type NotFoundError struct{}

func (e NotFoundError) Error() string {
	return "shikimori: Not found"
}

type TooManyRequestsError struct{}

func (e TooManyRequestsError) Error() string {
	return "shikimori: 429 Retry later"
}

type UnprocessableEntityErrors []string

func (e UnprocessableEntityErrors) Error() string {
	return fmt.Sprintf("shikimori: unprocessable entity: %v", []string(e))
}

const (
	rps = 5
	rpm = 90
)

type API struct {
	Client      *http.Client
	BaseURL     string
	UserAgent   string
	AccessToken string

	limits *limits
}

func NewAPI() *API {
	return &API{
		BaseURL:     "https://shikimori.me/api/",
		UserAgent:   "Go-http-client/1.1 (+https://github.com/SevereCloud/shikimori)",
		AccessToken: "",
		Client:      http.DefaultClient,
		limits:      newLimits(newLimit(rps, time.Second), newLimit(rpm, time.Minute)),
	}
}

func (s *API) buildRequest(ctx context.Context, method string, path string, params interface{}) (*http.Request, error) {
	url := s.BaseURL + path

	buf := new(bytes.Buffer)

	err := json.NewEncoder(buf).Encode(params)
	if err != nil {
		return nil, fmt.Errorf("shikimori: encode: %w", err)
	}

	log.Println(buf)

	req, err := http.NewRequestWithContext(ctx, method, url, buf)
	if err != nil {
		return nil, fmt.Errorf("shikimori: build request: %w", err)
	}

	req.Header.Set("User-Agent", s.UserAgent)
	req.Header.Set("Content-Type", "application/json")

	if s.AccessToken != "" {
		req.Header.Set("Authorization", "Bearer "+s.AccessToken)
	}

	return req, nil
}

func (s *API) request(ctx context.Context, obj interface{}, method string, path string, params interface{}) error {
	req, err := s.buildRequest(ctx, method, path, params)
	if err != nil {
		return err
	}

	err = s.limits.RateLimit(ctx)
	if err != nil {
		return err
	}

	resp, err := s.Client.Do(req)
	if err != nil {
		return fmt.Errorf("shikimori: do request: %w", err)
	}
	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)

	err = s.handleStatus(resp, dec)
	if err != nil {
		return err
	}

	err = dec.Decode(&obj)
	if err != nil {
		return fmt.Errorf("shikimori decode: %w", err)
	}

	return nil
}

func (*API) handleStatus(resp *http.Response, dec *json.Decoder) error {
	switch resp.StatusCode {
	case http.StatusOK:
		break
	case http.StatusUnprocessableEntity:
		var errUnprocessableEntity UnprocessableEntityErrors

		err := dec.Decode(&errUnprocessableEntity)
		if err != nil {
			return fmt.Errorf("shikimori: decode 422 error: %w", err)
		}

		return &errUnprocessableEntity
	case http.StatusNotFound:
		return &NotFoundError{}
	case http.StatusTooManyRequests:
		return &TooManyRequestsError{}
	default:
		bytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("shikimori: decode %d error: %w", resp.StatusCode, err)
		}

		return &UnknownError{
			Status: resp.StatusCode,
			Body:   string(bytes),
		}
	}

	return nil
}

func (s *API) get(ctx context.Context, obj interface{}, path string, params interface{}) error {
	return s.request(ctx, obj, http.MethodGet, path, params)
}

// func (s *API) post(ctx context.Context, obj interface{}, path string, params interface{}) error {
// 	return s.request(ctx, obj, http.MethodPost, path, params)
// }

// func (s *API) put(ctx context.Context, obj interface{}, path string, params interface{}) error {
// 	return s.request(ctx, obj, http.MethodPut, path, params)
// }

// func (s *API) patch(ctx context.Context, obj interface{}, path string, params interface{}) error {
// 	return s.request(ctx, obj, http.MethodPatch, path, params)
// }

// func (s *API) delete(ctx context.Context, obj interface{}, path string, params interface{}) error {
// 	return s.request(ctx, obj, http.MethodDelete, path, params)
// }
