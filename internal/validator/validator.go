package validator

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

// Default constructs a new validator instance with custom configuration and new rules.
func Default() *validator.Validate {
	validate := validator.New()

	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}

		return name
	})

	return validate
}
