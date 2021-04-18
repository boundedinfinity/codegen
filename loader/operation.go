package loader

import (
	"boundedinfinity/codegen/model"
	"boundedinfinity/codegen/util"
	"fmt"
)

func (t Loader) processOperation1(si int, m model.BiInput_Operation) error {
	t.reportStack.Push("operation[%v]", si)

	t.addUserMappedType(m.Name)

	t.reportStack.Pop()
	return nil
}

func (t Loader) processOperation2(si int, input model.BiInput_Operation) (model.BiOutput_Operation, error) {
	t.reportStack.Push("operation[%v]", si)

	ns := t.currentNamespace()
	output := model.BiOutput_Operation{
		Name:        input.Name,
		Description: input.Description,
		Namespace:   ns,
		Imports:     make([]string, 0),
		Templates:   make([]model.BiOutput_Template, 0),
	}

	{
		t.reportStack.Push("input")
		inputType, ok := t.getMappedType(input.Input.Type)

		if !ok {
			return output, fmt.Errorf("%v %w", input.Input, model.NotFoundErr)
		}

		output.Input = model.BiOutput_TypeProperty{
			Name:      input.Input.Name,
			Namespace: inputType.Namespace,
			Type:      inputType.BaseName,
		}

		if inputType.Namespace == ns {
			output.Input.Type = inputType.BaseName
		} else {
			output.Input.Type = inputType.ImportName

			if !inputType.BuiltIn {
				output.Imports = append(output.Imports, inputType.Namespace)
			}
		}

		t.reportStack.Pop()
	}

	{
		t.reportStack.Push("output")
		outputType, ok := t.getMappedType(input.Output.Type)

		if !ok {
			return output, fmt.Errorf("%v %w", input.Output, model.NotFoundErr)
		}

		output.Output = model.BiOutput_TypeProperty{
			Name:      input.Output.Name,
			Namespace: outputType.Namespace,
			Type:      outputType.BaseName,
		}

		if outputType.Namespace == ns {
			output.Input.Type = outputType.BaseName
		} else {
			output.Input.Type = outputType.ImportName

			if !outputType.BuiltIn {
				output.Imports = append(output.Imports, outputType.Namespace)
			}
		}

		output.Imports = util.StrSliceDedup(output.Imports)
		t.reportStack.Pop()
	}

	tmpls, err := t.getTemplates(ns, model.TemplateType_OPERATION)

	if err != nil {
		return output, err
	}

	for _, itmpl := range tmpls {
		otmpl, err := t.processTemplate2(ns, output.Name, itmpl)

		if err != nil {
			return output, err
		}

		output.Templates = append(output.Templates, otmpl)
	}

	t.reportStack.Pop()
	return output, nil
}
