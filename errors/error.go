package errors

import "errors"

var ErrInvalidOffset = errors.New("offset exceeds the total number of entries")
