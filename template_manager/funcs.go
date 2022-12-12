package template_manager

import (
	"boundedinfinity/codegen/template_manager/dumper"
	"fmt"
	"path"
	"strings"

	"github.com/boundedinfinity/go-commoner/caser"
	"github.com/boundedinfinity/go-commoner/optioner"
	"github.com/gertd/go-pluralize"
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
		TemplateFunc("SINGULAR", t.singular),
		TemplateFunc("PLURAL", t.plural),
	)

	for _, arg := range args {
		arg(t)
	}

	return nil
}

func dumpJson(obj any) string {
	return dumper.New().Dump(obj)
}

func (t *TemplateManager) singular(s string) string {
	return pluralize.NewClient().Singular(s)
}

func (t *TemplateManager) plural(s string) string {
	return pluralize.NewClient().Plural(s)
}

func (t *TemplateManager) pathBase(s string) string {
	return strings.ReplaceAll(path.Base(s), ".", "")
}

func (t *TemplateManager) pathDir(s string) string {
	return strings.ReplaceAll(path.Dir(s), ".", "")
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
