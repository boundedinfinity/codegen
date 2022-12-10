package loader

import (
	"boundedinfinity/codegen/cacher"
	"boundedinfinity/codegen/canonical"
	"boundedinfinity/codegen/model"
	"errors"

	"github.com/boundedinfinity/go-jsonschema"
)

func (t *Loader) init() error {
	t.jsonSchemas = jsonschema.New()

	if t.cacher == nil {
		return errors.New("cacher is nil")
	}

	if t.canonicals == nil {
		return errors.New("canonicals is nil")
	}

	if t.mergedCodeGen == nil {
		return errors.New("mergedCodeGen is nil")
	}

	return nil
}

type Arg func(*Loader)

func Cacher(v *cacher.Cacher) Arg {
	return func(t *Loader) {
		t.cacher = v
	}
}

func Canonicals(v *canonical.CanonicalCombined) Arg {
	return func(t *Loader) {
		t.canonicals = v
	}
}

func MergedCodeGen(v *model.CodeGenSchema) Arg {
	return func(t *Loader) {
		t.mergedCodeGen = v
	}
}
