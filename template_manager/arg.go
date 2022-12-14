package template_manager

import (
	"boundedinfinity/codegen/cacher"
	"boundedinfinity/codegen/canonical"
	"boundedinfinity/codegen/model"
	"boundedinfinity/codegen/template_delimiter"
	"fmt"

	"github.com/boundedinfinity/go-commoner/optioner"
)

func (t *TemplateManager) init() error {
	if t.cacher == nil {
		return fmt.Errorf("cacher is nil")
	}

	if t.pathMap == nil {
		return fmt.Errorf("pathMap is nil")
	}

	if t.canonicals == nil {
		return fmt.Errorf("canonicals is nil")
	}

	if err := t.initTemplatesFuncs(); err != nil {
		return nil
	}

	if t.codeGenSchema.Info.Delimiter.Empty() {
		t.codeGenSchema.Info.Delimiter = optioner.Some(template_delimiter.Square)
	}

	return nil
}

type Arg func(*TemplateManager)

func CanonicalCombined(v *canonical.CanonicalCombined) Arg {
	return func(t *TemplateManager) {
		t.canonicals = v
	}
}

func CodeGenSchema(v *model.CodeGenSchema) Arg {
	return func(t *TemplateManager) {
		t.codeGenSchema = v
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

func Verbose(v bool) Arg {
	return func(t *TemplateManager) {
		t.verbose = v
	}
}
