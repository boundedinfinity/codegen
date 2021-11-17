package system

import (
	"boundedinfinity/codegen/model"
	"boundedinfinity/codegen/template_type"
	"boundedinfinity/codegen/util"
	"io/fs"
	"path/filepath"
)

func (t *System) Process() error {
	for _, schema := range t.codeGen {
		if err := t.process1(schema); err != nil {
			return err
		}
	}

	for _, schema := range t.codeGen {
		if err := t.process2(schema); err != nil {
			return err
		}
	}

	for _, schema := range t.codeGen {
		if err := t.process3(schema); err != nil {
			return err
		}
	}

	if err := t.process4(); err != nil {
		return err
	}

	return nil
}

func (t *System) process1(schema *model.Schema) error {
	for _, v := range schema.Models {
		if err := t.jsonSchema.Add(v); err != nil {
			return err
		}
	}

	return nil
}

func (t *System) process2(schema *model.Schema) error {
	for _, v := range schema.Operations {
		if err := t.jsonSchema.Resolve(v.Input); err != nil {
			return err
		}

		if err := t.jsonSchema.Resolve(v.Output); err != nil {
			return err
		}
	}

	return nil
}

func (t *System) process3(schema *model.Schema) error {
	fn := func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		localUri := util.Path2Uri(path)

		if _, ok := t.template[localUri]; ok {
			return model.ErrTemplateDuplicateV(localUri)
		}

		return nil
	}

	for _, v := range schema.Templates.Files {
		if v.Name == "" {
			return model.ErrTemplateEmpty
		}

		path := util.Uri2Path(v.Name)
		file, err := util.IsFile(path)

		if err != nil {
			return err
		}

		if file {
			if _, ok := t.template[v.Name]; ok {
				return model.ErrTemplateDuplicateV(v.Name)
			}
		} else {
			if err := filepath.WalkDir(path, fn); err != nil {
				return err
			}
		}
	}

	return nil
}

func (t *System) process4() error {
	for _, v := range t.template {
		var typ template_type.TemplateType

		if err := t.detectTemplateType(v.Name, &typ); err != nil {
			return err
		}
	}

	return nil
}
