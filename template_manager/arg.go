package template_manager

import (
	"boundedinfinity/codegen/cacher"
	"fmt"
)

type Arg func(*TemplateManager)

func Cacher(v *cacher.Cacher) Arg {
	return func(t *TemplateManager) {
		t.cacher = v
	}
}

func (t *TemplateManager) init() error {
	if t.cacher == nil {
		return fmt.Errorf("cacher is nil")
	}

	return nil
}
