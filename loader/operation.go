package loader

import (
	"boundedinfinity/codegen/model"
)

func (t Loader) processOperation1(si int, m model.BiInput_Operation) error {
	t.reportStack.Push("operation[%v]", si)
	defer t.reportStack.Pop()

	// t.addUserMappedType(m.Name)

	return nil
}

func (t Loader) processOperation3(si int, input model.BiInput_Operation, output *model.BiOutput_Operation) error {
	t.reportStack.Push("operation[%v]", si)
	defer t.reportStack.Pop()

	ns := t.currentNamespace()

	output.Name = input.Name
	output.Description = t.splitDescription(input.Description)
	output.Namespace = ns

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
