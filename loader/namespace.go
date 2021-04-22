package loader

import (
	"boundedinfinity/codegen/model"
)

func (t Loader) namespaceProcssor1(input model.BiInput_Namespace, output *model.BiOutput_Namespace) error {
	if input.Namespaces != nil {
		for _, child := range input.Namespaces {
			output.Children = append(output.Children, child.Name)
		}
	}

	return nil
}
