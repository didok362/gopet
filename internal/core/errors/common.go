package core_errors

import "errors"

var (
	ErrNotFound        = errors.New("Not Found")
	ErrInvalidArgumnet = errors.New("Invalid Argument")
	ErrConflict        = errors.New("Conflict")
)
