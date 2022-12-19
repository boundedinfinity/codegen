package loader

func (t *Loader) ProcessTypes() error {
	var schemaPaths []string

	for _, lc := range t.projectManager.All {
		for _, file := range lc.Project.Schemas {
			if file.Path.Defined() {
				schemaPaths = append(schemaPaths, file.Path.Get())
			}
		}
	}

	if err := t.LoadTypePaths(schemaPaths...); err != nil {
		return err
	}

	return nil
}

func (t *Loader) ProcessTempaltes() error {
	var templatePaths []string

	for _, lc := range t.projectManager.All {
		for _, file := range lc.Project.Templates.Files {
			if file.Path.Defined() {
				templatePaths = append(templatePaths, file.Path.Get())
			}
		}
	}

	if err := t.LoadTemplatePaths(templatePaths...); err != nil {
		return err
	}

	return nil
}
