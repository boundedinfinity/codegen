package loader

import (
	"boundedinfinity/codegen/model"
	"fmt"
	"path"
	"strings"
)

func (t *Loader) modelProcessor1(namespace model.BiOutput_Namespace, input model.BiInput_Model, output *model.BiOutput_Model) error {
	output.Name = input.Name
	output.Description = t.splitDescription(input.Description)
	output.Namespace = namespace.Namespace
	out := fmt.Sprintf("%v.%v", path.Base(namespace.Namespace), input.Name)

	t.typeMap[output.SpecName] = &model.TypeInfo{
		SpecType:           output.SpecName,
		InNamespaceType:    input.Name,
		OutOfNamespaceType: out,
		Namespace:          namespace.Namespace,
	}

	return nil
}

func (t *Loader) modelProcessor8(namespace model.BiOutput_Namespace, input model.BiInput_Model, output *model.BiOutput_Model) error {
	if output.Properties == nil {
		return nil
	}

	m := make(map[string]bool)

	for _, property := range output.Properties {
		if !strings.HasPrefix(property.Namespace, model.NAMESPACE_BUILTIN) && property.Namespace != namespace.Namespace {
			if _, ok := m[property.Namespace]; !ok {
				m[property.Namespace] = true
			}
		}
	}

	for k := range m {
		output.Imports = append(output.Imports, k)
	}

	return nil
}
