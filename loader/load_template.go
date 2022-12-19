package loader

import (
	lc "boundedinfinity/codegen/loader_context"
	"boundedinfinity/codegen/template_type"
	"boundedinfinity/codegen/util"
	"path/filepath"

	"github.com/boundedinfinity/go-commoner/environmenter"
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
			lci := lc.LoaderFileInfo{
				Root:   pather.Dir(path),
				Source: path,
				IsFile: true,
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
			lci := lc.LoaderFileInfo{
				Root:   path,
				Source: source,
				IsFile: true,
			}

			if err := t.LoadTemplatePath(lci); err != nil {
				return err
			}
		}
	}

	return nil
}

func (t *Loader) LoadTemplatePath(lci lc.LoaderFileInfo) error {
	lc := lc.TemplateLoaderContext{
		FileInfo: lci,
	}

	if mt, err := file_extention.FromPath(lc.FileInfo.Source); err != nil {
		return err
	} else {
		lc.FileInfo.MimeType = mt
	}

	if try := util.GetOutputType(lc.FileInfo.Source); try.Failure() {
		return try.Error
	} else {
		lc.OutputMimeType = try.Result
	}

	if tt, err := template_type.FromUrl(lc.FileInfo.Source); err != nil {
		return err
	} else {
		lc.TemplateType = tt
	}

	typeId := util.GetSchemaTypeId(lc.FileInfo.Source)

	if typeId.Defined() {
		lc.TypeId = typeId.Get()
	}

	if err := t.renderer.Load(&lc); err != nil {
		return err
	}

	t.templateManager.Register(&lc)

	return nil
}
