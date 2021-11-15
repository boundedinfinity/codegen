package system

import "boundedinfinity/codegen/model"

func (t *System) Process() error {
	for _, schema := range t.codeGen {
		if err := t.process1(schema); err != nil {
			return err
		}
	}

	for _, schema := range t.codeGen {
		if err := t.process2(schema); err != nil {
			return err
		}
	}

	return nil
}

func (t *System) process1(schema *model.Schema) error {
	for _, m := range schema.Models {
		if err := t.jsonSchema.Add(m); err != nil {
			return err
		}
	}

	return nil
}

func (t *System) process2(schema *model.Schema) error {
	for _, o := range schema.Operations {
		if err := t.jsonSchema.Resolve(o.Input); err != nil {
			return err
		}

		if err := t.jsonSchema.Resolve(o.Output); err != nil {
			return err
		}
	}

	return nil
}
