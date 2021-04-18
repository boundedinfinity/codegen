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

func (t Loader) processModel2(si int, im model.BiInput_Model) (model.BiOutput_Model, error) {
	t.reportStack.Push("model[%v]", si)

	ns := t.currentNamespace()
	om := model.BiOutput_Model{
		Name:       im.Name,
		Namespace:  ns,
		Imports:    make([]string, 0),
		Properties: make([]model.BiOutput_TypeProperty, 0),
		Templates:  make([]model.BiOutput_Template, 0),
	}

	if im.Properties != nil {
		for i, ip := range im.Properties {
			t.reportStack.Push("properties[%v]", i)
			tf, ok := t.getMappedType(ip.Type)

			if !ok {
				return om, fmt.Errorf("%v %w", ip.Type, model.NotFoundErr)
			}

			op := model.BiOutput_TypeProperty{
				Name:      tf.BaseName,
				Namespace: tf.Namespace,
			}

			if ns == op.Namespace {
				op.Type = tf.BaseName
			} else {
				op.Type = tf.ImportName
			}

			if op.Namespace != ns && op.Namespace != model.NAMESPACE_BUILTIN {
				om.Imports = append(om.Imports, op.Namespace)
			}

			om.Imports = util.StrSliceDedup(om.Imports)
			om.Properties = append(om.Properties, op)
			t.reportStack.Pop()
		}
	}

	tmpls, err := t.getTemplates(ns, model.TemplateType_MODEL)

	if err != nil {
		return om, err
	}

	for _, itmpl := range tmpls {
		otmpl, err := t.processTemplate(ns, "", itmpl)

		if err != nil {
			return om, err
		}

		om.Templates = append(om.Templates, otmpl)
	}

	t.reportStack.Pop()
	return om, nil
}
