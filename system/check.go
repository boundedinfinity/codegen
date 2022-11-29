package system

import (
	"boundedinfinity/codegen/model"
	"path/filepath"

	o "github.com/boundedinfinity/go-commoner/optioner"
)

func (t *System) Check() error {
	if err := t.jsonSchemas.Check(); err != nil {
		return err
	}

	for schemaPath, schema := range t.pathMap {
		if err := t.mergeSchema(schemaPath, schema); err != nil {
			return err
		}
	}

	if err := t.tm.Register(t.combined.Templates); err != nil {
		return err
	}

	return nil
}

func (t *System) mergeSchema(schemaPath string, schema model.CodeGenSchema) error {
	for language, mappings := range schema.Mappings {
		if err := t.mergeMapping(language, *mappings); err != nil {
			return err
		}
	}

	for name, operation := range schema.Operations {
		if t.combined.Operations.Has(name) {
			return model.ErrCodeGenOperationDuplicatev(name)
		}

		t.combined.Operations[name] = operation
	}

	if err := t.mergeTemplate(schemaPath, schema.Templates); err != nil {
		return err
	}

	return nil
}

func (t *System) mergeTemplate(schemaPath string, templates model.CodeGenSchemaTemplates) error {
	dir := filepath.Dir(schemaPath)

	if templates.Header.Defined() {
		t.combined.Templates.Header = templates.Header
	}

	for _, file := range templates.Files {
		if err := t.cacher.CacheWithBase("templates", o.Some(dir), file.Path.Get()); err != nil {
			return err
		}
	}

	// for _, file := range templates.Files {
	// 	filter := func(f model.CodeGenSchemaTemplateFile) bool {
	// 		return f.Path.Get() == file.Path.Get()
	// 	}

	// 	if slicer.ContainsFn(t.combined.Templates.Files, filter) {
	// 		return model.ErrCodeGenTemplateFilePathDuplicatev(file.Path.Get())
	// 	}

	// 	t.combined.Templates.Files = append(t.combined.Templates.Files, file)
	// }

	return nil
}

func (t *System) mergeMapping(language string, b model.CodeGenSchemaMappings) error {
	a := t.combined.Mappings.Get(language)

	if a.Empty() {
		a = o.Some(model.NewMappings())
		t.combined.Mappings[language] = a.Get()
	}

	switch {
	case a.Get().Package.Defined() && b.Package.Defined():
		if a.Get().Package.Get() == b.Package.Get() {
			return model.ErrCodeGenMappingsPackageDuplicatev(b.Package)
		}
	case b.Package.Defined():
		a.Get().Package = b.Package
	}

	switch {
	case a.Get().RootDir.Defined() && b.RootDir.Defined():
		return model.ErrCodeGenMappingsRootDirDuplicatev(b.RootDir)
	case b.RootDir.Defined():
		abs, err := filepath.Abs(b.RootDir.Get())

		if err != nil {
			return err

		}
		a.Get().RootDir = o.Some(abs)
	}

	for from, to := range b.Replace {
		switch {
		case a.Get().Replace.Has(from):
			return model.ErrCodeGenMappingsRelpaceDuplicatev(from)
		default:
			a.Get().Replace[from] = to
		}
	}

	return nil
}
