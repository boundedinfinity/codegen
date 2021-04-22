package loader

import (
	"boundedinfinity/codegen/model"
	"fmt"
	"path"
)

func (t Loader) namespaceProcssor1(input model.BiInput_Namespace, output *model.BiOutput_Namespace) error {
	if input.Namespaces != nil {
		for _, child := range input.Namespaces {
			output.Children = append(output.Children, child.Name)
		}
	}

	pns := path.Dir(output.Namespace)

	if tmpls, ok := t.templateMap[pns]; ok {
		for _, tmpl := range tmpls {
			t.templateMap[output.Namespace] = append(t.templateMap[output.Namespace], tmpl)
		}
	}

	if input.Templates != nil {
		for tmplIndex, tmpl := range input.Templates {
			templateWrapper := func() error {
				t.reportStack.Push("template[%v]", tmplIndex)
				defer t.reportStack.Pop()

				var processed model.BiInput_Template

				if err := t.processTemplate1(*output, tmpl, &processed); err != nil {
					return err
				}

				t.templateMap[output.Namespace] = append(t.templateMap[output.Namespace], processed)

				return nil
			}

			if err := templateWrapper(); err != nil {
				return err
			}
		}
	}

	return nil
}

func (t Loader) namespaceProcssor7(input model.BiInput_Namespace, output *model.BiOutput_Namespace) error {
	if tmpls, ok := t.templateMap[output.Namespace]; ok {
		for _, v := range tmpls {
			switch v.Type {
			case string(model.TemplateType_NAMESPACE):
				if _, ok := t.namespaceMap[output.Namespace]; ok {
					fmt.Printf("namespace: %v\n", v)
				}
			case string(model.TemplateType_MODEL):
				for _, m := range t.modelMap {
					if m.Namespace == output.Namespace {
						for _, inputTemplate := range tmpls {
							outputTemplate := model.New_BiOutput_Template()

							if err := t.processTemplate2(*output, m.Name, inputTemplate, outputTemplate); err != nil {
								return err
							}

							m.Templates = append(m.Templates, outputTemplate)
						}
					}
				}
			case string(model.TemplateType_OPERATION):
				for _, o := range t.operationMap {
					if o.Namespace == output.Namespace {
						fmt.Printf("operation: %v\n", v)
					}
				}
			default:
				return t.InvalidateType()
			}
		}
	}

	return nil
}
