package utils

import (
	"encoding/json"
	"net/http"

	"github.com/devheaven-platform/golang/models"
)

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

func RespondError(w http.ResponseWriter, name string, message string, status int) {
	RespondJSON(w, status, models.Error{
		Name:    name,
		Message: message,
		Status:  status,
	})
}

func RespondValidationError(w http.ResponseWriter, name string, message string, status int, errors map[string]string) {
	RespondJSON(w, status, models.Error{
		Name:    name,
		Message: message,
		Status:  status,
		Errors:  errors,
	})
}
