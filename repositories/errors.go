package repositories

import "errors"

// ErrNotImplemented is returned by repository methods that are not yet implemented
var ErrNotImplemented = errors.New("not implemented")
