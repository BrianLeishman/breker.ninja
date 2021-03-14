package validators

import (
	"strings"

	"github.com/StirlingMarketingGroup/go-namecase"
	"github.com/go-playground/validator/v10"
)

var formatter = namecase.New()

// Name validates and formats a person's name
var Name validator.Func = func(fl validator.FieldLevel) bool {
	v, ok := fl.Field().Addr().Interface().(*string)
	if !ok {
		return false
	}

	*v = strings.TrimSpace(*v)

	*v = formatter.NameCase(*v)

	return true
}
