package loader

import (
	"boundedinfinity/codegen/codegen_type"
	"boundedinfinity/codegen/codegen_type/template_delimiter"

	o "github.com/boundedinfinity/go-commoner/optioner"
)

func (t *Loader) MergeProjects() error {
	for _, lc := range t.projectManager.Projects {
		if err := t.MergeProject(lc); err != nil {
			return err
		}
	}

	if t.projectManager.Merged.Info.Delimiter.Empty() {
		t.projectManager.Merged.Info.Delimiter = o.Some(template_delimiter.Square)
	}

	return nil
}

func (t *Loader) MergeProject(project *codegen_type.CodeGenProject) error {
	merged := t.projectManager.Merged

	if project.Info.Description.Defined() {
		merged.Info.Description = project.Info.Description
	}

	if project.Info.Namespace.Defined() {
		merged.Info.Namespace = project.Info.Namespace
	}

	if project.Info.DestDir.Defined() && project.Info.DestDir.Get() != "" {
		merged.Info.DestDir = project.Info.DestDir
	}

	if project.Info.FormatSource.Defined() {
		merged.Info.FormatSource = project.Info.FormatSource
	}

	if project.Info.TemplateDump.Defined() {
		merged.Info.TemplateDump = project.Info.TemplateDump
	}

	if project.Info.Delimiter.Defined() {
		merged.Info.Delimiter = project.Info.Delimiter
	}

	for k, v := range project.Mappings {
		t.projectManager.Merged.Mappings[k] = v
	}

	t.projectManager.Merged.Types = append(
		t.projectManager.Merged.Types,
		project.Types...,
	)

	t.projectManager.Merged.Operations = append(
		t.projectManager.Merged.Operations,
		project.Operations...,
	)

	if project.Templates.Header.Defined() {
		t.projectManager.Merged.Templates.Header = project.Templates.Header
	}

	t.projectManager.Merged.Templates.Files = append(
		t.projectManager.Merged.Templates.Files,
		project.Templates.Files...,
	)

	return nil
}
