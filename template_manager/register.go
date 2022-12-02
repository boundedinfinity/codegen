package template_manager

import (
	"boundedinfinity/codegen/model"
	"boundedinfinity/codegen/template_type"
	"fmt"
	"io/ioutil"
	"text/template"

	"github.com/boundedinfinity/go-commoner/pather"
	"github.com/boundedinfinity/go-mimetyper/file_extention"
	"github.com/boundedinfinity/go-mimetyper/mime_type"
)

func (t *TemplateManager) Register(templates model.CodeGenSchemaTemplates) error {
	for _, file := range templates.Files {
		if err := t.registerFileFile(file); err != nil {
			return err
		}
	}

	return nil
}

func (t *TemplateManager) registerFileFile(file model.CodeGenSchemaTemplateFile) error {
	if file.Path.Defined() {
		cdata := t.cacher.Find(file.Path.Get())

		if cdata.Empty() {
			return fmt.Errorf("cache not found for %v", file.Path.Get())
		}

		if t.pathMap.Has(cdata.Get().DestPath) {
			return model.ErrCodeGenTemplateFilePathDuplicatev(file.Path.Get())
		} else {
			path := cdata.Get().DestPath
			ext := pather.Ext(path)
			mt, err := file_extention.GetMimeType(ext)

			if err != nil {
				return err
			}

			bs, err := ioutil.ReadFile(cdata.Get().DestPath)

			if err != nil {
				return err
			}

			template, err := template.
				New("").
				Funcs(t.funcs).Parse(string(bs))

			if err != nil {
				return err
			}

			tt, err := template_type.FromUrl(cdata.Get().DestPath)

			if err != nil {
				return err
			}

			t.pathMap[file.Path.Get()] = TemplateContext{
				TemplateMimeType: mt,
				TemplateType:     tt,
				Template:         template,
				OutputMimeType:   mime_type.ApplicationXGo,
				Path:             cdata.Get().DestPath,
			}
		}
	}

	// if file.Content.Defined() {
	// 	t.pathMap[file.Path.Get()] = []byte(file.Content.Get())
	// }

	return nil
}
