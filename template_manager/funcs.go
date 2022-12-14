package template_manager

import (
	"boundedinfinity/codegen/render_context"
	"boundedinfinity/codegen/template_manager/dumper"
	"fmt"
	"path/filepath"
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
		TemplateFunc("PATH_REL", t.pathRel),
		TemplateFunc("DEFINED", t.defined),
		TemplateFunc("EMPTY", t.empty),
		TemplateFunc("SINGULAR", t.singular),
		TemplateFunc("PLURAL", t.plural),
		TemplateFunc("RES_SCHEMA_NS", t.resolveSchemaNs),
		TemplateFunc("RES_SCHEMA", t.resolveSchema),
	)

	for _, arg := range args {
		arg(t)
	}

	return nil
}

func dumpJson(obj any) string {
	return dumper.New().Dump(obj)
}

func (t *TemplateManager) resolveSchemaNs(schema render_context.RenderContext) string {
	var found string

	render_context.WalkBase(schema, func(base *render_context.RenderContextBase) error {
		if found != "" {
			return render_context.ErrExit
		}

		if base.SchemaNs != "" {
			found = base.SchemaNs
			return render_context.ErrExit
		}

		return nil
	})

	return found
}

func (t *TemplateManager) resolveSchema(schema render_context.RenderContext) render_context.RenderContext {
	switch c := schema.(type) {
	case *render_context.RenderContextArray:
		return t.resolveSchema(c.Items)
	case *render_context.RenderContextRef:
		return t.resolveSchema(c.Ref)
	default:
		return schema
	}
}

func (t *TemplateManager) singular(s string) string {
	return pluralize.NewClient().Singular(s)
}

func (t *TemplateManager) plural(s string) string {
	return pluralize.NewClient().Plural(s)
}

func (t *TemplateManager) pathRel(a, b string) (string, error) {
	r, err := filepath.Rel(a, b)

	if err != nil {
		return r, err
	}

	if !strings.HasPrefix(r, "..") && !strings.HasPrefix(r, "/") {
		r = "./" + r
	}

	return r, err
}

func (t *TemplateManager) pathBase(s string) string {
	p := s
	p = filepath.Base(p)

	// if p == "." {
	// 	p = ""
	// }

	return p
}

func (t *TemplateManager) pathDir(s string) string {
	p := s
	p = filepath.Dir(p)

	// if p == "." {
	// 	p = ""
	// }

	return p
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
