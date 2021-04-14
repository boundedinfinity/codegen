package model

import "errors"

var (
	CannotBeEmptyErr = errors.New("cannot be empty")
	NotFoundErr      = errors.New("not found")
)
