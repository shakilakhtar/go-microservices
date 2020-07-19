package utils

import (
	"errors"

	"dt-logger/logger"
)

var (
	// ErrUnknown is used when a item could not be found.
	ErrUnknown = NewError("unknown asset")

	// ErrInvalidArgument is returned when one or more arguments are invalid.
	ErrInvalidArgument = NewError("invalid argument")
)

// HandleError creates an errors.error type with a given string, logs the error and returns it
func HandleError(errMessage string) error {
	err := errors.New(errMessage)
	logger.Error(err.Error())
	return err
}

//Creates an error for a given error message
func NewError(msg string) error {
	return errors.New(msg)
}
