package system

import (
	"boundedinfinity/codegen/model"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/boundedinfinity/jsonschema"
	"github.com/boundedinfinity/mimetyper/mime_type"
	"gopkg.in/yaml.v2"
)

func (t *System) unmarshalJsonSchema(info *model.SourceInfo, bs []byte) error {
	var schemas []jsonschema.JsonSchmea

	if err := t.jsonSchema.Unmarshal(&schemas, info.MimeType, bs); err != nil {
		return err
	}

	for _, schema := range schemas {
		if err := t.jsonSchema.Add(&schema); err != nil {
			return err
		}
	}

	return nil
}

func (t *System) unmarshalCodeGen(info *model.SourceInfo, bs []byte) error {
	var schemas []model.Schema

	switch info.MimeType {
	case mime_type.ApplicationXYaml:
		if err := t.unmarshalYaml(&schemas, bs); err != nil {
			return err
		}
	case mime_type.ApplicationJson:
		if err := t.unmarshalJson(&schemas, bs); err != nil {
			return err
		}
	default:
		mime_type.Error(model.SUPPORTED_MIMETYPES, string(info.MimeType))
	}

	for _, schema := range schemas {
		if schema.Id == "" {
			return model.ErrCodeGenIdEmpty
		}

		if _, ok := t.codeGen[schema.Id]; ok {
			return model.ErrCodeGenIdDuplicateV(schema.Id)
		}

		t.codeGen[schema.Id] = &schema
	}

	return nil
}

func (t *System) unmarshalYaml(ss *[]model.Schema, bs []byte) error {
	d := yaml.NewDecoder(bytes.NewReader(bs))

	for {
		s := new(model.Schema)

		err := d.Decode(&s)

		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			return err
		}

		if s == nil {
			continue
		}

		*ss = append(*ss, *s)
	}

	return nil
}

func (t *System) unmarshalJson(ss *[]model.Schema, bs []byte) error {
	s := string(bs)
	s = strings.TrimSpace(s)
	f := s[0:1]

	switch f {
	case "{":
		var x model.Schema
		if err := json.Unmarshal(bs, &s); err != nil {
			return err
		}
		*ss = append(*ss, x)
	case "[":
		var xs []model.Schema

		if err := json.Unmarshal(bs, &xs); err != nil {
			return err
		}

		*ss = append(*ss, xs...)
	default:
		return fmt.Errorf("unsupported file type")
	}

	return nil
}
