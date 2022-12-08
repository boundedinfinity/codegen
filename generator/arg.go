package generator

import (
	"boundedinfinity/codegen/canonical"
	"boundedinfinity/codegen/model"
	"boundedinfinity/codegen/template_manager"
	"fmt"
)

type Arg func(*Generator)

func DestDir(v string) Arg {
	return func(t *Generator) {
		t.destDir = v
	}
}

func GenExt(v string) Arg {
	return func(t *Generator) {
		t.genExt = v
	}
}

func Canonicals(v *canonical.CanonicalCombined) Arg {
	return func(t *Generator) {
		t.canonicals = v
	}
}

func TemplateManager(v *template_manager.TemplateManager) Arg {
	return func(t *Generator) {
		t.tm = v
	}
}

func CodeGenSchema(v *model.CodeGenSchema) Arg {
	return func(t *Generator) {
		t.codeGenSchema = v
	}
}

const (
	DEFAULT_GENEXT  = "gen"
	DEFAULT_DESTDIR = "/tmp/codegen"
)

func (t *Generator) init() error {
	if t.destDir == "" {
		t.destDir = DEFAULT_DESTDIR
	}

	if t.genExt == "" {
		t.genExt = DEFAULT_GENEXT
	}

	if t.tm == nil {
		return fmt.Errorf("template manager is nil")
	}

	if t.canonicals == nil {
		return fmt.Errorf("canonicals is nil")
	}

	if t.codeGenSchema == nil {
		return fmt.Errorf("codeGenSchema is nil")
	}

	return nil
}
