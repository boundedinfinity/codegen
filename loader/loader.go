package loader

import (
	"boundedinfinity/codegen/loader_context"
	"boundedinfinity/codegen/renderer"

	"github.com/boundedinfinity/go-jsonschema"
)

type Loader struct {
	jsonSchemas     *jsonschema.System
	typeManager     *loader_context.CodeGenTypeManager
	projectManager  *loader_context.CodeGenProjectManager
	templateManager *loader_context.CodeGenTemplateManager
	renderer        *renderer.Renderer
	// cacher         *cacher.Cacher
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
