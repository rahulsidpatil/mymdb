package utils

// NewError example
func NewError(status int, err error) HTTPError {
	return HTTPError{
		Message: err.Error(),
	}
}

// HTTPError example
type HTTPError struct {
	Message string `json:"message,omitempty" example:"status bad request"`
}
