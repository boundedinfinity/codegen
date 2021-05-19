package loader

import "boundedinfinity/codegen/model"

func (t *Loader) processOperation1() error {
	for name, inputOperation := range t.inputOperations {
		if _, ok := t.inputOperations[name]; !ok {
			return t.ErrInvalidOperation(name)
		}

		outputOperation := model.NewOutputOperation()
		t.outputOperations[name] = outputOperation
		outputOperation.Description = t.splitDescription(inputOperation.Description)
		outputOperation.Name = inputOperation.Name

		switch inputOperation.Input.Type {
		// case model.SchemaType_Ref:
		// 	if _, ok := t.outputModels[inputOperation.Input.Ref.Ref]; !ok {
		// 		return t.ErrInvalidModel(inputOperation.Input.Name)
		// 	}
		default:
			return t.ErrInvalidModel(inputOperation.Input.Name)
		}

		switch inputOperation.Output.Type {
		// case model.SchemaType_Ref:
		// 	if _, ok := t.outputModels[inputOperation.Output.Ref.Ref]; !ok {
		// 		return t.ErrInvalidModel(inputOperation.Output.Name)
		// 	}
		default:
			return t.ErrInvalidModel(inputOperation.Output.Name)
		}

	}

	return nil
}
