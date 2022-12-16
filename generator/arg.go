package generator

import (
	"boundedinfinity/codegen/codegen_project"
	"boundedinfinity/codegen/codegen_type"
	"boundedinfinity/codegen/loader"
	"boundedinfinity/codegen/template_manager"
	"fmt"
	"io/fs"
)

const (
	DEFAULT_FILE_MODE = fs.FileMode(0644)
)

type Arg func(*Generator)

func GenExt(v string) Arg {
	return func(t *Generator) {
		t.genExt = v
	}
}

func TypeManager(v *codegen_type.CodeGenTypeManager) Arg {
	return func(t *Generator) {
		t.typeManager = v
	}
}

func TemplateManager(v *template_manager.TemplateManager) Arg {
	return func(t *Generator) {
		t.tm = v
	}
}

func ProjectManager(v *codegen_project.CodeGenProjectManager) Arg {
	return func(t *Generator) {
		t.projectManager = v
	}
}

func FileMode(v fs.FileMode) Arg {
	return func(t *Generator) {
		t.fileMode = v
	}
}

func Loader(v *loader.Loader) Arg {
	return func(t *Generator) {
		t.loader = v
	}
}

const (
	DEFAULT_GENEXT  = "gen"
	DEFAULT_DESTDIR = "/tmp/codegen"
)

func (t *Generator) init() error {
	if t.genExt == "" {
		t.genExt = DEFAULT_GENEXT
	}

	if t.tm == nil {
		return fmt.Errorf("template manager is nil")
	}

	if t.typeManager == nil {
		return fmt.Errorf("codeGenTypeManager is nil")
	}

	if t.projectManager == nil {
		return fmt.Errorf("codeGenSchema is nil")
	}

	if t.loader == nil {
		return fmt.Errorf("loader is nil")
	}

	if t.fileMode == 0 {
		t.fileMode = DEFAULT_FILE_MODE
	}

	return nil
}
