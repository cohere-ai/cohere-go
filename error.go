package cohere

import (
	"fmt"
)

type APIError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"-"`
}

func (e *APIError) Error() string {
	return fmt.Sprintf("%s (%d)", e.Message, e.StatusCode)
}

func (e *APIError) Is(target error) bool {
	_, ok := target.(*APIError)
	return ok
}
