package setup

import (
	"regexp"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func SetupValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("key", validateKey())
	}
}

func validateKey() func(fl validator.FieldLevel) bool {
	keyRegexp := regexp.MustCompile(`^[a-zA-Z-]+$`)

	return func(fl validator.FieldLevel) bool {
		keyValue, ok := fl.Field().Interface().(string)
		if !ok {
			return false
		}

		return keyRegexp.MatchString(keyValue)
	}
}
