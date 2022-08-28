package url

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type Url struct {
	addr string
}

func NewUrl(url string) *Url {
	a := new(Url)
	a.addr = url
	return a
}

func (a *Url) Validate() error {
	err := validation.Validate(a.addr,
		validation.Required,
		validation.Length(10, 100),
		is.URL,
	)
	return err
}

func (a *Url) GetAddr() string {
	return a.addr
}
