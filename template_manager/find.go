package template_manager

import (
	"boundedinfinity/codegen/template_type"

	"github.com/boundedinfinity/go-commoner/slicer"
)

func (t *TemplateManager) Find(typ template_type.TemplateType) []TemplateContext {
	return slicer.Filter(t.pathMap.Values().Get(), func(x TemplateContext) bool {
		return x.TemplateType == typ
	})
}
