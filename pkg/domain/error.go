package domain

// Error represents an API error. An Error always contains
// a message and an optional errors map if the error is an
// validation error.
type Error struct {
	Message string            `json:"message"`
	Errors  map[string]string `json:"errors,omitempty"`
}