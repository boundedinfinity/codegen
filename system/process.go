package system

func (t *System) Load(uris ...string) error {
	if err := t.loader.LoadUri(uris...); err != nil {
		return err
	}

	if err := t.loader.Merge(); err != nil {
		return err
	}

	if err := t.loader.Validate(); err != nil {
		return err
	}

	return nil
}

func (t *System) ProcessTemplates() error {
	if err := t.tm.Register(t.mergedCodeGen.Templates); err != nil {
		return err
	}

	return nil
}

func (t *System) Generate() error {
	if err := t.generator.Process(); err != nil {
		return err
	}

	for _, schema := range t.canonicals.All() {
		if err := t.generator.GenerateModel(schema); err != nil {
			return err
		}
	}

	// for _, operation := range t.mergedCodeGen.Operations {
	// 	if operation.Input.Defined() {
	// 		schema := t.jsonSchemas.Get(string(operation.Input.Get()))

	// 		switch {
	// 		case schema.Defined():
	// 			if err := t.generator.GenerateJsonSchema(schema.Get()); err != nil {
	// 				return err
	// 			}
	// 		default:
	// 			// TODO
	// 		}
	// 	}
	// }

	return nil
}
