package errors

import (
	"errors"
)

var ErrInvalidOffset = errors.New("offset exceeds the total number of entries")
var ErrInvalidSorting = errors.New("Invalid Sorting type")