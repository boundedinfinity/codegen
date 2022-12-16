package template_manager

import (
	"boundedinfinity/codegen/codegen_project"
	"boundedinfinity/codegen/codegen_type"
	"boundedinfinity/codegen/template_delimiter"
	"fmt"

	"github.com/boundedinfinity/go-commoner/optioner"
)

func (t *TemplateManager) init() error {
	// if t.cacher == nil {
	// 	return fmt.Errorf("cacher is nil")
	// }

	if t.pathMap == nil {
		return fmt.Errorf("pathMap is nil")
	}

	if t.typeManager == nil {
		return fmt.Errorf("typeManager is nil")
	}

	if t.projectManager == nil {
		return fmt.Errorf("projectManager is nil")
	}

	if err := t.initTemplatesFuncs(); err != nil {
		return nil
	}

	if t.projectManager.Merged.Info.Delimiter.Empty() {
		t.projectManager.Merged.Info.Delimiter = optioner.Some(template_delimiter.Square)
	}

	return nil
}

type Arg func(*TemplateManager)

func TypeManaager(v *codegen_type.CodeGenTypeManager) Arg {
	return func(t *TemplateManager) {
		t.typeManager = v
	}
}

func ProjectManager(v *codegen_project.CodeGenProjectManager) Arg {
	return func(t *TemplateManager) {
		t.projectManager = v
	}
}

// func Cacher(v *cacher.Cacher) Arg {
// 	return func(t *TemplateManager) {
// 		t.cacher = v
// 	}
// }

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
