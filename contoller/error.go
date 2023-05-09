package contoller

import "errors"

var (
	// ErrInvalidID is returned when an invalid ID is provided to a method that need ID parameter.
	ErrInvalidID = errors.New("provided id was invalid")

	// ErrInvalidCredentials is returned when the provided credentials are invalid.
	ErrInvalidCredentials = errors.New("invalid credentials")
)
