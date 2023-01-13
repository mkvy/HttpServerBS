package utils

import "errors"

var (
	ErrNotFound    = errors.New("Record not found")
	ErrDbConnect   = errors.New("Error connecting to database")
	ErrWrongEntity = errors.New("Invalid data")
)
