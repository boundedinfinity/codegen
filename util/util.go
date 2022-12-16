package util

import (
	"boundedinfinity/codegen/codegen_type/codegen_type_id"
	"path/filepath"
	"strings"

	"github.com/boundedinfinity/go-commoner/environmenter"
	"github.com/boundedinfinity/go-commoner/extentioner"
	o "github.com/boundedinfinity/go-commoner/optioner"
	"github.com/boundedinfinity/go-commoner/pather"
	"github.com/boundedinfinity/go-commoner/slicer"
	"github.com/boundedinfinity/go-commoner/trier"
	"github.com/boundedinfinity/go-jsonschema/schematype"
	"github.com/boundedinfinity/go-mimetyper/file_extention"
	"github.com/boundedinfinity/go-mimetyper/mime_type"
)

var (
	codegenExts = []string{
		".codegen-project.json",
		".codegen-project.yaml",
		".codegen-project.yml",
	}

	codegenTypeExts = []string{
		".codegen-type.json",
		".codegen-type.yaml",
		".codegen-type.yml",
	}

	jsonSchemaExts = []string{
		".json-schema.json",
		".json-schema.yaml",
		".json-schema.yml",
	}
)

func ExpandPath(root, path string) string {
	new := path
	new = environmenter.Sub(new)

	if !filepath.IsAbs(new) {
		new = filepath.Join(root, new)
	}

	new = filepath.Clean(new)

	return new
}

func ExpandPatho(root, path o.Option[string]) o.Option[string] {
	new := ExpandPath(root.Get(), path.Get())

	return o.Some(new)
}

func IsCodeGenSchemaTypeFile(v string) bool {
	return slicer.ContainsFn(codegenTypeExts, func(x string) bool {
		return strings.HasSuffix(v, x)
	})
}

func IsCodeGenSchemaFile(v string) bool {
	return slicer.ContainsFn(codegenExts, func(x string) bool {
		return strings.HasSuffix(v, x)
	})
}

func IsJsonSchemaFile(v string) bool {
	return slicer.ContainsFn(jsonSchemaExts, func(x string) bool {
		return strings.HasSuffix(v, x)
	})
}

func IsSchemaFile(v string) bool {
	return IsCodeGenSchemaFile(v) || IsJsonSchemaFile(v) || IsCodeGenSchemaTypeFile(v)
}

func IsJsonSchemaTemplate(typ o.Option[schematype.SchemaType], path string) bool {
	if typ.Empty() {
		return false
	}

	base := filepath.Base(path)
	ts := typ.String()
	return strings.HasPrefix(base, ts)
}

func GetCanonicalType(path string) o.Option[codegen_type_id.CodgenTypeId] {
	filename := pather.Base(path)
	found, ok := slicer.FindFn(codegen_type_id.All, func(v codegen_type_id.CodgenTypeId) bool {
		return strings.HasPrefix(filename, string(v))
	})

	if ok {
		return o.Some(found)
	}

	return o.None[codegen_type_id.CodgenTypeId]()
}

func GetTemplateType(path string) trier.Try[mime_type.MimeType] {
	var ext string
	ext = extentioner.Ext(path)
	tm, err := file_extention.GetMimeType(ext)
	return trier.Complete(tm, err)
}

func GetOutputType(path string) trier.Try[mime_type.MimeType] {
	ext := path
	ext = extentioner.Strip(ext)
	ext = extentioner.Ext(ext)
	tm, err := file_extention.GetMimeType(ext)
	return trier.Complete(tm, err)
}
