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

	if info, ok := t.typeMap[inputBuiltin]; ok {
		output.Input.SpecType = inputBuiltin
		output.Input.Namespace = info.Namespace
		output.Input.Type = info.InNamespaceType
	} else if info, ok := t.typeMap[inputAbs]; ok {
		output.Input.SpecType = inputAbs
		output.Input.Namespace = info.Namespace

		if namespace.Namespace == inputAbs {
			output.Input.Type = info.InNamespaceType
		} else {
			output.Input.Type = info.OutOfNamespaceType
		}
	} else if info, ok := t.typeMap[inputRel]; ok {
		output.Input.SpecType = inputRel
		output.Input.Namespace = info.Namespace

		if namespace.Namespace == inputRel {
			output.Input.Type = info.InNamespaceType
		} else {
			output.Input.Type = info.OutOfNamespaceType
		}
	} else {
		return t.NotFound()
	}

	outputBuiltin := path.Join(model.NAMESPACE_BUILTIN, input.Output.Type)
	outputAbs := path.Join(t.rootNamespace(), input.Output.Type)
	outputRel := path.Join(namespace.Namespace, input.Output.Type)

	if info, ok := t.typeMap[outputBuiltin]; ok {
		output.Output.SpecType = outputBuiltin
		output.Output.Namespace = info.Namespace
		output.Output.Type = info.InNamespaceType
	} else if info, ok := t.typeMap[outputAbs]; ok {
		output.Output.SpecType = outputAbs
		output.Output.Namespace = info.Namespace

		if namespace.Namespace == outputAbs {
			output.Output.Type = info.InNamespaceType
		} else {
			output.Output.Type = info.OutOfNamespaceType
		}
	} else if info, ok := t.typeMap[outputRel]; ok {
		output.Output.SpecType = outputRel
		output.Namespace = info.Namespace

		if namespace.Namespace == outputRel {
			output.Output.Type = info.InNamespaceType
		} else {
			output.Output.Type = info.OutOfNamespaceType
		}
	} else {
		return t.NotFound()
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
