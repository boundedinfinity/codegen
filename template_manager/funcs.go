package template_manager

import (
	"boundedinfinity/codegen/canonical"
	"boundedinfinity/codegen/template_manager/dumper"
	"boundedinfinity/codegen/util"
	"fmt"
	"path"

	"github.com/boundedinfinity/go-commoner/caser"
	"github.com/boundedinfinity/go-commoner/optioner"
)

func (t *TemplateManager) initTemplatesFuncs() error {
	args := make([]Arg, 0)

	args = append(args,
		TemplateFunc("DUMP", dumpJson),
		TemplateFunc("PASCAL", t.pascal),
		TemplateFunc("CAMEL", t.camel),
		TemplateFunc("SNAKE", t.camel),
		TemplateFunc("BASE", t.pathBase),
		TemplateFunc("DIR", t.pathDir),
		TemplateFunc("DEFINED", t.defined),
		TemplateFunc("EMPTY", t.empty),
	)

	for _, arg := range args {
		arg(t)
	}

	return nil
}

func dumpJson(obj any) string {
	return dumper.New().Dump(obj)
}

func (t *TemplateManager) pathBase(s string) string {
	return path.Base(s)
}

func (t *TemplateManager) pathDir(s string) string {
	return path.Dir(s)
}

func (t *TemplateManager) namespace(schema canonical.Canonical) string {
	return util.SchemaNamepace(t.codeGenSchema.Info, schema)
}

func (t *TemplateManager) camel(s any) string {
	return caser.KebabToCamel(a2s(s))
}

func (t *TemplateManager) pascal(s any) string {
	return caser.KebabToPascal(a2s(s))
}

func (t *TemplateManager) snake(s any) string {
	return caser.KebabToSnake(a2s(s))
}

func (t *TemplateManager) defined(s any) bool {
	return a2s(s) != ""
}

func (t *TemplateManager) empty(s any) bool {
	return a2s(s) == ""
}

func a2s(a any) string {
	var s string

	switch v := a.(type) {
	case string:
		s = v
	case fmt.Stringer:
		s = v.String()
	case optioner.Option[string]:
		s = v.OrElse("")
	default:
		s = fmt.Sprintf("%v", a)
	}

	return s
}
