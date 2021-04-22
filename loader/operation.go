package loader

import (
	"boundedinfinity/codegen/model"
	"path"
)

func (t Loader) processOperation6(namespace model.BiOutput_Namespace, input model.BiInput_Operation, output *model.BiOutput_Operation) error {
	output.Namespace = namespace.Namespace
	output.Description = t.splitDescription(input.Description)
	output.Input.Name = input.Input.Name
	output.Input.Description = t.splitDescription(input.Input.Description)
	output.Output.Name = input.Output.Name
	output.Output.Description = t.splitDescription(input.Output.Description)

	inputBuiltin := path.Join(model.NAMESPACE_BUILTIN, input.Input.Type)
	inputAbs := path.Join(t.rootNamespace(), input.Input.Type)
	inputRel := path.Join(namespace.Namespace, input.Input.Type)
	var inputInfo *model.TypeInfo

	if info, ok := t.typeMap[inputBuiltin]; ok {
		inputInfo = info
	} else if info, ok := t.typeMap[inputAbs]; ok {
		inputInfo = info
	} else if info, ok := t.typeMap[inputRel]; ok {
		inputInfo = info
	} else {
		return t.NotFound()
	}

	output.Input.SpecName = path.Join(output.SpecName, output.Input.Name)
	output.Input.Namespace = inputInfo.Namespace
	output.Input.SpecType = inputInfo.SpecType

	if namespace.Namespace == inputInfo.Namespace {
		output.Input.Type = inputInfo.InNamespaceType
	} else {
		output.Input.Type = inputInfo.OutOfNamespaceType
	}

	outputBuiltin := path.Join(model.NAMESPACE_BUILTIN, input.Output.Type)
	outputAbs := path.Join(t.rootNamespace(), input.Output.Type)
	outputRel := path.Join(namespace.Namespace, input.Output.Type)
	var outputInfo *model.TypeInfo

	if info, ok := t.typeMap[outputBuiltin]; ok {
		outputInfo = info
	} else if info, ok := t.typeMap[outputAbs]; ok {
		outputInfo = info
	} else if info, ok := t.typeMap[outputRel]; ok {
		outputInfo = info
	} else {
		return t.NotFound()
	}

	output.Output.SpecName = path.Join(output.SpecName, output.Output.Name)
	output.Output.Namespace = outputInfo.Namespace
	output.Output.SpecType = outputInfo.SpecType

	if namespace.Namespace == outputInfo.Namespace {
		output.Output.Type = outputInfo.InNamespaceType
	} else {
		output.Output.Type = outputInfo.OutOfNamespaceType
	}

	return nil
}

func (t Loader) processOperation3(si int, input model.BiInput_Operation, output *model.BiOutput_Operation) error {
	checkInput := func() error {
		// 	t.reportStack.Push("input")
		// 	defer t.reportStack.Pop()

		// 	inputType, ok := t.getMappedType(input.Input.Type)

		// 	if !ok {
		// 		return t.NotFound()
		// 	}

		// 	output.Input = model.BiOutput_Property{
		// 		Name:      input.Input.Name,
		// 		Namespace: inputType.Namespace,
		// 		Type:      inputType.BaseName,
		// 	}

		// 	if inputType.Namespace == ns {
		// 		output.Input.Type = inputType.BaseName
		// 	} else {
		// 		output.Input.Type = inputType.ImportName

		// 		if !inputType.BuiltIn {
		// 			output.Imports = append(output.Imports, inputType.Namespace)
		// 		}
		// 	}

		return nil
	}

	checkOutput := func() error {
		// t.reportStack.Push("output")
		// defer t.reportStack.Pop()

		// outputType, ok := t.getMappedType(input.Output.Type)

		// if !ok {
		// 	return t.NotFound()
		// }

		// output.Output = model.BiOutput_Property{
		// 	Name:      input.Output.Name,
		// 	Namespace: outputType.Namespace,
		// 	Type:      outputType.BaseName,
		// }

		// if outputType.Namespace == ns {
		// 	output.Input.Type = outputType.BaseName
		// } else {
		// 	output.Input.Type = outputType.ImportName

		// 	if !outputType.BuiltIn {
		// 		output.Imports = append(output.Imports, outputType.Namespace)
		// 	}
		// }

		// output.Imports = util.StrSliceDedup(output.Imports)
		return nil
	}

	checkTemplates := func() error {
		// tmpls, err := t.getTemplates(ns, model.TemplateType_OPERATION)

		// if err != nil {
		// 	return err
		// }

		// for _, itmpl := range tmpls {
		// 	otmpl, err := t.processTemplate2(ns, output.Name, itmpl)

		// 	if err != nil {
		// 		return err
		// 	}

		// 	output.Templates = append(output.Templates, otmpl)
		// }
		return nil
	}

	if err := checkInput(); err != nil {
		return err
	}

	if err := checkOutput(); err != nil {
		return err
	}

	if err := checkTemplates(); err != nil {
		return err
	}

	return nil
}
