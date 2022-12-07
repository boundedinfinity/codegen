package template_manager

import (
	"boundedinfinity/codegen/canonical"
	"boundedinfinity/codegen/template_manager/dumper"
	"path"
	"path/filepath"

	"github.com/boundedinfinity/go-commoner/caser"
)

func dumpJson(obj any) string {
	return dumper.New().Dump(obj)
}

func (t *TemplateManager) typePath(rc RenderContext) string {
	v := rc.Schema.(canonical.CanonicalBase).Id.Get()
	v = filepath.Join(filepath.Dir(v), caser.KebabToPascal(filepath.Base(v)))
	return v
}

func (t *TemplateManager) importPath(rc RenderContext) string {
	v := filepath.Dir(t.typePath(rc))
	return v
}

func (t *TemplateManager) importType(rc RenderContext) string {
	v := t.typePath(rc)
	v2 := filepath.Base(filepath.Dir(v))
	v = v2 + "." + v
	return v
}

func (t *TemplateManager) objPath(rc RenderContext) string {
	var out string

	// id := t.jsonSchemas.Id(rc.Schema)

	// if id.Defined() {
	// 	out = string(id.Get())
	// 	_, p, _ := urischemer.Break(out)
	// 	out = p
	// } else {
	// 	// TODO
	// }

	// if t.codeGenSchema.Info.Package.Defined() {
	// 	out = path.Join(t.codeGenSchema.Info.Package.Get(), out)
	// } else {
	// 	// TODO
	// }

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
