package loader

import (
	ct "boundedinfinity/codegen/codegen_type"
	"boundedinfinity/codegen/codegen_type/template_delimiter"
	"fmt"

	o "github.com/boundedinfinity/go-commoner/optioner"
)

func (t *Loader) MergeProjects(projects ...*ct.CodeGenProject) error {
	if err := t.projectManager.RegisterProject(projects...); err != nil {
		return err
	}

	merged := t.projectManager.Merged

	err := ct.Walker().
		Type(func(_ *ct.CodeGenProject, typ ct.CodeGenType) error {
			if err := t.typeManager.Register(typ); err != nil {
				return err
			}

			switch typ.(type) {
			case *ct.CodeGenTypePath:
				// ignore
			default:
				merged.Types = append(merged.Types, typ)
			}

			return nil
		}).
		Operation(func(project *ct.CodeGenProject, operation *ct.CodeGenProjectOperation) error {
			if err := t.projectManager.RegisterOperation(*operation); err != nil {
				return err
			}

			merged.Operations = append(merged.Operations, operation)

			return nil
		}).Walk(projects...)

	if err != nil {
		return err
	}

	err = ct.Walker().
		Info(func(project *ct.CodeGenProject, info *ct.CodeGenInfo) error {
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
				if merged.Mappings.Has(k) {
					fmt.Printf("project.Mappings already contains %v", k)
				} else {
					merged.Mappings[k] = v
				}
			}

			t.projectManager.Merged.Templates.Types = append(
				t.projectManager.Merged.Templates.Types,
				project.Templates.Types...,
			)

			return nil
		}).Walk(projects...)

	if merged.Info.Delimiter.Empty() {
		merged.Info.Delimiter = o.Some(template_delimiter.Square)
	}

	// if t.projectManager.Merged.Info.DestDir.Empty() {
	// 	fmt.Printf("implement default dest dir")
	// }

	return nil
}
