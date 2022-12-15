package util

import (
	"boundedinfinity/codegen/canonical"
	"boundedinfinity/codegen/canonical/canonical_type"
	"boundedinfinity/codegen/model"
	"boundedinfinity/codegen/render_context"
	"path"
	"path/filepath"
	"strings"

	"github.com/boundedinfinity/go-commoner/environmenter"
	"github.com/boundedinfinity/go-commoner/extentioner"
	"github.com/boundedinfinity/go-commoner/optioner"
	"github.com/boundedinfinity/go-commoner/pather"
	"github.com/boundedinfinity/go-commoner/slicer"
	"github.com/boundedinfinity/go-commoner/trier"
	"github.com/boundedinfinity/go-jsonschema/schematype"
	"github.com/boundedinfinity/go-mimetyper/file_extention"
	"github.com/boundedinfinity/go-mimetyper/mime_type"
	"github.com/boundedinfinity/go-urischemer"
)

var (
	codegenExts = []string{
		".codegen.json",
		".codegen.yaml",
		".codegen.yml",
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

func ResolveUri(source string, v optioner.Option[string]) (optioner.Option[string], error) {
	if v.Empty() {
		return optioner.None[string](), nil
	}

	s, p, err := urischemer.Break(v.Get())

	if err != nil {
		return optioner.None[string](), nil
	}

	p2 := ResolvePath(source, optioner.Some(p))
	p3 := urischemer.Combine(s, p2.Get())

	return optioner.Some(p3), nil
}

func ResolvePath(source string, v optioner.Option[string]) optioner.Option[string] {
	if v.Empty() {
		return optioner.None[string]()
	}

	p := v.Get()
	p = environmenter.Sub(p)

	if !filepath.IsAbs(p) {
		p = filepath.Join(source, p)

	}

	p = filepath.Clean(p)

	return optioner.Some(p)
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

func IsJsonSchemaTemplate(typ optioner.Option[schematype.SchemaType], path string) bool {
	if typ.Empty() {
		return false
	}

	base := filepath.Base(path)
	ts := typ.String()
	return strings.HasPrefix(base, ts)
}

func GetCanonicalType(path string) optioner.Option[canonical_type.CanonicalType] {
	filename := pather.Base(path)
	found, ok := slicer.FindFn(canonical_type.All, func(v canonical_type.CanonicalType) bool {
		return strings.HasPrefix(filename, string(v))
	})

	if ok {
		return optioner.Some(found)
	}

	return optioner.None[canonical_type.CanonicalType]()
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

func SchemaNamepace(info model.CodeGenSchemaInfo, schema canonical.Canonical) string {
	id := schema.SchemaId()

	if id.Empty() {
		return ""
	}

	ns := id.Get()
	_, ns, _ = urischemer.Break(ns)
	ns = path.Join(info.Namespace.Get(), ns)
	ns = path.Join(path.Dir(ns), path.Base(ns))

	return ns
}

func RelNamepace(info model.CodeGenSchemaInfo, schema canonical.Canonical) string {
	schemaNs := SchemaNamepace(info, schema)

	if schemaNs == "" {
		return ""
	}

	relNs := schemaNs
	relNs = strings.ReplaceAll(schemaNs, info.Namespace.Get(), "")
	relNs = strings.Replace(relNs, "/", "", 1)
	return relNs
}

func DestPath(info model.CodeGenSchemaInfo, schema render_context.RenderContext, tmplPath string) string {
	file := tmplPath
	ns := schema.Base().SchemaNs
	file = filepath.Base(file)
	file = extentioner.Strip(file)
	file = filepath.Base(ns) + "." + file
	path := ns
	path = strings.ReplaceAll(ns, info.Namespace.Get(), "")
	path = filepath.Dir(path)
	path = filepath.Join(info.DestDir.Get(), path, file)
	return path
}

func CurrentNs(info model.CodeGenSchemaInfo, outputPath string) string {
	out := outputPath
	out = path.Dir(out)
	out = strings.ReplaceAll(out, info.DestDir.Get(), info.Namespace.Get())
	return out
}
