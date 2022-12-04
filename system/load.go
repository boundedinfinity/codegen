package system

import (
	"boundedinfinity/codegen/model"
	"boundedinfinity/codegen/util"
	"encoding/json"

	"github.com/boundedinfinity/go-marshaler"
	"github.com/boundedinfinity/go-mimetyper/mime_type"
	"github.com/ghodss/yaml"
)

func (t *System) LoadUri(uris ...string) error {
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
		case util.IsCodeGenSchemaFile(cdata.DestPath):
			if err := t.LoadPath(cdata.DestPath); err != nil {
				return err
			}
		default:
			return model.ErrUnsupportedSchemev(cdata.DestPath)
		}
	}

	return nil
}

func (t *System) LoadPath(root string) error {
	m, err := marshaler.ReadFromPath(root)

	if err != nil {
		return err
	}

	for path, content := range m {
		if t.pathMap.Has(path) {
			return model.ErrPathDuplicatev(path)
		}

		if err := t.LoadSchema(content.Data, content.MimeType, path); err != nil {
			return err
		}
	}

	return nil
}

func (t *System) LoadSchema(data []byte, mt mime_type.MimeType, path string) error {
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

	var schema model.CodeGenSchema

	if err := json.Unmarshal(bs, &schema); err != nil {
		return err
	}

	t.pathMap[path] = schema

	return nil
}
