package url

import (
	"io"
	"net/http"
)

// Interface class providing data on URLs.
type Provider interface {
	ReadBody(url *Url) (string, error)
}

type WebProvider struct{}

func NewWebProvider() Provider {
	return new(WebProvider)
}

func (p *WebProvider) ReadBody(url *Url) (string, error) {
	addr := url.Addr()

	r, err := http.Get(addr)
	if err != nil {
		return "", err
	}
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
