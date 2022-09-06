// Package url manages and manipulates URLs.

package url

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

// The struct Url is a value object.
type Url struct {
	addr string
}

func NewUrl(url string) (*Url, error) {
	i := new(Url)
	i.addr = url

	err := i.validate()

	return i, err
}

func (u *Url) validate() error {
	err := validation.Validate(u.addr,
		validation.Required,
		validation.Length(10, 100),
		is.URL,
	)
	return err
}

func (u *Url) Addr() string {
	return u.addr
}
