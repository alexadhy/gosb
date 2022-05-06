package dto

import (
	"net/url"

	"github.com/sonh/qs"
)

type URLer interface {
	URL(string) string
}

// encodeQS takes baseURL and struct with `qs` tag
// encodes struct with qs tag to URL string
// will return the encoded url string and error if any
func encodeQS[T URLer](baseURL string, input T) (string, error) {
	_, err := url.Parse(baseURL)
	if err != nil {
		return "", err
	}

	encoder := qs.NewEncoder()
	values, err := encoder.Values(input)
	if err != nil {
		return "", err
	}

	return baseURL + "?" + values.Encode(), nil
}
