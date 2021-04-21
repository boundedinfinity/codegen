package loader

import (
	"boundedinfinity/codegen/model"
	"fmt"
	"path"
)

func (t *Loader) modelProcessor1(namespace model.BiOutput_Namespace, input model.BiInput_Model, output *model.BiOutput_Model) error {
	output.Name = input.Name
	output.Description = t.splitDescription(input.Description)
	output.Namespace = t.currentNamespace()
	out := fmt.Sprintf("%v.%v", path.Base(t.currentNamespace()), input.Name)

	t.typeMap[output.SpecName] = &model.TypeInfo{
		SpecType:           output.SpecName,
		InNamespaceType:    input.Name,
		OutOfNamespaceType: out,
		Namespace:          t.currentNamespace(),
	}

	return nil
}

func (t *Loader) modelProcessor5(namespace model.BiOutput_Namespace, input model.BiInput_Model, output *model.BiOutput_Model) error {

	return nil
}
