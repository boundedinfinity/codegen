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
	if results, err := t.cacher.Cache(uris...); err != nil {
		return err
	} else {
		for _, result := range results {
			t.cacheResults[result.DestPath] = result
		}
	}

	for _, path := range t.cacheResults.Keys().Get() {
		switch {
		case util.IsJsonSchemaFile(path):
			if err := t.jsonSchemas.LoadPath(path); err != nil {
				return err
			}
		case util.IsCodeGenSchemaFile(path):
			if err := t.LoadPath(path); err != nil {
				return err
			}
		default:
			return model.ErrUnsupportedSchemev(path)
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
