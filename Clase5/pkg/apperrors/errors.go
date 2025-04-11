package apperrors

import "errors"

var (
	ErrNotFound       = NewAppError(404, "Resource not found")
	ErrBadRequest     = NewAppError(400, "Bad request")
	ErrInternalServer = NewAppError(500, "Internal server error")
	ErrUnauthorized   = NewAppError(401, "Unauthorized")
)

type AppError struct {
	Code    int
	Message string
	Err     error
}

func (e *AppError) Error() string {
	return e.Message
}

func NewAppError(code int, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Err:     errors.New(message),
	}
}

// Helpers
func IsAppError(err error) (*AppError, bool) {
	appErr, ok := err.(*AppError)
	return appErr, ok
}
