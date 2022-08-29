package url

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type Url struct {
	addr string
}

func NewUrl(url string) (*Url, error) {
	i := new(Url)
	i.addr = url

	err := i.validate()

	return i, err
}

func (a *Url) validate() error {
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
