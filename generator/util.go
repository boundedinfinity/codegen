package generator

import (
	cp "boundedinfinity/codegen/codegen_project"
	lc "boundedinfinity/codegen/loader_context"
	rc "boundedinfinity/codegen/render_context"
	"path"
	"path/filepath"
	"strings"

	"github.com/boundedinfinity/go-commoner/caser"
	"github.com/boundedinfinity/go-commoner/extentioner"
)

func (t *Generator) SchemaNamepace(rootNs string, lc lc.TypeLoaderContext) string {
	if lc.Schema.Base().Base().Id.Empty() {
		return ""
	}

	var ns string

	ns = lc.FileInfo.Source
	ns = strings.Replace(ns, lc.FileInfo.Root, "", 1)
	ns = path.Join(rootNs, ns)
	ns = extentioner.Strip(ns)
	ns = extentioner.Strip(ns)

	return ns
}

func (t *Generator) RelNamepace(rootNs string, lc lc.TypeLoaderContext) string {
	schemaNs := t.SchemaNamepace(rootNs, lc)

	if schemaNs == "" {
		return ""
	}

	relNs := schemaNs
	relNs = strings.ReplaceAll(schemaNs, rootNs, "")
	relNs = strings.Replace(relNs, "/", "", 1)
	return relNs
}

func (t *Generator) DestPath(info cp.CodeGenProjectInfo, schema rc.RenderContext, tmplPath string) string {
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

func CurrentNs(info cp.CodeGenProjectInfo, outputPath string) string {
	out := outputPath
	out = path.Dir(out)
	out = strings.ReplaceAll(out, info.DestDir.Get(), info.Namespace.Get())
	return out
}
