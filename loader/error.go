package loader

import (
	"boundedinfinity/codegen/errutil"
)

func (t *Loader) Errorf(format string, a ...interface{}) error {
	return errutil.Errorf(t.modelStack.S(), format, a...)
}

func (t *Loader) CannotBeEmpty() error {
	return errutil.CannotBeEmpty(t.modelStack.S())
}

func (t *Loader) NotFound() error {
	return errutil.NotFound(t.modelStack.S())
}

func (t *Loader) CustomTypeNotFound(v string) error {
	return errutil.Errorf(t.modelStack.S(), "type not found %v", v)
}

func (t *Loader) DuplicateType(v string) error {
	return errutil.Errorf(t.modelStack.S(), "duplicate type %v", v)
}

func (t *Loader) InvalidateType() error {
	return errutil.InvalidateType(t.modelStack.S())
}

func (t *Loader) MustBeOneOf(oneOf []string) error {
	return errutil.MustBeOneOf(t.modelStack.S(), oneOf)
}
