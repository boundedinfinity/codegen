package loader

import (
	"boundedinfinity/codegen/errutil"
)

func (t *Loader) Errorf(format string, a ...interface{}) error {
	return errutil.Errorf(t.reportStack.S(), format, a...)
}

func (t *Loader) ErrCannotBeEmpty() error {
	return errutil.CannotBeEmpty(t.reportStack.S())
}

func (t *Loader) ErrNotFound() error {
	return errutil.NotFound(t.reportStack.S())
}

func (t *Loader) ErrCustomTypeDuplicate(v string) error {
	return errutil.Errorf(t.reportStack.S(), "duplicate %v", v)
}

func (t *Loader) ErrCustomTypeNotFound(v string) error {
	return errutil.Errorf(t.reportStack.S(), "type not found %v", v)
}

func (t *Loader) ErrDuplicateType(v string) error {
	return errutil.Errorf(t.reportStack.S(), "duplicate type %v", v)
}

func (t *Loader) ErrDuplicatePrimitive(v string) error {
	return errutil.Errorf(t.reportStack.S(), "duplicate primitive %v", v)
}

func (t *Loader) ErrInvalidPrimitive(v string) error {
	return errutil.Errorf(t.reportStack.S(), "invalid primitive %v", v)
}

func (t *Loader) ErrInvalidType(v string) error {
	return errutil.InvalidType(t.reportStack.S(), v)
}

func (t *Loader) ErrMustBeOneOf(oneOf []string) error {
	return errutil.MustBeOneOf(t.reportStack.S(), oneOf)
}
