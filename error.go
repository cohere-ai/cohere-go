package cohere

import (
	"fmt"
)

type ApiError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"-"`
}

func (e *ApiError) Error() string {
	return fmt.Sprintf("%s (%d)", e.Message, e.StatusCode)
}

func (e *ApiError) Is(target error) bool {
	_, ok := target.(*ApiError)
	return ok
}
