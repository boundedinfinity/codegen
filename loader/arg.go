package loader

import (
	"boundedinfinity/codegen/cacher"
	"boundedinfinity/codegen/canonical"
	"boundedinfinity/codegen/model"

	"github.com/boundedinfinity/go-jsonschema"
)

type Arg func(*Loader)

func Jsonschemas(v *jsonschema.System) Arg {
	return func(t *Loader) {
		t.jsonSchemas = v
	}
}

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
