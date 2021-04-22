package loader

import (
	"boundedinfinity/codegen/errutil"
)

func (t *Loader) Errorf(format string, a ...interface{}) error {
	return errutil.Errorf(t.namespaceStack.S(), format, a...)
}

func (t *Loader) CannotBeEmpty() error {
	return errutil.CannotBeEmpty(t.reportStack.S())
}

func (t *Loader) NotFound() error {
	return errutil.NotFound(t.reportStack.S())
}

func (t *Loader) CustomTypeDuplicate(v string) error {
	return errutil.Errorf(t.reportStack.S(), "duplicate %v", v)
}

func (t *Loader) CustomTypeNotFound(v string) error {
	return errutil.Errorf(t.reportStack.S(), "type not found %v", v)
}

func (t *Loader) DuplicateType(v string) error {
	return errutil.Errorf(t.reportStack.S(), "duplicate type %v", v)
}

func (t *Loader) InvalidateType() error {
	return errutil.InvalidateType(t.reportStack.S())
}

func (t *Loader) MustBeOneOf(oneOf []string) error {
	return errutil.MustBeOneOf(t.reportStack.S(), oneOf)
}
