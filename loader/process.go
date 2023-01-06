package loader

import ct "boundedinfinity/codegen/codegen_type"

func (t *Loader) ProcessTemplates() error {
	var typePaths []string

	err := ct.Walker().TemplateType(func(project *ct.CodeGenProject, template *ct.CodeGenProjectTemplates, file *ct.CodeGenProjectTypeTemplateFile) error {
		if file.SourcePath.Defined() {
			typePaths = append(typePaths, file.SourcePath.Get())
		}

		return nil
	}).Walk(t.projectManager.Projects...)

	if err != nil {
		return err
	}

	var operationPaths []string

	err = ct.Walker().TemplateOperation(func(project *ct.CodeGenProject, template *ct.CodeGenProjectTemplates, file *ct.CodeGenProjectTypeTemplateFile) error {
		if file.SourcePath.Defined() {
			operationPaths = append(operationPaths, file.SourcePath.Get())
		}

		return nil
	}).Walk(t.projectManager.Projects...)

	if err != nil {
		return err
	}

	metas, err := t.LoadTemplatePath(typePaths...)

	if err != nil {
		return err
	}

	for _, meta := range metas {
		t.templateManager.Register(&meta)
	}

	return nil
}
