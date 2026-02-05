package services

import "errors"

var (
	// ErrInvalidCredentials returned when signin password does not match
	ErrInvalidCredentials = errors.New("invalid credentials")
)
