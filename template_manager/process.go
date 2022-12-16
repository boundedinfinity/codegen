package template_manager

import (
	cp "boundedinfinity/codegen/codegen_project"

	o "github.com/boundedinfinity/go-commoner/optioner"
	"github.com/boundedinfinity/go-commoner/pather"
)

func (t *TemplateManager) Process() error {
	for _, project := range t.projectManager.All {
		for _, file := range project.Templates.Files {
			if file.Path.Empty() {
				return nil
			}

			ok, err := pather.IsDir(file.Path.Get())

			if err != nil {
				return err
			}

			if !ok {
				return t.Register(o.None[string](), file)
			}

			sources, err := pather.GetFiles(file.Path.Get())

			if err != nil {
				return err
			}

			for _, source := range sources {
				subFile := cp.CodeGenProjectTemplateFile{
					Header:  file.Header,
					Path:    o.Some(source),
					Content: file.Content,
				}

				if err := t.Register(file.Path, subFile); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
