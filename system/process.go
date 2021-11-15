package system

import "boundedinfinity/codegen/model"

func (t *System) Process() error {
	for _, schema := range t.codeGen {
		if err := t.process1(schema); err != nil {
			return err
		}
	}

	return nil
}

func (t *System) process1(schema *model.Schema) error {
	for _, m := range schema.Models {
		if err := t.jsonSchema.AddtoMap(m); err != nil {
			return err
		}
	}

	return nil
}
