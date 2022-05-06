package gosb

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/go-resty/resty/v2"
)

func TestCallHTTP(t *testing.T) {
	hc := defaultHC("API_TOKEN")

	cases := []struct {
		name          string
		input         func(*resty.Request) (*resty.Response, error)
		expected      map[string]any
		expectedError error
	}{
		{
			name: "should be able to do GET request",
			input: func(r *resty.Request) (*resty.Response, error) {
				uri := "https://jsonplaceholder.typicode.com/todos/1"
				return r.Get(uri)
			},
			expected: map[string]any{
				"userId":    float64(1),
				"id":        float64(1),
				"title":     "delectus aut autem",
				"completed": false,
			},
			expectedError: nil,
		},
		{
			name: "should be able to do POST request",
			input: func(r *resty.Request) (*resty.Response, error) {
				uri := "https://jsonplaceholder.typicode.com/posts"
				input := struct {
					title  string
					body   string
					userID int
				}{
					"foo",
					"bar",
					2,
				}
				return r.SetBody(&input).Post(uri)
			},
			expected:      map[string]any{"id": float64(101)},
			expectedError: nil,
		},
	}

	for _, tt := range cases {
		resp, err := makeRequest[map[string]any](context.Background(), hc, tt.input)
		assert.NoError(t, err)
		assert.Equal(t, tt.expected, *resp)
	}
}
