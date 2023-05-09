package database

import "errors"

var (
	// ErrIDNotFound is returned when an invalid ID is provided to a method like Delete.
	ErrIDNotFound = errors.New("provided id was not found")
)
