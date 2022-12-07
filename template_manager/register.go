package template_manager

import (
	"boundedinfinity/codegen/model"
	"boundedinfinity/codegen/template_type"
	"fmt"
	"io/ioutil"
	"text/template"

	"github.com/boundedinfinity/go-commoner/extentioner"
	"github.com/boundedinfinity/go-mimetyper/file_extention"
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
			tmplExt := extentioner.Ext(path)
			outputExt := extentioner.Ext(extentioner.Strip(path))
			tmt, err := file_extention.GetMimeType(tmplExt)

			if err != nil {
				return err
			}

			omt, err := file_extention.GetMimeType(outputExt)

			if err != nil {
				return err
			}

			bs, err := ioutil.ReadFile(cdata.Get().DestPath)

			if err != nil {
				return err
			}

			tmpl, err := template.New("").Funcs(t.funcs).Parse(string(bs))

			if err != nil {
				return err
			}

			tt, err := template_type.FromUrl(cdata.Get().DestPath)

			if err != nil {
				return err
			}

			_, err = t.combinedTemplates.AddParseTree(cdata.Get().DestPath, tmpl.Tree)

			if err != nil {
				return err
			}

			tc := TemplateContext{
				TemplateMimeType: tmt,
				TemplateType:     tt,
				OutputMimeType:   omt,
				Path:             cdata.Get().DestPath,
			}

			t.pathMap[file.Path.Get()] = tc
			t.AppendTemplateContext(tc)
		}
	}

	return nil
}
