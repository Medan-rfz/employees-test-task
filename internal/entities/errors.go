package entities

import "errors"

var (
	ErrEmployeeId       = errors.New("incorrect employee id")
	ErrEmployeeNotFound = errors.New("employee not found")
)
