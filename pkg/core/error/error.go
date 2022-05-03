package error

import "errors"

var (
	ErrMethodNotFound = errors.New("method not found")
	ErrTokenNotFound  = errors.New("token not found")
)
