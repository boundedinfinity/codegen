package loader

import "boundedinfinity/codegen/model"

func (t Loader) namespaceProcssor1(input model.BiInput_Namespace, output *model.BiOutput_Namespace) error {
	output.Namespace = t.currentNamespace()
	output.RelativeNamespace = t.relativeNamespace("")

	return nil
}
