package loader

import (
	"boundedinfinity/codegen/loader_context"
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

	if t.templateManager == nil {
		return errors.New("templateManager is nil")
	}

	return nil
}

type Arg func(*Loader)

func TemplateManager(v *loader_context.CodeGenTemplateManager) Arg {
	return func(t *Loader) {
		t.templateManager = v
	}
}

func Canonicals(v *loader_context.CodeGenTypeManager) Arg {
	return func(t *Loader) {
		t.typeManager = v
	}
}

func ProjectManager(v *loader_context.CodeGenProjectManager) Arg {
	return func(t *Loader) {
		t.projectManager = v
	}
}

// func Cacher(v *cacher.Cacher) Arg {
// 	return func(t *Loader) {
// 		t.cacher = v
// 	}
// }
