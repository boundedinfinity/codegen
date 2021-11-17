package system

import (
	"boundedinfinity/codegen/schema_ext"
	"boundedinfinity/codegen/template_type"
	"boundedinfinity/codegen/uritype"
	"net/url"
	"path/filepath"
	"strings"

	"github.com/boundedinfinity/jsonschema/mimetype"
)

func (t *System) detectTemplateType(uri string, typ *template_type.TemplateType) error {
	ext := filepath.Ext(uri)
	ext = strings.TrimPrefix(ext, ".")
	parsed, err := template_type.Parse(ext)

	if err != nil {
		return err
	}

	*typ = parsed
	return nil
}

func (t *System) detectUriType(uri string, typ *uritype.UriType) error {
	parsed, err := url.Parse(uri)

	if err != nil {
		return err
	}

	x, err := uritype.Parse(parsed.Scheme)

	if err != nil {
		return err
	}

	*typ = x

	return nil
}

func (t *System) detectMimeType(uri string, typ *mimetype.MimeType) error {
	ext := filepath.Ext(uri)
	test := strings.TrimPrefix(ext, ".")
	mt, err := schema_ext.Parse(test)

	if err != nil {
		return err
	}

	switch mt {
	case schema_ext.Json:
		*typ = mimetype.ApplicationJson
	case schema_ext.Yaml, schema_ext.Yml:
		*typ = mimetype.ApplicationXYaml
	}

	return nil
}
