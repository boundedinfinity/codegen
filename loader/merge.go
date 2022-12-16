package loader

import (
	cp "boundedinfinity/codegen/codegen_project"
)

func (t *Loader) MergeProject() error {
	for _, project := range t.projectManager.All {
		if err := t.mergeProject(*project); err != nil {
			return err
		}
	}

	return nil
}

func (t *Loader) mergeProject(project cp.CodeGenProjectProject) error {
	if err := t.mergeInfo(project.Info); err != nil {
		return err
	}

	for k, v := range project.Mappings {
		t.projectManager.Merged.Mappings[k] = v
	}

	for _, file := range project.Schemas {
		t.projectManager.Merged.Schemas = append(t.projectManager.Merged.Schemas, file)
	}

	for name, operation := range project.Operations {
		if t.projectManager.Merged.Operations.Has(name) {
			return cp.ErrCodeGenOperationDuplicatev(name)
		}

		t.projectManager.Merged.Operations[name] = operation
	}

	if project.Templates.Header.Defined() {
		t.projectManager.Merged.Templates.Header = project.Templates.Header
	}

	for _, file := range project.Templates.Files {
		t.projectManager.Merged.Templates.Files = append(t.projectManager.Merged.Templates.Files, file)
	}

	return nil
}

func (t *Loader) mergeInfo(info cp.CodeGenProjectInfo) error {
	if info.Description.Defined() {
		t.projectManager.Merged.Info.Description = info.Description
	}

	if info.Namespace.Defined() {
		t.projectManager.Merged.Info.Namespace = info.Namespace
	} else {
		// TODO
	}

	if info.DestDir.Defined() {
		t.projectManager.Merged.Info.DestDir = info.DestDir
	} else {
		// TODO
	}

	if info.FormatSource.Defined() {
		t.projectManager.Merged.Info.FormatSource = info.FormatSource
	} else {
		// TODO
	}

	if info.TemplateDump.Defined() {
		t.projectManager.Merged.Info.TemplateDump = info.TemplateDump
	} else {
		// TODO
	}

	return nil
}
