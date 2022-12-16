package codegen_project

import (
	"boundedinfinity/codegen/codegen_type"
	"boundedinfinity/codegen/render_context"
	"path"
	"path/filepath"
	"strings"

	"github.com/boundedinfinity/go-commoner/extentioner"
)

func SchemaNamepace(info CodeGenProjectInfo, schema codegen_type.CodeGenType) string {
	if schema.Base().Id.Empty() {
		return ""
	}

	ns := schema.Base().Source
	ns = strings.Replace(ns, schema.Base().Root, "", 1)
	ns = path.Join(info.Namespace.Get(), ns)
	ns = extentioner.Strip(ns)
	ns = extentioner.Strip(ns)

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

func DestPath(info CodeGenProjectInfo, schema render_context.RenderContext, tmplPath string) string {
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

func CurrentNs(info CodeGenProjectInfo, outputPath string) string {
	out := outputPath
	out = path.Dir(out)
	out = strings.ReplaceAll(out, info.DestDir.Get(), info.Namespace.Get())
	return out
}
