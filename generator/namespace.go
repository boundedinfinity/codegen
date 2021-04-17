package generator

import (
	"boundedinfinity/codegen/model"
)

func (t *Generator) runNamespace(ns model.BiOutput_Model_Namespace) error {
	for _, typ := range ns.Models {
		for _, tmpl := range typ.Templates {
			ctx := model.BiOutput_TemplateModelContext{
				Model: typ,
				Spec:  t.spec,
			}

			if err := t.renderFile(tmpl.Input, tmpl.Output, ctx); err != nil {
				return err
			}
		}
	}

	for _, childNs := range ns.Namespaces {
		if err := t.runNamespace(childNs); err != nil {
			return err
		}
	}

	return nil
}
