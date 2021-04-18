package model

import (
	"errors"
	"fmt"
	"strings"
)

var (
	CannotBeEmptyErr = errors.New("cannot be empty")
	NotFoundErr      = errors.New("not found")
	InvalidateType   = func(vs []string) error {
		str := strings.Join(vs, ",")
		return fmt.Errorf("invalid value.  Can be one of %v", str)
	}
)
