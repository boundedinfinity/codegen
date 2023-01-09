package util

import (
	ct "boundedinfinity/codegen/codegen_type"
	"boundedinfinity/codegen/codegen_type/codegen_type_id"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/boundedinfinity/go-commoner/caser"
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
	codegenProjectExts = []string{
		".codegen.json",
		".codegen.yaml",
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

func RemoveSchema(s string) string {
	vs := append([]string{}, jsonSchemaExts...)
	vs = append(vs, codegenProjectExts...)
	vs = append(vs, codegenTypeExts...)

	for _, v := range vs {
		v = extentioner.Strip(v)
		s = strings.ReplaceAll(s, v, "")
	}

	return s
}

func EnsureAbs(root, path any) string {
	eRoot := environmenter.Sub(fmt.Sprint(root))
	new := fmt.Sprint(path)
	new = environmenter.Sub(new)

	if !filepath.IsAbs(new) {
		new = filepath.Join(eRoot, new)
	}

	new = filepath.Clean(new)

	return new
}

func ExpandPatho(root, path o.Option[string]) o.Option[string] {
	new := EnsureAbs(root.Get(), path.Get())

	return o.Some(new)
}

func IsCodeGenSchemaTypeFile(v string) bool {
	return slicer.ContainsFn(codegenTypeExts, func(x string) bool {
		return strings.HasSuffix(v, x)
	})
}

func IsCodeGenSchemaFile(v string) bool {
	return slicer.ContainsFn(codegenProjectExts, func(x string) bool {
		return strings.HasSuffix(v, x)
	})
}

func IsCodeGenFile(v string) bool {
	var s []string
	s = append(s, codegenProjectExts...)
	s = append(s, codegenTypeExts...)

	return slicer.ContainsFn(s, func(x string) bool {
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

func GetSchemaTypeId(path o.Option[string]) o.Option[codegen_type_id.CodgenTypeId] {
	if path.Empty() {
		return o.None[codegen_type_id.CodgenTypeId]()
	}

	filename := pather.Base(path.Get())

	found, ok := slicer.FindFn(codegen_type_id.All, func(v codegen_type_id.CodgenTypeId) bool {
		return strings.HasPrefix(filename, string(v))
	})

	if ok {
		return o.Some(found)
	}

	return o.None[codegen_type_id.CodgenTypeId]()
}

func GetTemplateMimeType(path string) trier.Try[mime_type.MimeType] {
	ext := path
	ext = extentioner.Ext(ext)
	tm, err := file_extention.FromExt(ext)
	return trier.Complete(tm, err)
}

func GetTemplateExt(path string) trier.Try[string] {
	mimeType := GetTemplateMimeType(path)
	return getExt(path, mimeType)
}

func GetOutputMimeType(path string) trier.Try[mime_type.MimeType] {
	ext := path
	ext = extentioner.Strip(ext)
	ext = extentioner.Ext(ext)
	tm, err := file_extention.FromExt(ext)
	return trier.Complete(tm, err)
}

func getExt(path string, mimeType trier.Try[mime_type.MimeType]) trier.Try[string] {
	if mimeType.Failure() {
		return trier.Failure[string](mimeType.Error)
	}

	exts, err := file_extention.GetExts(mimeType.Result)

	if err != nil {
		return trier.Failure[string](err)
	}

	name := pather.Base(path)
	var found string

	for _, ext := range exts {
		if strings.Contains(name, ext.String()) {
			found = ext.String()
			break
		}
	}

	if found == "" {
		return trier.Failure[string](fmt.Errorf("can't find extention"))
	}

	return trier.Success(found)
}

func GetOutputExt(path string) trier.Try[string] {
	mimeType := GetOutputMimeType(path)
	return getExt(path, mimeType)
}

func GetOutputPath(outputDir string, file ct.CodeGenProjectTemplateFile, typ ct.CodeGenType) trier.Try[string] {
	if file.SourcePath.Empty() {
		return trier.Failuref[string]("sourcePath missing")
	}

	var name string
	var path string

	name = typ.Namespace().SchemaNs
	name = pather.Base(name)
	name = caser.KebabToPascal(name)
	name = name + extentioner.Ext(extentioner.Strip(file.SourcePath.Get()))

	path = file.SourcePath.Get()
	path = strings.ReplaceAll(path, file.RootPath.Get(), "")
	path = pather.Dir(path)
	path = pather.Join(outputDir, path, name)

	return trier.Success(path)
}
