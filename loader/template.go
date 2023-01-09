package loader

import (
	ct "boundedinfinity/codegen/codegen_type"
	"boundedinfinity/codegen/codegen_type/template_type"
	"boundedinfinity/codegen/util"
	"fmt"

	o "github.com/boundedinfinity/go-commoner/optioner"
	"github.com/boundedinfinity/go-commoner/pather"
	"github.com/boundedinfinity/go-mimetyper/file_extention"
)

func templatePath(project *ct.CodeGenProject, _ *ct.CodeGenProjectTemplates, file *ct.CodeGenProjectTemplateFile) error {
	*file.Source() = project.SourceMeta
	return nil
}

func (t *Loader) ProcessTemplates(projects ...*ct.CodeGenProject) error {
	merged := t.projectManager.Merged

	processTemplate := func(project *ct.CodeGenProject, _ *ct.CodeGenProjectTemplates, file *ct.CodeGenProjectTemplateFile) error {
		if file.SourcePath.Defined() {
			if metas, err := t.LoadTemplatePath(project.RootPath, file.SourcePath.Get()); err != nil {
				return err
			} else {
				for _, meta := range metas {
					if meta.SourcePath.Empty() {
						continue
					}

					new := &ct.CodeGenProjectTemplateFile{
						TemplateMeta: meta,
						Header:       file.Header,
					}

					switch meta.TemplateType {
					case template_type.Model:
						merged.Templates.Types = append(merged.Templates.Types, new)
					case template_type.Operation:
						merged.Templates.Operations = append(merged.Templates.Operations, new)
					default:
						fmt.Printf("template type %v not implemented\n", meta.TemplateType)
					}

					t.templateManager.Register(new)
				}
			}
		}

		return nil
	}

	if err := ct.WalkTemplate(processTemplate, projects...); err != nil {
		return err
	}

	if err := ct.WalkTemplate(templatePath, projects...); err != nil {
		return err
	}

	return nil
}

func (t *Loader) LoadTemplatePath(root o.Option[string], paths ...string) ([]ct.TemplateMeta, error) {
	var templateMetas []ct.TemplateMeta

	paths, err := normalizePath(root, paths...)

	if err != nil {
		return templateMetas, err
	}

	for _, path := range paths {
		ok, err := pather.IsFile(path)

		if err != nil {
			return templateMetas, err
		}

		if ok {
			sourceMeta := ct.SourceMeta{
				RootPath:   o.Some(pather.Dir(path)),
				SourcePath: o.Some(path),
			}

			if mt, err := file_extention.FromPath(sourceMeta.SourcePath.Get()); err != nil {
				return templateMetas, err
			} else {
				sourceMeta.SourceMimeType = mt
			}

			if templateMeta, err := t.loadTemplatePath(sourceMeta); err != nil {
				return templateMetas, err
			} else {
				templateMetas = append(templateMetas, templateMeta)
			}

			continue
		}

		sources, err := pather.GetFiles(path)

		if err != nil {
			return templateMetas, err
		}

		for _, source := range sources {
			sourceMeta := ct.SourceMeta{
				RootPath:   o.Some(path),
				SourcePath: o.Some(source),
			}

			if mt, err := file_extention.FromPath(sourceMeta.SourcePath.Get()); err != nil {
				return templateMetas, err
			} else {
				sourceMeta.SourceMimeType = mt
			}

			if templateMeta, err := t.loadTemplatePath(sourceMeta); err != nil {
				return templateMetas, err
			} else {
				templateMetas = append(templateMetas, templateMeta)
			}
		}
	}

	return templateMetas, nil
}

func (t *Loader) loadTemplatePath(sourceMeta ct.SourceMeta) (ct.TemplateMeta, error) {
	templateMeta := ct.TemplateMeta{
		SourceMeta: sourceMeta,
	}

	if try := util.GetOutputMimeType(templateMeta.SourcePath.Get()); try.Failure() {
		return templateMeta, try.Error
	} else {
		templateMeta.OutputMimeType = try.Result
	}

	if try := util.GetTemplateMimeType(templateMeta.SourcePath.Get()); try.Failure() {
		return templateMeta, try.Error
	} else {
		templateMeta.TemplateMimeTime = try.Result
	}

	if try := util.GetTemplateExt(templateMeta.SourcePath.Get()); try.Failure() {
		return templateMeta, try.Error
	} else {
		templateMeta.TemplateExt = try.Result
	}

	if try := util.GetOutputExt(templateMeta.SourcePath.Get()); try.Failure() {
		return templateMeta, try.Error
	} else {
		templateMeta.OutputExt = try.Result
	}

	if tt, err := template_type.FromUrl(templateMeta.SourcePath.Get()); err != nil {
		return templateMeta, err
	} else {
		templateMeta.TemplateType = tt
	}

	templateMeta.Type = util.GetSchemaTypeId(templateMeta.SourcePath)

	if err := t.renderer.Load(&templateMeta); err != nil {
		return templateMeta, err
	}

	return templateMeta, nil
}
