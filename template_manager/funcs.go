package template_manager

import (
	"boundedinfinity/codegen/canonical"
	"boundedinfinity/codegen/template_manager/dumper"
	"path"
	"path/filepath"

	"github.com/boundedinfinity/go-commoner/caser"
	"github.com/boundedinfinity/go-urischemer"
)

func dumpJson(obj any) string {
	return dumper.New().Dump(obj)
}

func (t *TemplateManager) namespace(schema canonical.Canonical) string {
	id := schema.SchemaId()

	if id.Empty() {
		return "NO-ID"
	}

	ns := id.Get()
	_, ns, _ = urischemer.Break(ns)
	ns = path.Join(t.codeGenSchema.Info.Namespace.Get(), ns)
	ns = path.Join(path.Dir(ns), caser.KebabToPascal(path.Base(ns)))

	return ns
}

func (t *TemplateManager) getPackage(schema canonical.Canonical) string {
	pkg := t.namespace(schema)
	pkg = path.Dir(pkg)
	pkg = path.Base(pkg)
	return pkg
}

func (t *TemplateManager) importType(schema canonical.Canonical) string {
	v := t.namespace(schema)
	v2 := filepath.Base(filepath.Dir(v))
	v = v2 + "." + v
	return v
}

func (t *TemplateManager) objPath(schema canonical.Canonical) string {
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

func (t *TemplateManager) objName(schema canonical.Canonical) string {
	out := t.objPath(schema)
	out = path.Base(out)
	out = caser.KebabToPascal(out)
	return out
}

func (t *TemplateManager) objPackage(schema canonical.Canonical) string {
	out := t.objPath(schema)
	out = path.Dir(out)
	return out
}

func (t *TemplateManager) objPackageBase(schema canonical.Canonical) string {
	out := t.objPath(schema)
	out = path.Dir(out)
	out = path.Base(out)
	return out
}
