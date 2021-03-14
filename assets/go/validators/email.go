package validators

import (
	"net/mail"
	"strings"

	"github.com/go-playground/validator/v10"
)

// Email validates and formats an email address
var Email validator.Func = func(fl validator.FieldLevel) bool {
	v, ok := fl.Field().Addr().Interface().(*string)
	if !ok {
		return false
	}

	*v = strings.TrimSpace(*v)

	e, err := mail.ParseAddress("<" + *v + ">")
	if err != nil {
		return false
	}

	*v = e.Address

	*v = strings.ToLower(*v)

	return true
}
