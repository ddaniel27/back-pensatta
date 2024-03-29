package validators

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidateRole(fl validator.FieldLevel) bool {
	role, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}

	myRole := strings.ToUpper(role)

	switch myRole {
	case "STUDENT", "TEACHER", "COORDINATOR":
		return true
	default:
		return false
	}
}
