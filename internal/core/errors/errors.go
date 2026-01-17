package errors

import "errors"

var (
	ErrNotFound          = errors.New("entity not found")
	ErrInsufficientFunds = errors.New("insufficient funds")
	ErrValidation        = errors.New("validation error")
	ErrBinding           = errors.New("binding error")
)
