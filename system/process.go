package system

func (t *System) Process(paths ...string) error {
	if err := t.loader.LoadTypePaths(paths...); err != nil {
		return err
	}

	if err := t.loader.Process(); err != nil {
		return err
	}

	if err := t.loader.MergeProjects(); err != nil {
		return err
	}

	if err := t.loader.Validate(); err != nil {
		return err
	}

	if err := t.generator.Process(); err != nil {
		return err
	}

	if err := t.generator.Generate(); err != nil {
		return err
	}

	return nil
}
