package system

import (
	"boundedinfinity/codegen/cacher"
	"boundedinfinity/codegen/model"
	"path/filepath"
	"strings"

	o "github.com/boundedinfinity/go-commoner/optioner"
	"github.com/boundedinfinity/go-commoner/optioner/mapper"
	"github.com/boundedinfinity/go-commoner/slicer"
	"github.com/boundedinfinity/go-urischemer"
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

	return nil
}

func (t *System) mergeSchema(schemaPath string, schema model.CodeGenSchema) error {
	if err := t.mergeInfo(schema.Info); err != nil {
		return err
	}

	if err := t.mergeMapping(schema.Mappings); err != nil {
		return err
	}

	for name, operation := range schema.Operations {
		if t.mergedCodeGen.Operations.Has(name) {
			return model.ErrCodeGenOperationDuplicatev(name)
		}

		t.mergedCodeGen.Operations[name] = operation
	}

	if err := t.mergeTemplate(schemaPath, schema.Templates); err != nil {
		return err
	}

	return nil
}

func (t *System) mergeInfo(info model.CodeGenSchemaInfo) error {
	if info.Description.Defined() {
		t.mergedCodeGen.Info.Description = info.Description
	}

	if info.Package.Defined() {
		t.mergedCodeGen.Info.Package = info.Package
	} else {
		// TODO
	}

	if info.RootDir.Defined() {
		t.mergedCodeGen.Info.RootDir = info.RootDir
	} else {
		// TODO
	}

	return nil
}

func (t *System) mergeTemplate(schemaPath string, templates model.CodeGenSchemaTemplates) error {
	if templates.Header.Defined() {
		t.mergedCodeGen.Templates.Header = templates.Header
	}

	dir := filepath.Dir(schemaPath)
	files := make([]model.CodeGenSchemaTemplateFile, 0)

	for _, file := range templates.Files {
		if file.Path.Defined() {
			scheme, path, err := urischemer.Break(file.Path.Get())

			if err != nil {
				return err
			}

			if scheme == urischemer.File && !filepath.IsAbs(path) {
				rel := path
				path = filepath.Join(dir, rel)
				path, err = filepath.Abs(path)

				if err != nil {
					return err
				}

				path = urischemer.Combine(scheme, path)
				file.Path = o.Some(path)
			}
		}

		files = append(files, file)

		if err := t.cacher.Cache("templates", file.Path.Get()); err != nil {
			return err
		}
	}

	cached := t.cacher.FindByGroup("templates")

	if cached.Empty() {
		return nil
	}

	for _, file := range files {
		cds := slicer.Filter(cached.Get(), func(cd *cacher.CachedData) bool {
			return cd.SourceUrl == file.Path.Get() || strings.HasPrefix(cd.SourceUrl, file.Path.Get())
		})

		for _, cd := range cds {
			t.mergedCodeGen.Templates.Files = append(t.mergedCodeGen.Templates.Files, model.CodeGenSchemaTemplateFile{
				Header: file.Header,
				Path:   o.Some(cd.DestPath),
			})
		}
	}

	return nil
}

func (t *System) mergeMapping(mappings mapper.Mapper[string, string]) error {
	for k, v := range mappings {
		if t.mergedCodeGen.Mappings.Has(k) {
			// TODO
		}

		t.mergedCodeGen.Mappings[k] = v
	}

	return nil
}
