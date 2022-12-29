package codegen_type

type walker struct {
	infoFn      func(*CodeGenProjectProject, *CodeGenProjectInfo) error
	operationFn func(*CodeGenProjectProject, *CodeGenProjectOperation) error
	schemaFn    func(*CodeGenProjectProject, *CodeGenProjectTypeFile) error
	templateFn  func(*CodeGenProjectProject, *CodeGenProjectTemplates, *CodeGenProjectTemplateFile) error
}

func Walk() *walker {
	return &walker{}
}

func (w *walker) Each(projects []*CodeGenProjectProject) error {
	for _, project := range projects {
		if err := w.Run(project); err != nil {
			return err
		}
	}

	return nil
}

func (w *walker) Run(project *CodeGenProjectProject) error {
	if w.infoFn != nil {
		if err := w.infoFn(project, &project.Info); err != nil {
			return err
		}
	}

	if w.operationFn != nil && project.Operations != nil {
		for _, operation := range project.Operations {
			if err := w.operationFn(project, operation); err != nil {
				return err
			}
		}
	}

	if w.schemaFn != nil && project.Schemas != nil {
		for _, file := range project.Schemas {
			if err := w.schemaFn(project, file); err != nil {
				return err
			}
		}
	}

	if w.templateFn != nil && project.Templates.Files != nil {
		for _, file := range project.Templates.Files {
			if err := w.templateFn(project, &project.Templates, file); err != nil {
				return err
			}
		}
	}

	return nil
}

func (w *walker) Info(fn func(*CodeGenProjectProject, *CodeGenProjectInfo) error) *walker {
	w.infoFn = fn
	return w
}

func (w *walker) Operation(fn func(*CodeGenProjectProject, *CodeGenProjectOperation) error) *walker {
	w.operationFn = fn
	return w
}

func (w *walker) Schema(fn func(*CodeGenProjectProject, *CodeGenProjectTypeFile) error) *walker {
	w.schemaFn = fn
	return w
}

func (w *walker) Template(fn func(*CodeGenProjectProject, *CodeGenProjectTemplates, *CodeGenProjectTemplateFile) error) *walker {
	w.templateFn = fn
	return w
}
