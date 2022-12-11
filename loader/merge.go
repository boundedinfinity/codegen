package loader

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

func (t *Loader) Merge() error {
	for _, jsSchema := range t.jsonSchemas.AllPath() {
		if schema, err := t.convertJsonSchema(jsSchema, o.None[string]()); err != nil {
			return err
		} else {
			source := t.jsonSchemas.GetSource(jsSchema.GetId().Get())

			if err := t.canonicals.Register(schema, source); err != nil {
				return err
			}
		}
	}

	for source, schema := range t.canonicalPathMap {
		if err := t.canonicals.Register(schema, o.Some(source)); err != nil {
			return err
		}
	}

	for schemaPath, schema := range t.cgsPathMap {
		if err := t.mergeSchema(schemaPath, schema); err != nil {
			return err
		}
	}

	return nil
}

func (t *Loader) mergeSchema(schemaPath string, schema model.CodeGenSchema) error {
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

	if schema.Templates.Header.Defined() {
		t.mergedCodeGen.Templates.Header = schema.Templates.Header
	}

	for _, file := range schema.Templates.Files {
		if err := t.mergeTemplates(schemaPath, file); err != nil {
			return err
		}
	}

	return nil
}

func (t *Loader) mergeTemplates(schemaPath string, file model.CodeGenSchemaTemplateFile) error {
	if file.Path.Empty() {
		return nil
	}

	dir := filepath.Dir(schemaPath)
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

	if err := t.cacher.Cache("templates", file.Path.Get()); err != nil {
		return err
	}

	cached := t.cacher.FindByGroup("templates")

	if cached.Empty() {
		return nil
	}

	cds := slicer.Filter(cached.Get(), func(cd *cacher.CachedData) bool {
		return cd.SourceUrl == file.Path.Get() || strings.HasPrefix(cd.SourceUrl, file.Path.Get())
	})

	for _, cd := range cds {
		newFile := model.CodeGenSchemaTemplateFile{
			Header: file.Header,
			Path:   o.Some(cd.DestPath),
		}

		t.mergedCodeGen.Templates.Files = append(t.mergedCodeGen.Templates.Files, newFile)
	}

	return nil
}

func (t *Loader) mergeMapping(mappings mapper.Mapper[string, string]) error {
	for k, v := range mappings {
		if t.mergedCodeGen.Mappings.Has(k) {
			// TODO
		}

		t.mergedCodeGen.Mappings[k] = v
	}

	return nil
}

func (t *Loader) mergeInfo(info model.CodeGenSchemaInfo) error {
	if info.Description.Defined() {
		t.mergedCodeGen.Info.Description = info.Description
	}

	if info.Namespace.Defined() {
		t.mergedCodeGen.Info.Namespace = info.Namespace
	} else {
		// TODO
	}

	if info.DestDir.Defined() {
		t.mergedCodeGen.Info.DestDir = info.DestDir
	} else {
		// TODO
	}

	return nil
}
