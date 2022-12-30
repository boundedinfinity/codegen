package codegen_type

type projectWalker struct {
	infoFn         func(*CodeGenProject, *CodeGenInfo) error
	operationFn    func(*CodeGenProject, *CodeGenProjectOperation) error
	templateFn     func(*CodeGenProject, *CodeGenProjectTemplates, *CodeGenProjectTemplateFile) error
	schemaFn       func(*CodeGenProject, *CodeGenProjectTypeFile) error
	schemaAll      func(*CodeGenProject, *CodeGenProjectTypeFile, CodeGenTypeContext) error
	schemaStringFn func(CodeGenTypeContext, *CodeGenTypeString) error
	schemaArrayFn  func(CodeGenTypeContext, *CodeGenTypeArray) error
	schemaObjectFn func(CodeGenTypeContext, *CodeGenTypeObject) error
}

func Walker() *projectWalker {
	return &projectWalker{}
}

func (w *projectWalker) Walk(projects ...*CodeGenProject) error {
	for _, project := range projects {
		if err := w.walk(project); err != nil {
			return err
		}
	}

	return nil
}

func (w *projectWalker) walk(project *CodeGenProject) error {
	if project == nil {
		return nil
	}

	if w.infoFn != nil {
		if err := w.infoFn(project, &project.Info); err != nil {
			return err
		}
	}

	if w.operationFn != nil {
		for _, operation := range project.Operations {
			if err := w.operationFn(project, operation); err != nil {
				return err
			}
		}
	}

	if w.schemaFn != nil {
		for _, file := range project.Schemas {
			if err := w.schemaFn(project, file); err != nil {
				return err
			}
		}
	}

	if w.templateFn != nil && project.Templates.Files != nil {
		for _, file := range project.Templates.Files {
			if file == nil {
				continue
			}

			if err := w.templateFn(project, &project.Templates, file); err != nil {
				return err
			}
		}
	}

	return nil
}

func (w *projectWalker) Info(fn func(*CodeGenProject, *CodeGenInfo) error) *projectWalker {
	w.infoFn = fn
	return w
}

func (w *projectWalker) Operation(fn func(*CodeGenProject, *CodeGenProjectOperation) error) *projectWalker {
	w.operationFn = fn
	return w
}

func (w *projectWalker) Schema(fn func(*CodeGenProject, *CodeGenProjectTypeFile) error) *projectWalker {
	w.schemaFn = fn
	return w
}

func (t *projectWalker) SchemaAll(v func(*CodeGenProject, *CodeGenProjectTypeFile, CodeGenTypeContext) error) *projectWalker {
	t.schemaAll = v
	return t
}

func (t *projectWalker) SchemaString(v func(CodeGenTypeContext, *CodeGenTypeString) error) *projectWalker {
	t.schemaStringFn = v
	return t
}

func (t *projectWalker) SchemaArray(v func(CodeGenTypeContext, *CodeGenTypeArray) error) *projectWalker {
	t.schemaArrayFn = v
	return t
}

func (t *projectWalker) SchemaObject(v func(CodeGenTypeContext, *CodeGenTypeObject) error) *projectWalker {
	t.schemaObjectFn = v
	return t
}

func (w *projectWalker) Template(fn func(*CodeGenProject, *CodeGenProjectTemplates, *CodeGenProjectTemplateFile) error) *projectWalker {
	w.templateFn = fn
	return w
}
