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

func (t *Loader) LoadTemplatePaths(paths ...string) error {
	paths = slicer.Map(paths, environmenter.Sub)
	paths = slicer.Map(paths, filepath.Clean)

	for _, path := range paths {
		ok, err := pather.IsFile(path)

		if err != nil {
			return err
		}

		if ok {
			lci := ct.SourceMeta{
				RootPath:   o.Some(pather.Dir(path)),
				SourcePath: o.Some(path),
			}

			if err := t.LoadTemplatePath(lci); err != nil {
				return err
			}

			continue
		}

		sources, err := pather.GetFiles(path)

		if err != nil {
			return err
		}

		for _, source := range sources {
			lci := ct.SourceMeta{
				RootPath:   o.Some(path),
				SourcePath: o.Some(source),
			}

			if err := t.LoadTemplatePath(lci); err != nil {
				return err
			}
		}
	}

	return nil
}

func (t *Loader) LoadTemplatePath(lci ct.SourceMeta) error {
	lc := ct.TemplateContext{
		FileInfo: lci,
	}

	if mt, err := file_extention.FromPath(lc.FileInfo.SourcePath.Get()); err != nil {
		return err
	} else {
		lc.FileInfo.SourceMimeType = mt
	}

	if try := util.GetOutputType(lc.FileInfo.SourcePath.Get()); try.Failure() {
		return try.Error
	} else {
		lc.OutputMimeType = try.Result
	}

	if tt, err := template_type.FromUrl(lc.FileInfo.SourcePath.Get()); err != nil {
		return err
	} else {
		lc.TemplateType = tt
	}

	typeId := util.GetSchemaTypeId(lc.FileInfo.SourcePath.Get())

	if typeId.Defined() {
		lc.TypeId = typeId.Get()
	}

	if err := t.renderer.Load(&lc); err != nil {
		return err
	}

	t.templateManager.Register(&lc)

	return nil
}
