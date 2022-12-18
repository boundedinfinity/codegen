package codegen_project

import (
	"boundedinfinity/codegen/codegen_type"
	rc "boundedinfinity/codegen/render_context"
	"path"
	"path/filepath"
	"strings"

	"github.com/boundedinfinity/go-commoner/caser"
	"github.com/boundedinfinity/go-commoner/extentioner"
)

func SchemaNamepace(info CodeGenProjectInfo, schema codegen_type.CodeGenType) string {
	if schema.Base().Id.Empty() {
		return ""
	}

	var ns string
	// ns = schema.Base().Source
	// ns = strings.Replace(ns, schema.Base().Root, "", 1)
	// ns = path.Join(info.Namespace.Get(), ns)
	// ns = extentioner.Strip(ns)
	// ns = extentioner.Strip(ns)

	return ns
}

func RelNamepace(info CodeGenProjectInfo, schema codegen_type.CodeGenType) string {
	schemaNs := SchemaNamepace(info, schema)

	if schemaNs == "" {
		return ""
	}

	relNs := schemaNs
	relNs = strings.ReplaceAll(schemaNs, info.Namespace.Get(), "")
	relNs = strings.Replace(relNs, "/", "", 1)
	return relNs
}

func DestPath(info CodeGenProjectInfo, schema rc.RenderContext, tmplPath string) string {
	sourceDir, sourceFile := filepath.Split(schema.Base().Source)
	rootPath := schema.Base().Root
	destPath := info.DestDir.Get()

	tmplFile := tmplPath
	tmplFile = filepath.Base(tmplFile)
	tmplFile = extentioner.Strip(tmplFile)

	outFile := sourceFile
	outFile = extentioner.Strip(outFile)
	outFile = extentioner.Strip(outFile)
	outFile = caser.KebabToPascal(outFile)
	outFile = extentioner.Join(outFile, tmplFile)

	outPath := sourceDir
	outPath = strings.Replace(outPath, rootPath, "", 1)
	outPath = filepath.Join(destPath, outPath, outFile)

	return outPath
}

func CurrentNs(info CodeGenProjectInfo, outputPath string) string {
	out := outputPath
	out = path.Dir(out)
	out = strings.ReplaceAll(out, info.DestDir.Get(), info.Namespace.Get())
	return out
}
