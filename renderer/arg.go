package renderer

import (
	"boundedinfinity/codegen/loader_context"
	"errors"
)

type Arg func(*Renderer)

func (t *Renderer) init(args ...Arg) error {
	if t.typeManager == nil {
		return errors.New("typeManager is nil")
	}

	if t.projectManager == nil {
		return errors.New("projectManager is nil")
	}

	if t.templateManager == nil {
		return errors.New("templateManager is nil")
	}

	if len(t.funcs) == 0 {
		TemplateFunc("DUMP", dumpJson)(t)
		TemplateFunc("PASCAL", t.pascal)(t)
		TemplateFunc("CAMEL", t.camel)(t)
		TemplateFunc("SNAKE", t.camel)(t)
		TemplateFunc("BASE", t.pathBase)(t)
		TemplateFunc("DIR", t.pathDir)(t)
		TemplateFunc("PATH_REL", t.pathRel)(t)
		TemplateFunc("DEFINED", t.defined)(t)
		TemplateFunc("EMPTY", t.empty)(t)
		TemplateFunc("SINGULAR", t.singular)(t)
		TemplateFunc("PLURAL", t.plural)(t)
		TemplateFunc("RES_SCHEMA_NS", t.resolveSchemaNs)(t)
		TemplateFunc("RES_SCHEMA", t.resolveSchema)(t)
		TemplateFunc("RESOLVE", t.resolveSchema)(t)
	}

	return nil
}

func TemplateFunc(name string, fn any) Arg {
	return func(t *Renderer) {
		t.funcs[name] = fn
	}
}

func TemplateManager(v *loader_context.CodeGenTemplateManager) Arg {
	return func(t *Renderer) {
		t.templateManager = v
	}
}

func TypeManager(v *loader_context.CodeGenTypeManager) Arg {
	return func(t *Renderer) {
		t.typeManager = v
	}
}

func ProjectManager(v *loader_context.CodeGenProjectManager) Arg {
	return func(t *Renderer) {
		t.projectManager = v
	}
}
