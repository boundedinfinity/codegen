package loader

import "boundedinfinity/codegen/model"

func (t *Loader) processOperation1() error {
	for name, inputOperation := range t.inputOperations {
		if _, ok := t.inputOperations[name]; !ok {
			return t.ErrInvalidOperation(name)
		}

		outputOperation := model.NewOutputOperationWithInput(inputOperation)
		t.outputOperations[name] = outputOperation

		assign := func(i model.InputModel, o **model.OutputModel) error {
			switch i.Type {
			case model.SchemaType_Ref:
				if ref, ok := t.outputModels[i.Ref]; ok {
					nm := model.NewOutputModelWithOutput(ref)
					nm.Name = i.Name
					*o = nm
				} else {
					return t.ErrInvalidModel(i.Ref)
				}
			default:
				return t.ErrInvalidModel(i.Name)
			}

			return nil
		}

		if err := assign(inputOperation.Input, &outputOperation.Input); err != nil {
			return err
		}

		if err := assign(inputOperation.Output, &outputOperation.Output); err != nil {
			return err
		}

		t.OutputSpec.Operations = append(t.OutputSpec.Operations, outputOperation)
	}

	return nil
}
