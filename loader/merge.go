package loader

import (
	"boundedinfinity/codegen/loader_context"
	"fmt"
)

func alreadyMerged(name string) {
	fmt.Printf("already merged %v\n", name)
}

func (t *Loader) MergeProjects() error {
	for _, lc := range t.projectManager.All {
		if err := t.MergeProject(lc); err != nil {
			return err
		}
	}

	return nil
}

func (t *Loader) MergeProject(lc *loader_context.ProjectLoaderContext) error {
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

	for k, v := range lc.Project.Mappings {
		t.projectManager.Merged.Mappings[k] = v
	}

	for _, file := range lc.Project.Schemas {
		t.projectManager.Merged.Schemas = append(t.projectManager.Merged.Schemas, file)
	}

	for _, operation := range lc.Project.Operations {
		// if t.projectManager.Merged.Operations.Has(name) {
		// 	return cp.ErrCodeGenOperationDuplicatev(name)
		// }

		t.projectManager.Merged.Operations = append(t.projectManager.Merged.Operations, operation)
	}

	if lc.Project.Templates.Header.Defined() {
		t.projectManager.Merged.Templates.Header = lc.Project.Templates.Header
	}

	for _, file := range lc.Project.Templates.Files {
		t.projectManager.Merged.Templates.Files = append(t.projectManager.Merged.Templates.Files, file)
	}

	return nil
}
