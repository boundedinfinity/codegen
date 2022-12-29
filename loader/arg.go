package loader

import (
	"boundedinfinity/codegen/manager"
	"boundedinfinity/codegen/renderer"
	"errors"

	"github.com/boundedinfinity/go-jsonschema"
)

func (t *Loader) init() error {
	t.jsonSchemas = jsonschema.New()

	// if t.cacher == nil {
	// 	return errors.New("cacher is nil")
	// }

	if t.renderer == nil {
		return errors.New("renderer is nil")
	}

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

func TemplateManager(v *manager.CodeGenTemplateManager) Arg {
	return func(t *Loader) {
		t.templateManager = v
	}
}

func TypeManager(v *manager.CodeGenTypeManager) Arg {
	return func(t *Loader) {
		t.typeManager = v
	}
}

func ProjectManager(v *manager.CodeGenProjectManager) Arg {
	return func(t *Loader) {
		t.projectManager = v
	}
}

func Renderer(v *renderer.Renderer) Arg {
	return func(t *Loader) {
		t.renderer = v
	}
}

// func Cacher(v *cacher.Cacher) Arg {
// 	return func(t *Loader) {
// 		t.cacher = v
// 	}
// }
