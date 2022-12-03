package template_manager

import (
	"boundedinfinity/codegen/template_manager/dumper"
	"path"

	"github.com/boundedinfinity/go-commoner/caser"
	"github.com/boundedinfinity/go-urischemer"
)

func dumpJson(obj any) string {
	return dumper.New().Dump(obj)
}

func (t *TemplateManager) objPath(rc RenderContext) string {
	var out string

	id := t.jsonSchemas.Id(rc.Schema)

	if id.Defined() {
		out = string(id.Get())
		_, p, _ := urischemer.Break(out)
		out = p
	} else {
		// TODO
	}

	if t.codeGenSchema.Info.Package.Defined() {
		out = path.Join(t.codeGenSchema.Info.Package.Get(), out)
	} else {
		// TODO
	}

	return out
}

func (t *TemplateManager) objName(rc RenderContext) string {
	out := t.objPath(rc)
	out = path.Base(out)
	out = caser.KebabToPascal(out)
	return out
}

func (t *TemplateManager) objPackage(rc RenderContext) string {
	out := t.objPath(rc)
	out = path.Dir(out)
	return out
}

func (t *TemplateManager) objPackageBase(rc RenderContext) string {
	out := t.objPath(rc)
	out = path.Dir(out)
	out = path.Base(out)
	return out
}
