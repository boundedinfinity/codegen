package loader

import (
	"boundedinfinity/codegen/model"
	"boundedinfinity/codegen/util"
	"fmt"
)

func (t Loader) processModel1(si int, m model.BiInput_Model) error {
	t.reportStack.Push("model[%v]", si)

	t.addUserMappedType(m.Name)

	t.reportStack.Pop()
	return nil
}

func (t Loader) processModel2(si int, input model.BiInput_Model) (model.BiOutput_Model, error) {
	t.reportStack.Push("model[%v]", si)

	ns := t.currentNamespace()
	output := model.BiOutput_Model{
		Name:        input.Name,
		Description: t.splitDescription(input.Description),
		Namespace:   ns,
		Imports:     make([]string, 0),
		Properties:  make([]model.BiOutput_TypeProperty, 0),
		Templates:   make([]model.BiOutput_Template, 0),
	}

	if input.Properties != nil {
		for i, ip := range input.Properties {
			t.reportStack.Push("properties[%v]", i)
			tf, ok := t.getMappedType(ip.Type)

			if !ok {
				return output, fmt.Errorf("%v %w", ip.Type, model.NotFoundErr)
			}

			op := model.BiOutput_TypeProperty{
				Name:        ip.Name,
				Description: t.splitDescription(ip.Description),
				Namespace:   tf.Namespace,
				Validations: make([]model.BiOutput_Validation, 0),
			}

			if ip.Validations != nil {
				for _, val := range ip.Validations {
					op.Validations = append(op.Validations, model.BiOutput_Validation{
						Minimum:  val.Minimum,
						Maximum:  val.Maximum,
						Required: val.Required,
					})
				}
			}

			if ns == op.Namespace {
				op.Type = tf.BaseName
			} else {
				op.Type = tf.ImportName
			}

			if op.Namespace != ns && op.Namespace != model.NAMESPACE_BUILTIN {
				output.Imports = append(output.Imports, op.Namespace)
			}

			output.Imports = util.StrSliceDedup(output.Imports)
			output.Properties = append(output.Properties, op)
			t.reportStack.Pop()
		}
	}

	tmpls, err := t.getTemplates(ns, model.TemplateType_MODEL)

	if err != nil {
		return output, err
	}

	for _, itmpl := range tmpls {
		otmpl, err := t.processTemplate2(ns, input.Name, itmpl)

		if err != nil {
			return output, err
		}

		output.Templates = append(output.Templates, otmpl)
	}

	t.reportStack.Pop()
	return output, nil
}
