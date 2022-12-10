package loader

import (
	"boundedinfinity/codegen/cacher"
	"boundedinfinity/codegen/canonical"
	"boundedinfinity/codegen/model"

	"github.com/boundedinfinity/go-commoner/optioner/mapper"
	"github.com/boundedinfinity/go-jsonschema"
)

type Loader struct {
	jsonSchemas      *jsonschema.System
	cacher           *cacher.Cacher
	cgsPathMap       mapper.Mapper[string, model.CodeGenSchema]
	cgsId2path       mapper.Mapper[string, string]
	canonicalPathMap mapper.Mapper[string, canonical.Canonical]
	canonicalId2path mapper.Mapper[string, string]
	canonicals       *canonical.CanonicalCombined
	mergedCodeGen    *model.CodeGenSchema
}

func New(args ...Arg) (*Loader, error) {
	t := &Loader{
		cgsPathMap:       mapper.Mapper[string, model.CodeGenSchema]{},
		cgsId2path:       mapper.Mapper[string, string]{},
		canonicalPathMap: mapper.Mapper[string, canonical.Canonical]{},
		canonicalId2path: mapper.Mapper[string, string]{},
	}

	for _, arg := range args {
		arg(t)
	}

	if err := t.init(); err != nil {
		return nil, err
	}

	return t, nil
}
