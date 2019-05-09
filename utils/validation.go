package utils

import (
	"fmt"

	validator "gopkg.in/go-playground/validator.v9"
)

func Validate(request interface{}) map[string]string {
	validate := validator.New()
	err := validate.Struct(request)

	if err != nil {
		errs := map[string]string{}

		for _, err := range err.(validator.ValidationErrors) {
			// TODO: convert field name to lowercase
			errs[err.StructField()] = formatMessage(err)
		}

		return errs
	}

	return nil
}

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
