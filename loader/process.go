package loader

func (t *Loader) Process() error {

	for _, project := range t.projectManager.All {
		for _, file := range project.Schemas {
			if file.Path.Defined() {
				if err := t.LoadTypePaths(file.Path.Get()); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
