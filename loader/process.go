package loader

func (t *Loader) ProcessTypes() error {
	var schemaPaths []string

	for _, lc := range t.projectManager.Projects {
		for _, file := range lc.Project.Types {
			if file.Source().SourcePath.Defined() {
				schemaPaths = append(schemaPaths, file.Source().SourcePath.Get())
			}
		}
	}

	if err := t.LoadTypePaths(schemaPaths...); err != nil {
		return err
	}

	return nil
}

func (t *Loader) ProcessTemplates() error {
	var templatePaths []string

	for _, lc := range t.projectManager.Projects {
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
