package loader

import (
	"boundedinfinity/codegen/codegen_type"
	"boundedinfinity/codegen/template_delimiter"

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

func (t *Loader) MergeProject(lc *codegen_type.ProjectContext) error {
	merged := t.projectManager.Merged

	if lc.Project.Info.Description.Defined() {
		merged.Info.Description = lc.Project.Info.Description
	}

	if lc.Project.Info.Namespace.Defined() {
		merged.Info.Namespace = lc.Project.Info.Namespace
	}

	if lc.Project.Info.DestDir.Defined() && lc.Project.Info.DestDir.Get() != "" {
		merged.Info.DestDir = lc.Project.Info.DestDir
	}

	if lc.Project.Info.FormatSource.Defined() {
		merged.Info.FormatSource = lc.Project.Info.FormatSource
	}

	if lc.Project.Info.TemplateDump.Defined() {
		merged.Info.TemplateDump = lc.Project.Info.TemplateDump
	}

	if lc.Project.Info.Delimiter.Defined() {
		merged.Info.Delimiter = lc.Project.Info.Delimiter
	}

	for k, v := range lc.Project.Mappings {
		t.projectManager.Merged.Mappings[k] = v
	}

	t.projectManager.Merged.Schemas = append(
		t.projectManager.Merged.Schemas,
		lc.Project.Schemas...,
	)

	t.projectManager.Merged.Operations = append(
		t.projectManager.Merged.Operations,
		lc.Project.Operations...,
	)

	if lc.Project.Templates.Header.Defined() {
		t.projectManager.Merged.Templates.Header = lc.Project.Templates.Header
	}

	t.projectManager.Merged.Templates.Files = append(
		t.projectManager.Merged.Templates.Files,
		lc.Project.Templates.Files...,
	)

	return nil
}
