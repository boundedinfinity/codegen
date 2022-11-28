package generator

import (
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

func TemplateManager(v *template_manager.TemplateManager) Arg {
	return func(t *Generator) {
		t.tm = v
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
		return fmt.Errorf("missing template manager")
	}

	return nil
}
