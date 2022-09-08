package url

import (
	"io"
	"net/http"
)

// Interface class providing data on URL.
type Provider interface {
	ReadBody(url *Url) (string, error)
}

// A WebPrivider uses web as the data source.
// This is usually used in production.
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

// A InMemDummyProvider uses in-memory as the data source.
// This is usually used in development such as testing.
type InMemDummyProvider struct{}

func NewInMemDummyProvider() Provider {
	return new(InMemDummyProvider)
}

func (p *InMemDummyProvider) ReadBody(url *Url) (string, error) {
	body := `
	<title>InMem TITLE</title>
	<h1>InMem H1</h1>
	`
	return body, nil
}
