package validation

import (
	"fmt"
	"strings"

	validator "gopkg.in/go-playground/validator.v9"
)

// Validate verifies a given request using validator
// tags from the go-playground/validator.v9 package.
// The function takes an interface as parameter.
// The errors will be returned in the form of a map
// the key is the field name, the value is the error
// value. Nil will be returned if the request doesn't
// contain errors.
func Validate(request interface{}) map[string]string {
	validate := validator.New()
	err := validate.Struct(request)

	if err != nil {
		errs := map[string]string{}

		for _, err := range err.(validator.ValidationErrors) {
			errs[strings.ToLower(err.StructField())] = formatMessage(err)
		}

		return errs
	}

	return nil
}

// formatMessage formats a FieldError in a more human
// readable form. It takes an instance of validator.FieldError
// as parameter and returns an string as result.
func formatMessage(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return fmt.Sprintf("The %s field is required.", err.StructField())
	case "min":
		if err.Kind().String() == "string" {
			return fmt.Sprintf("The %s field must be at least %s characters.", err.StructField(), err.Param())
		}
		return fmt.Sprintf("The %s field must be at least %s.", err.StructField(), err.Param())
	case "max":
		if err.Kind().String() == "string" {
			return fmt.Sprintf("The %s field cannot be larger than %s characters.", err.StructField(), err.Param())
		}
		return fmt.Sprintf("The %s field cannot be larger than %s.", err.StructField(), err.Param())
	default:
		return fmt.Sprintf("The %s field is invalid.", err.StructField())
	}
}
