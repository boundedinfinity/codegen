package uri_scheme

import (
	"errors"

	"github.com/boundedinfinity/go-commoner/errorer"
)

var (
	ErrUriTypeInvalid = errors.New("invalid URI scheme")
)

var (
	ErrUriTypeInvalidV = errorer.Errorfn(ErrUriTypeInvalid)
)
