package loader

import (
	"boundedinfinity/codegen/model"
)

func (t *Loader) modelProcessor1(namespace model.BiOutput_Namespace, input model.BiInput_Model, output *model.BiOutput_Model) error {
	output.Name = input.Name
	output.Namespace = t.currentNamespace()

	return nil
}
