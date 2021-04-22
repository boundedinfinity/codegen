package loader

import (
	"boundedinfinity/codegen/model"
	"fmt"
	"path"
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
