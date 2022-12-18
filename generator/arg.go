package generator

import (
	"boundedinfinity/codegen/loader_context"
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

func TypeManager(v *loader_context.CodeGenTypeManager) Arg {
	return func(t *Generator) {
		t.typeManager = v
	}
}

func TemplateManager(v *loader_context.CodeGenTemplateManager) Arg {
	return func(t *Generator) {
		t.templateManager = v
	}
}

func ProjectManager(v *loader_context.CodeGenProjectManager) Arg {
	return func(t *Generator) {
		t.projectManager = v
	}
}

func FileMode(v fs.FileMode) Arg {
	return func(t *Generator) {
		t.fileMode = v
	}
}

// func Loader(v *loader.Loader) Arg {
// 	return func(t *Generator) {
// 		t.loader = v
// 	}
// }

const (
	DEFAULT_GENEXT  = "gen"
	DEFAULT_DESTDIR = "/tmp/codegen"
)

func (t *Generator) init() error {
	if t.genExt == "" {
		t.genExt = DEFAULT_GENEXT
	}

	if t.templateManager == nil {
		return fmt.Errorf("templateManager is nil")
	}

	if t.typeManager == nil {
		return fmt.Errorf("typeManager is nil")
	}

	if t.projectManager == nil {
		return fmt.Errorf("projectManager is nil")
	}

	if t.fileMode == 0 {
		t.fileMode = DEFAULT_FILE_MODE
	}

	return nil
}
