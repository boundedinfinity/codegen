package system

func (t *System) ProcessTemplates() error {
	if err := t.tm.Register(t.combined.Templates); err != nil {
		return err
	}
	return nil
}

func (t *System) Generate() error {
	for _, operation := range t.combined.Operations {
		if operation.Input.Defined() {
			schema := t.jsonSchemas.Get(string(operation.Input.Get()))

			switch {
			case schema.Defined():
				if err := t.generator.GenerateJsonSchema(schema.Get()); err != nil {
					return err
				}
			default:
				// TODO
			}
		}
	}

	return nil
}
