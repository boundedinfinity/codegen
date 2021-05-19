package loader

import (
	"boundedinfinity/codegen/errutil"
	"boundedinfinity/codegen/model"
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

func (t *Loader) ErrTemplatePathNotFound(v string) error {
	return errutil.Errorf(t.reportStack.S(), "template path not found %v", v)
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

func (t *Loader) ErrorDuplicateOperation(v string) error {
	return errutil.Errorf(t.reportStack.S(), "duplicate operation %v", v)
}

func (t *Loader) ErrInvalidOperation(v string) error {
	return errutil.Errorf(t.reportStack.S(), "invalide operation %v", v)
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

func (t *Loader) ErrInvalidModel(v string) error {
	return errutil.InvalidType(t.reportStack.S(), v)
}

func (t *Loader) ErrInvalidSource(v model.InputSourceEnum) error {
	return errutil.Errorf(t.reportStack.S(), "invalid source type %v", v)
}

func (t *Loader) ErrMustBeOneOf(oneOf []string) error {
	return errutil.MustBeOneOf(t.reportStack.S(), oneOf)
}
