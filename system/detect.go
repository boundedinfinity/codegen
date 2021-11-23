package system

import (
	"boundedinfinity/codegen/template_type"
	"boundedinfinity/codegen/uritype"
	"net/url"
	"path/filepath"
	"strings"

	"github.com/boundedinfinity/mimetyper/file_extention"
	"github.com/boundedinfinity/mimetyper/mime_type"
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

func (t *System) detectMimeType(uri string, typ *mime_type.MimeType) error {
	mt, err := file_extention.DetectMimeType(uri)

	if err != nil {
		return nil
	}

	*typ = mt

	return nil
}
