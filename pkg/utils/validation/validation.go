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
	field := strings.ToLower(err.StructField())
	switch err.Tag() {
	case "required":
		return fmt.Sprintf("The %s field is required.", field)
	case "min":
		if err.Kind().String() == "string" {
			return fmt.Sprintf("The %s field must be at least %s characters.", field, err.Param())
		}
		return fmt.Sprintf("The %s field must be at least %s.", field, err.Param())
	case "max":
		if err.Kind().String() == "string" {
			return fmt.Sprintf("The %s field cannot be larger than %s characters.", field, err.Param())
		}
		return fmt.Sprintf("The %s field cannot be larger than %s.", field, err.Param())
	case "gt":
		if err.Kind().String() == "slice" {
			return fmt.Sprintf("The %s field must contain more than %s element(s).", field, err.Param())
		}
		return fmt.Sprintf("The %s field must be greater than %s.", field, err.Param())
	case "gte":
		if err.Kind().String() == "slice" {
			return fmt.Sprintf("The %s field must contain at least %s element(s).", field, err.Param())
		}
		return fmt.Sprintf("The %s field must be greater than or equal to %s.", field, err.Param())
	case "lt":
		if err.Kind().String() == "slice" {
			return fmt.Sprintf("The %s field cannot contain more than %s element(s).", field, err.Param())
		}
		return fmt.Sprintf("The %s field must be less than %s.", field, err.Param())
	case "lte":
		if err.Kind().String() == "slice" {
			return fmt.Sprintf("The %s field cannot contain more than %s element(s).", field, err.Param())
		}
		return fmt.Sprintf("The %s field must be less than or equal to %s.", field, err.Param())
	case "email":
		return fmt.Sprintf("The %s field must be valid email.", field)
	case "oneof":
		return fmt.Sprintf("The %s fiele must be one of %s", field, err.Param())
	default:
		fmt.Println(err.Tag())
		return fmt.Sprintf("The %s field is invalid.", field)
	}
}
