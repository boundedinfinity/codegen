package loader

import (
	"boundedinfinity/codegen/canonical"
	"boundedinfinity/codegen/model"
	"boundedinfinity/codegen/util"
	"encoding/json"

	"github.com/boundedinfinity/go-marshaler"
	"github.com/boundedinfinity/go-mimetyper/mime_type"
	"github.com/ghodss/yaml"
)

func (t *Loader) LoadUri(uris ...string) error {
	if err := t.cacher.Cache("schema", uris...); err != nil {
		return err
	}

	cdatas := t.cacher.FindByGroup("schema").Get()

	for _, cdata := range cdatas {
		switch {
		case util.IsJsonSchemaFile(cdata.DestPath):
			if err := t.jsonSchemas.LoadPath(cdata.DestPath); err != nil {
				return err
			}
		case util.IsCodeGenSchemaFile(cdata.DestPath), util.IsCodeGenSchemaTypeFile(cdata.DestPath):
			if err := t.LoadCodeGenPath(cdata.DestPath); err != nil {
				return err
			}
		default:
			return model.ErrUnsupportedSchemev(cdata.DestPath)
		}
	}

	return nil
}

func (t *Loader) LoadCodeGenTypePath(root string) error {

	return nil
}

func (t *Loader) LoadCodeGenPath(root string) error {
	m, err := marshaler.ReadFromPath(root)

	if err != nil {
		return err
	}

	for path, content := range m {
		if t.cgsPathMap.Has(path) {
			return model.ErrPathDuplicatev(path)
		}

		if err := t.LoadSchema(content.Data, content.MimeType, path); err != nil {
			return err
		}
	}

	return nil
}

func (t *Loader) LoadSchema(data []byte, mt mime_type.MimeType, path string) error {
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
		return model.ErrMimeTypeUnsupportedv(mt)
	}

	switch {
	case util.IsCodeGenSchemaFile(path):
		var schema model.CodeGenSchema

		if err := json.Unmarshal(bs, &schema); err != nil {
			return err
		}

		t.cgsPathMap[path] = schema
	case util.IsCodeGenSchemaTypeFile(path):
		if schema, err := canonical.UnmarshalCanonicalSchemaJson(bs); err != nil {
			return err
		} else {
			t.canonicalPathMap[path] = schema
		}
	}

	return nil
}
