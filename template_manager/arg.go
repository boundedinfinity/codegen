package template_manager

import (
	"boundedinfinity/codegen/cacher"
	"boundedinfinity/codegen/model"

	"github.com/boundedinfinity/go-jsonschema"
)

type Arg func(*TemplateManager)

func CodeGenSchema(v *model.CodeGenSchema) Arg {
	return func(t *TemplateManager) {
		t.codeGenSchema = v
	}
}

func JsonSchema(v *jsonschema.System) Arg {
	return func(t *TemplateManager) {
		t.jsonSchemas = v
	}
}

func Cacher(v *cacher.Cacher) Arg {
	return func(t *TemplateManager) {
		t.cacher = v
	}
}

func TemplateFunc(v string, fn any) Arg {
	return func(t *TemplateManager) {
		t.funcs[v] = fn
	}
}

func SetTemplateFuncs(v map[string]any) Arg {
	return func(t *TemplateManager) {
		t.funcs = v
	}
}

func FormatSource(v bool) Arg {
	return func(t *TemplateManager) {
		t.formatSource = v
	}
}

func Verbose(v bool) Arg {
	return func(t *TemplateManager) {
		t.verbose = v
	}
}
