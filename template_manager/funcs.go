package template_manager

import (
	"boundedinfinity/codegen/canonical"
	"boundedinfinity/codegen/template_manager/dumper"
	"boundedinfinity/codegen/util"
	"fmt"
	"path"

	"github.com/boundedinfinity/go-commoner/caser"
)

func dumpJson(obj any) string {
	return dumper.New().Dump(obj)
}

func (t *TemplateManager) namespace(schema canonical.Canonical) string {
	return util.SchemaNamepace(t.codeGenSchema.Info, schema)
}

func (t *TemplateManager) getPackage(schema canonical.Canonical) string {
	return util.SchemaPackage(t.codeGenSchema.Info, schema)
}

func (t *TemplateManager) baseType(schema canonical.Canonical) string {
	return util.SchemaBaseType(t.codeGenSchema.Info, schema)
}

func (t *TemplateManager) camel(s fmt.Stringer) string {
	return caser.KebabToCamel(s.String())
}

func (t *TemplateManager) pascal(s fmt.Stringer) string {
	return caser.KebabToPascal(s.String())
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
