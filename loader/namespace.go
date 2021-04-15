package loader

import (
	"boundedinfinity/codegen/model"
	"strings"
)

func (t *Loader) processNamespace(input model.BiSpecNamespace, tmpls []model.BiSpecTemplate) (model.BiGenNamespace, error) {
	t.stack.Push(input.Name)
	output := model.BiGenNamespace{
		Name:       strings.Join(t.stack.S(), "/"),
		Types:      make([]model.BiGenType, 0),
		Namespaces: make([]model.BiGenNamespace, 0),
	}

	if input.Types != nil {
		for _, styp := range input.Types {
			name := strings.Join([]string{output.Name, styp.Name}, "/")

			gtyp := model.BiGenType{
				Name:       name,
				Type:       styp.Type,
				Properties: make([]model.BiGenTypeProperty, 0),
				Templates:  make([]model.BiGenTemplate, 0),
			}

			if styp.Properties != nil {
				for _, sprop := range styp.Properties {
					gtyp.Properties = append(gtyp.Properties, model.BiGenTypeProperty{
						Name: sprop.Name,
						Type: sprop.Type,
					})
				}
			}

			if tmpls != nil {
				for _, stmpl := range tmpls {
					gtmpl, err := t.processTemplate(name, stmpl)

					if err != nil {
						return output, err
					}

					gtyp.Templates = append(gtyp.Templates, gtmpl)
				}
			}

			output.Types = append(output.Types, gtyp)
			t.Gen.Lookup[name] = gtyp
		}
	}

	if input.Namespaces != nil {
		for _, ns := range input.Namespaces {
			if gen, err := t.processNamespace(ns, tmpls); err != nil {
				return gen, err
			} else {
				output.Namespaces = append(output.Namespaces, gen)
			}
		}
	}

	t.stack.Pop()
	return output, nil
}
