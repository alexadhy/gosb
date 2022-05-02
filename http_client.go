package gosb

import (
	"github.com/go-resty/resty/v2"
)

// defaultHC returns default http.Client (*resty.Client) to use
// in facing SendBird API
func defaultHC(opts ...resty.Option) *resty.Client {
	options := &resty.Options{}
	c := resty.New()

	for _, opt := range opts {
		opt(options)
	}
	c.SetRetryCount(3)
	return c
}
