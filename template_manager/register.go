package template_manager

import (
	"boundedinfinity/codegen/model"
	"fmt"
	"io/ioutil"

	o "github.com/boundedinfinity/go-commoner/optioner"
	"github.com/boundedinfinity/go-commoner/pather"
)

func (t *TemplateManager) Register(templates ...model.CodeGenSchemaTemplates) error {
	for _, template := range templates {
		if err := t.registerTemplate(template); err != nil {
			return err
		}
	}

	return nil
}

func (t *TemplateManager) registerTemplate(template model.CodeGenSchemaTemplates) error {
	for _, file := range template.Files {
		if err := t.registerFilePath(file); err != nil {
			return err
		}
	}

	return nil
}

func (t *TemplateManager) registerFilePath(file model.CodeGenSchemaTemplateFile) error {
	if file.Path.Defined() {
		if ok, err := pather.IsDir(file.Path.Get()); err != nil {
			return err
		} else {
			if ok {
				if err := t.registerFileDir(file); err != nil {
					return err
				}
			} else {
				if err := t.registerFileFile(o.None[string](), file); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (t *TemplateManager) registerFileDir(file model.CodeGenSchemaTemplateFile) error {
	if file.Path.Defined() {
		if paths, err := pather.GetPaths(file.Path.Get()); err != nil {
			return err
		} else {
			for _, path := range paths {
				new := model.CodeGenSchemaTemplateFile{
					Header: file.Header,
					Path:   o.Some(path),
				}

				if err := t.registerFilePath(new); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (t *TemplateManager) registerFileFile(base o.Option[string], file model.CodeGenSchemaTemplateFile) error {
	if file.Path.Defined() {
		err := t.cacher.Cache("templates", file.Path.Get())

		if err != nil {
			return err
		}

		cdata := t.cacher.Find(file.Path.Get())

		if cdata.Empty() {
			return fmt.Errorf("cache not found for %v", file.Path.Get())
		}

		if t.pathMap.Has(cdata.Get().DestPath) {
			return model.ErrCodeGenTemplateFilePathDuplicatev(file.Path.Get())
		} else {
			data, err := ioutil.ReadFile(cdata.Get().DestPath)

			if err != nil {
				return err
			}

			t.pathMap[file.Path.Get()] = data
		}
	}

	return nil
}
