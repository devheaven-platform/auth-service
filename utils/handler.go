package utils

import (
	"encoding/json"
	"net/http"

	"github.com/devheaven-platform/auth-service/models"
)

// RespondJSON is used as a helper function to write an
// go model to json and send it as response. It takes
// an http.ResponseWriter, status code and model as
// parameters.
func RespondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

// RespondError is used as a helper function to write
// an error as api response. It takes an http.ResponseWriter,
// name, message and status as parameters.
func RespondError(w http.ResponseWriter, name string, message string, status int) {
	RespondJSON(w, status, models.Error{
		Message: message,
	})
}

// RespondValidationError is used as an helper function
// to write an validation error as api response. It takes
// an http.ResponseWriter, name, message, status and errors
// map as parameters.
func RespondValidationError(w http.ResponseWriter, name string, message string, status int, errors map[string]string) {
	RespondJSON(w, status, models.Error{
		Message: message,
		Errors:  errors,
	})
}
