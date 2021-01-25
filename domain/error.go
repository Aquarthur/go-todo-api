package domain

import "errors"

var (
	ErrNotFound = errors.New("not found")
	ErrConflict = errors.New("conflict")
)
