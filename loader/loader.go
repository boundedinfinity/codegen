package loader

import (
	"boundedinfinity/codegen/cacher"
	"boundedinfinity/codegen/canonical"
	"boundedinfinity/codegen/model"
	"errors"

	"github.com/boundedinfinity/go-commoner/optioner/mapper"
	"github.com/boundedinfinity/go-jsonschema"
)

type Loader struct {
	jsonSchemas      *jsonschema.System
	cacher           *cacher.Cacher
	cgsPathMap       mapper.Mapper[string, model.CodeGenSchema]
	canonicalPathMap mapper.Mapper[string, canonical.Canonical]
	canonicals       *canonical.CanonicalCombined
	mergedCodeGen    *model.CodeGenSchema
}

func New(args ...Arg) (*Loader, error) {
	t := &Loader{
		cgsPathMap:       mapper.Mapper[string, model.CodeGenSchema]{},
		canonicalPathMap: mapper.Mapper[string, canonical.Canonical]{},
	}

	for _, arg := range args {
		arg(t)
	}

	if err := t.init(); err != nil {
		return nil, err
	}

	return t, nil
}

func (t *Loader) init() error {
	if t.jsonSchemas == nil {
		return errors.New("jsonSchemas is nil")
	}

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
