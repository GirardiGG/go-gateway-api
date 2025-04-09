package domain

import "errors"

var (
	ErrAccountNotFound    = errors.New("Account not found")
	ErrDuplicateAPIKey    = errors.New("api key already exists")
	ErrInvoiceNotFound    = errors.New("invoice not found")
	ErrUnauthorizedAccess = errors.New("unauthorized access")
)
