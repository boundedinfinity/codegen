package loader

func (t *Loader) ProcessTypes() error {
	var schemaPaths []string

	for _, lc := range t.projectManager.Projects {
		for _, file := range lc.Types {
			if file.Source().SourcePath.Defined() {
				schemaPaths = append(schemaPaths, file.Source().SourcePath.Get())
			}
		}
	}

	if err := t.LoadProjectPaths(schemaPaths...); err != nil {
		return err
	}

	return nil
}

func (t *Loader) ProcessTemplates() error {
	var templatePaths []string

	for _, lc := range t.projectManager.Projects {
		for _, file := range lc.Templates.Files {
			if file.SourcePath.Defined() {
				templatePaths = append(templatePaths, file.SourcePath.Get())
			}
		}
	}

	if err := t.LoadTemplatePaths(templatePaths...); err != nil {
		return err
	}

	return nil
}
