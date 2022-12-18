package loader

import (
	lc "boundedinfinity/codegen/loader_context"
	"boundedinfinity/codegen/template_delimiter"
	"boundedinfinity/codegen/template_type"
	"boundedinfinity/codegen/util"
	"fmt"
	"html/template"
	"io/ioutil"
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
			mt, err := file_extention.FromPath(path)

			if err != nil {
				return err
			}

			lci := lc.LoaderFileInfo{
				Root:     pather.Dir(path),
				Source:   path,
				IsFile:   true,
				MimeType: mt,
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
			mt, err := file_extention.FromPath(path)

			if err != nil {
				return err
			}

			lci := lc.LoaderFileInfo{
				Root:     path,
				Source:   source,
				IsFile:   true,
				MimeType: mt,
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

	if try := util.GetTemplateType(lc.FileInfo.Source); try.Failure() {
		return try.Error
	} else {
		lc.TemplateMimeType = try.Result
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

	bs, err := ioutil.ReadFile(lci.Source)

	if err != nil {
		return err
	}

	l, r := template_delimiter.Get(t.projectManager.Merged.Info.Delimiter.Get())

	if t.projectManager.Merged.Info.TemplateDump.Defined() && t.projectManager.Merged.Info.TemplateDump.Get() {
		tmp := string(bs)
		tmp += fmt.Sprintf("\n\n%v DUMP . %v", l, r)
		bs = []byte(tmp)
	}

	if tmpl, err := template.New("").Funcs(t.funcs).Delims(l, r).Parse(string(bs)); err != nil {
		return err
	} else {
		tc.Template = tmpl
	}

	t.path2project[file.Path.Get()] = tc
	t.AppendTemplateContext(tc)
	util.MapListAdd(t.root2project, file.Root.Get(), tc)

	return nil
}
