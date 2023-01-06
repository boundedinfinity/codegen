package loader

import (
	ct "boundedinfinity/codegen/codegen_type"
	"boundedinfinity/codegen/codegen_type/template_type"
	"boundedinfinity/codegen/util"
	"path/filepath"

	"github.com/boundedinfinity/go-commoner/environmenter"
	o "github.com/boundedinfinity/go-commoner/optioner"
	"github.com/boundedinfinity/go-commoner/pather"
	"github.com/boundedinfinity/go-commoner/slicer"
	"github.com/boundedinfinity/go-mimetyper/file_extention"
)

func (t *Loader) LoadTemplatePath(paths ...string) ([]ct.TemplateMeta, error) {
	var templateMetas []ct.TemplateMeta

	paths = slicer.Map(paths, environmenter.Sub)
	paths = slicer.Map(paths, filepath.Clean)

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

	if try := util.GetOutputType(templateMeta.SourcePath.Get()); try.Failure() {
		return templateMeta, try.Error
	} else {
		templateMeta.OutputMimeType = try.Result
	}

	if tt, err := template_type.FromUrl(templateMeta.SourcePath.Get()); err != nil {
		return templateMeta, err
	} else {
		templateMeta.TemplateType = tt
	}

	typ := util.GetSchemaTypeId(templateMeta.SourcePath.Get())

	if typ.Defined() {
		templateMeta.Type = typ.Get()
	}

	if err := t.renderer.Load(&templateMeta); err != nil {
		return templateMeta, err
	}

	return templateMeta, nil
}
