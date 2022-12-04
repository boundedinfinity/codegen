package template_manager

import (
	"boundedinfinity/codegen/template_type"
)

func (t *TemplateManager) Find(typ template_type.TemplateType) []TemplateContext {
	if t.modelMap.Has(string(typ)) {
		return t.modelMap.Get(string(typ)).Get()
	} else {
		return make([]TemplateContext, 0)
	}
}
