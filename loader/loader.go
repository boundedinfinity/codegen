package loader

import (
	"boundedinfinity/codegen/codegen_project"
	"boundedinfinity/codegen/codegen_type"

	"github.com/boundedinfinity/go-jsonschema"
)

type Loader struct {
	jsonSchemas *jsonschema.System
	// cacher         *cacher.Cacher
	typeManager    *codegen_type.CodeGenTypeManager
	projectManager *codegen_project.CodeGenProjectManager
}

func New(args ...Arg) (*Loader, error) {
	t := &Loader{}

	for _, arg := range args {
		arg(t)
	}

	if err := t.init(); err != nil {
		return nil, err
	}

	return t, nil
}
