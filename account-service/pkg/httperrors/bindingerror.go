package httperrors

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

func messageForTag(ve validator.FieldError) string {
	switch ve.Tag() {
	case "required":
		return fmt.Sprintf("field %s is required", ve.Field())
	case "email":
		return fmt.Sprintf("field %s must be a valid email", ve.Field())
	case "min=8":
		return fmt.Sprintf("field %s requires %d amount of characters", ve.Field(), 8)
	default:
		return fmt.Sprintf("error in field %s", ve.Field())
	}
}

func HttpBindingJSONError(err error, code int) ErrorRes {
	errs := Errors{}
	if verr, ok := err.(validator.ValidationErrors); ok {
		for _, ve := range verr {
			errs = append(errs, Error{Field: strings.ToLower(ve.Field()), Message: messageForTag(ve)})
		}
	}

	return ErrorRes{Errors: errs, StatusCode: code}
}
