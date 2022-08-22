package addr

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type Addr struct {
	addr string
}

func NewAddr(addr string) *Addr {
	a := new(Addr)
	a.addr = addr
	return a
}

func (a *Addr) Validate() error {
	err := validation.Validate(a.addr,
		validation.Required,
		validation.Length(10, 100),
		is.URL,
	)
	return err
}

func (a *Addr) GetAddr() string {
	return a.addr
}
