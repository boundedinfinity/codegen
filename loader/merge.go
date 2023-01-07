package loader

import (
	ct "boundedinfinity/codegen/codegen_type"
	"boundedinfinity/codegen/codegen_type/template_delimiter"
	"fmt"

	o "github.com/boundedinfinity/go-commoner/optioner"
)

func (t *Loader) MergeProjects(projects ...*ct.CodeGenProject) error {
	merged := t.projectManager.Merged

	err := ct.WalkProject(func(project *ct.CodeGenProject) error {
		return t.projectManager.RegisterProject(project)
	}, projects...)

	if err != nil {
		return err
	}

	err = ct.WalkType(func(_ *ct.CodeGenProject, typ ct.CodeGenType) error {
		if _, ok := typ.(*ct.CodeGenTypePath); !ok {
			merged.Types = append(merged.Types, typ)
			return t.typeManager.Register(typ)
		}

		return nil
	}, projects...)

	if err != nil {
		return err
	}

	err = ct.WalkOperation(func(cgp *ct.CodeGenProject, operation *ct.CodeGenProjectOperation) error {
		if operation.Name.Defined() {
			merged.Operations = append(merged.Operations, operation)
			return t.projectManager.RegisterOperation(*operation)
		}

		return nil
	}, projects...)

	if err != nil {
		return err
	}

	err = ct.WalkProject(func(project *ct.CodeGenProject) error {
		return t.projectManager.RegisterProject(project)
	}, projects...)

	err = ct.WalkInfo(func(project *ct.CodeGenProject, info *ct.CodeGenInfo) error {
		if info.Description.Defined() {
			merged.Info.Description = info.Description
		}

		if info.Namespace.Defined() {
			merged.Info.Namespace = info.Namespace
		}

		if info.DestDir.Defined() && info.DestDir.Get() != "" {
			merged.Info.DestDir = info.DestDir
		}

		if info.FormatSource.Defined() {
			merged.Info.FormatSource = info.FormatSource
		}

		if info.TemplateDump.Defined() {
			merged.Info.TemplateDump = info.TemplateDump
		}

		if info.Delimiter.Defined() {
			merged.Info.Delimiter = info.Delimiter
		}

		for k, v := range project.Mappings {
			if merged.Mappings.Has(k) {
				fmt.Printf("project.Mappings already contains %v", k)
			} else {
				merged.Mappings[k] = v
			}
		}

		merged.Templates.Types = append(merged.Templates.Types, project.Templates.Types...)

		return nil
	}, projects...)

	if merged.Info.Delimiter.Empty() {
		merged.Info.Delimiter = o.Some(template_delimiter.Square)
	}

	if merged.Info.Namespace.Empty() {
		merged.Info.Namespace = o.Some("NAMESPACE")
	}

	if t.projectManager.Merged.Info.DestDir.Empty() {
		fmt.Printf("implement default dest dir")
	}

	return nil
}
