package gosb

import (
	"context"

	"github.com/go-resty/resty/v2"

	"github.com/alexadhy/gosb/dto"
)

// defaultHC returns default http.Client (*resty.Client) to use
// in facing SendBird API
func defaultHC(apiToken string, opts ...resty.Option) *resty.Client {
	options := &resty.Options{}
	c := resty.New()

	for _, opt := range opts {
		opt(options)
	}
	c.SetRetryCount(3)

	c.Header.Add("Api-Token", apiToken)
	return c
}

type urler interface {
	URL(string) string
}

func makeRequestURL(input urler, c *Client) (string, error) {
	u := input.URL(c.baseURL)
	if u == "" {
		return "", dto.NewErr(dto.ErrEmptyURLValue, "empty request URL")
	}
	return u, nil
}

func makeRequest[T any](ctx context.Context, hc *resty.Client, fn func(request *resty.Request) (*resty.Response, error)) (*T, error) {
	var out T
	resp, err := fn(hc.R().
		SetResult(&out).
		SetError(&out).
		SetContext(ctx))

	if err != nil {
		return nil, err
	}

	result := resp.Result().(*T)
	return result, nil
}

func getRequest[T any](ctx context.Context, hc *resty.Client, u string) (*T, error) {
	return makeRequest[T](ctx, hc, func(request *resty.Request) (*resty.Response, error) {
		return request.Get(u)
	})
}

func postRequest[T any](ctx context.Context, hc *resty.Client, u string, body any) (*T, error) {
	return makeRequest[T](ctx, hc, func(request *resty.Request) (*resty.Response, error) {
		return request.
			SetBody(body).
			Post(u)
	})
}

func putRequest[T any](ctx context.Context, hc *resty.Client, u string, body any) (*T, error) {
	return makeRequest[T](ctx, hc, func(request *resty.Request) (*resty.Response, error) {
		return request.
			SetBody(body).
			Put(u)
	})
}

func deleteRequest[T any](ctx context.Context, hc *resty.Client, u string) (*T, error) {
	return makeRequest[T](ctx, hc, func(request *resty.Request) (*resty.Response, error) {
		return request.Delete(u)
	})
}
