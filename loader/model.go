package loader

import (
	"boundedinfinity/codegen/model"
	"path"
)

func (t Loader) processModel1(si int, input model.BiInput_Model) error {
	t.reportStack.Push("model[%v]", si)
	defer t.reportStack.Pop()

	ns := t.currentNamespace()
	name := path.Join(ns, input.Name)

	t.modelMap[name] = &model.BiOutput_Model{
		Name:         input.Name,
		Namespace:    ns,
		Description:  t.splitDescription(input.Description),
		Imports:      make([]string, 0),
		JsonStruture: make(map[string]interface{}),
		Properties:   make([]model.BiOutput_Property, 0),
		Templates:    make([]model.BiOutput_Template, 0),
	}

	return nil
}

func (t Loader) processModel2(si int, input model.BiInput_Model) error {
	t.reportStack.Push("model[%v]", si)
	defer t.reportStack.Pop()

	ns := t.currentNamespace()
	name := path.Join(ns, input.Name)
	output, ok := t.modelMap[name]

	if !ok {
		t.CustomTypeNotFound(name)
	}

	if input.Properties != nil {
		for i, inputProperty := range input.Properties {
			var outputProperty model.BiOutput_Property

			if err := t.processProperty(i, inputProperty, &outputProperty); err != nil {
				return err
			}

			output.Properties = append(output.Properties, outputProperty)
		}
	}

	return nil
}

func (t Loader) processModel3(si int, input model.BiInput_Model, output *model.BiOutput_Model) error {
	t.reportStack.Push("model[%v]", si)
	defer t.reportStack.Pop()

	ns := t.currentNamespace()
	output.Name = input.Name
	output.Description = t.splitDescription(input.Description)
	output.Namespace = ns
	output.Imports = make([]string, 0)
	output.Properties = make([]model.BiOutput_Property, 0)
	output.Templates = make([]model.BiOutput_Template, 0)

	// if input.Properties != nil {
	// 	for i, ip := range input.Properties {
	// 		t.reportStack.Push("properties[%v]", i)
	// 		tf, ok := t.getMappedType(ip.Type)

	// 		if !ok {
	// 			return t.NotFound()
	// 		}

	// 		op := model.BiOutput_Property{
	// 			Name:        ip.Name,
	// 			Description: t.splitDescription(ip.Description),
	// 			Namespace:   tf.Namespace,
	// 			Validations: make([]model.BiOutput_Validation, 0),
	// 		}

	// 		if ip.Validations != nil {
	// 			for _, val := range ip.Validations {
	// 				op.Validations = append(op.Validations, model.BiOutput_Validation{
	// 					Minimum:  val.Minimum,
	// 					Maximum:  val.Maximum,
	// 					Required: val.Required,
	// 				})
	// 			}
	// 		}

	// 		if err := t.propertyJsonStructure(ip, &op); err != nil {
	// 			return err
	// 		}

	// 		if ns == op.Namespace {
	// 			op.Type = tf.BaseName
	// 		} else {
	// 			op.Type = tf.ImportName
	// 		}

	// 		if op.Namespace != ns && op.Namespace != model.NAMESPACE_BUILTIN {
	// 			output.Imports = append(output.Imports, op.Namespace)
	// 		}

	// 		output.Imports = util.StrSliceDedup(output.Imports)
	// 		output.Properties = append(output.Properties, op)
	// 		t.reportStack.Pop()
	// 	}
	// }

	tmpls, err := t.getTemplates(ns, model.TemplateType_MODEL)

	if err != nil {
		return err
	}

	for _, itmpl := range tmpls {
		otmpl, err := t.processTemplate2(ns, input.Name, itmpl)

		if err != nil {
			return err
		}

		output.Templates = append(output.Templates, otmpl)
	}

	if err := t.modelJsonStructure(input, output); err != nil {
		return nil
	}

	return nil
}
