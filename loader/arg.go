package loader

import (
	"boundedinfinity/codegen/codegen_project"
	"boundedinfinity/codegen/codegen_type"
	"errors"

	"github.com/boundedinfinity/go-jsonschema"
)

func (t *Loader) init() error {
	t.jsonSchemas = jsonschema.New()

	// if t.cacher == nil {
	// 	return errors.New("cacher is nil")
	// }

	if t.typeManager == nil {
		return errors.New("typeManager is nil")
	}

	if t.projectManager == nil {
		return errors.New("projectManager is nil")
	}

	return nil
}

type Arg func(*Loader)

// func Cacher(v *cacher.Cacher) Arg {
// 	return func(t *Loader) {
// 		t.cacher = v
// 	}
// }

func Canonicals(v *codegen_type.CodeGenTypeManager) Arg {
	return func(t *Loader) {
		t.typeManager = v
	}
}

func ProjectManager(v *codegen_project.CodeGenProjectManager) Arg {
	return func(t *Loader) {
		t.projectManager = v
	}
}
