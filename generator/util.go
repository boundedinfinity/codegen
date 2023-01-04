package generator

import (
	ct "boundedinfinity/codegen/codegen_type"
	rc "boundedinfinity/codegen/render_context"
	"path"
	"path/filepath"
	"strings"

	"github.com/boundedinfinity/go-commoner/caser"
	"github.com/boundedinfinity/go-commoner/extentioner"
)

func (t *Generator) SchemaNamepace(rootNs string, lc ct.CodeGenType) string {
	if lc.Base().Base().Id.Empty() {
		return ""
	}

	var ns string

	ns = lc.Source().SourcePath.Get()
	ns = strings.Replace(ns, lc.Source().RootPath.Get(), "", 1)
	ns = path.Join(rootNs, ns)
	ns = extentioner.Strip(ns)
	ns = extentioner.Strip(ns)

	return ns
}

func (t *Generator) RelNamepace(rootNs string, lc ct.CodeGenType) string {
	schemaNs := t.SchemaNamepace(rootNs, lc)

	if schemaNs == "" {
		return ""
	}

	relNs := schemaNs
	relNs = strings.ReplaceAll(schemaNs, rootNs, "")
	relNs = strings.Replace(relNs, "/", "", 1)
	return relNs
}

func (t *Generator) DestPath(info ct.CodeGenInfo, schema rc.RenderContext, tmplPath string) string {
	sourceDir, sourceFile := filepath.Split(schema.Base().SourcePath.Get())
	rootPath := schema.Base().RootPath
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
	outPath = strings.Replace(outPath, rootPath.Get(), "", 1)
	outPath = filepath.Join(destPath, outPath, outFile)

	return outPath
}

func CurrentNs(info ct.CodeGenInfo, outputPath string) string {
	out := outputPath
	out = path.Dir(out)
	out = strings.ReplaceAll(out, info.DestDir.Get(), info.Namespace.Get())
	return out
}
