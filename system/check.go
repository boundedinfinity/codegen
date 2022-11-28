package system

import (
	"boundedinfinity/codegen/model"
	"path/filepath"

	"github.com/boundedinfinity/go-commoner/optioner"
)

func (t *System) Check() error {
	if err := t.jsonSchemas.Check(); err != nil {
		return err
	}

	for _, schema := range t.pathMap {
		if err := t.mergeSchema(schema); err != nil {
			return err
		}
	}

	return nil
}

func (t *System) mergeSchema(schema model.CodeGenSchema) error {
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

	// for x := range schema.Templates {

	// }

	return nil
}

func (t *System) mergeTemplate(tmpl model.CodeGenSchemaTemplateFile) error {
	return nil
}

func (t *System) mergeMapping(language string, b model.CodeGenSchemaMappings) error {
	a := t.combined.Mappings.Get(language)

	if a.Empty() {
		a = optioner.Some(model.NewMappings())
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
		a.Get().RootDir = optioner.Some(abs)
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
