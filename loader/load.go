package loader

import (
	cp "boundedinfinity/codegen/codegen_project"
	"boundedinfinity/codegen/codegen_type"
	"boundedinfinity/codegen/util"
	"encoding/json"
	"path/filepath"

	o "github.com/boundedinfinity/go-commoner/optioner"
	"github.com/boundedinfinity/go-commoner/slicer"
	"github.com/boundedinfinity/go-jsonschema/model"
	"github.com/boundedinfinity/go-marshaler"
	"github.com/boundedinfinity/go-mimetyper/mime_type"
	"github.com/ghodss/yaml"
)

func (t *Loader) LoadTypePaths(paths ...string) error {
	paths = slicer.Map(paths, filepath.Clean)

	for _, path := range paths {
		m, err := marshaler.ReadFromPath(path)

		if err != nil {
			return err
		}

		for source, content := range m {
			if err := t.LoadTypePath(path, source, content.Data, content.MimeType); err != nil {
				return err
			}
		}
	}

	return nil
}

func (t *Loader) LoadTypePath(root, source string, data []byte, mt mime_type.MimeType) error {
	var bs []byte
	var err error

	switch mt {
	case mime_type.ApplicationXYaml:
		bs, err = yaml.YAMLToJSON(data)

		if err != nil {
			return err
		}
	case mime_type.ApplicationJson:
		bs = data
	default:
		return cp.ErrMimeTypeUnsupportedv(mt)
	}

	switch {
	case util.IsJsonSchemaFile(source):
		if schema, err := model.UnmarshalSchema(bs); err != nil {
			return err
		} else {
			t.jsonSchemas.Register(root, source, schema)
		}
	case util.IsCodeGenSchemaFile(source):
		var schema cp.CodeGenProjectProject

		if err := json.Unmarshal(bs, &schema); err != nil {
			return err
		}

		schema.Info.DestDir = util.ExpandPatho(o.Some(source), schema.Info.DestDir)

		schema.Schemas = slicer.Map(schema.Schemas, func(f cp.CodeGenProjectTypeFile) cp.CodeGenProjectTypeFile {
			f.Path = util.ExpandPatho(o.Some(root), f.Path)
			return f
		})

		schema.Templates.Files = slicer.Map(schema.Templates.Files, func(f cp.CodeGenProjectTemplateFile) cp.CodeGenProjectTemplateFile {
			f.Path = util.ExpandPatho(o.Some(root), f.Path)
			return f
		})

		t.projectManager.Register(root, source, &schema)
	case util.IsCodeGenSchemaTypeFile(source):
		if schema, err := codegen_type.UnmarshalJson(bs); err != nil {
			return err
		} else {
			t.typeManager.Register(root, source, schema)
		}
	}

	return nil
}
