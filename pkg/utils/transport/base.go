package transport

import (
	"encoding/json"
	"net/http"

	"github.com/devheaven-platform/auth-service/pkg/domain"
)

// BaseTransport represents a base transport object.
type BaseTransport struct{}

// RespondJSON is used as a helper function to write an
// go model to json and send it as response. It takes
// an http.ResponseWriter, status code and model as
// parameters.
func (t BaseTransport) RespondJSON(res http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(err.Error()))
		return
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(status)
	res.Write([]byte(response))
}

// RespondError is used as a helper function to write
// an error as api response. It takes an http.ResponseWriter,
// message and status as parameters.
func (t BaseTransport) RespondError(res http.ResponseWriter, message string, status int) {
	t.RespondJSON(res, status, domain.APIError{
		Message: message,
	})
}

// RespondValidationError is used as an helper function
// to write an validation error as api response. It takes
// an http.ResponseWriter, message, status and errors
// map as parameters.
func (t BaseTransport) RespondValidationError(res http.ResponseWriter, message string, status int, errors map[string]string) {
	t.RespondJSON(res, status, domain.APIError{
		Message: message,
		Errors:  errors,
	})
}
