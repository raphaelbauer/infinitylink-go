package httpmockable

import (
	"net/http"
)

// Just a tiny helper to make http.* requests mockable in tests
type HttpMockable interface {
	Get(url string) (resp *http.Response, err error)
}

type HttpImpl struct {
}

func (h *HttpImpl) Get(url string) (resp *http.Response, err error) {
	return http.Get(url)
}
