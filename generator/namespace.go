package generator

import (
	"boundedinfinity/codegen/model"
)

func (t *Generator) runNamespace(ns model.BiGenNamespace) error {
	for _, typ := range ns.Types {
		for _, tmpl := range typ.Templates {
			ctx := model.BiGenTemplateTypeContext{
				Type: typ,
				Spec: t.spec,
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
