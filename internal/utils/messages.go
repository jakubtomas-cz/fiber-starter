package utils

import (
	"errors"
)

var (
	ErrNoService    = errors.New("no internal service provided")
	ErrNoRepository = errors.New("no internal repository provided")
)
