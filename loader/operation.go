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

		if _, ok := t.outputModels[inputOperation.Input.Type.String()]; !ok {
			return t.ErrInvalidModel(inputOperation.Input.Name)
		}

		if _, ok := t.outputModels[inputOperation.Output.Type.String()]; !ok {
			return t.ErrInvalidModel(inputOperation.Input.Name)
		}
	}

	return nil
}
