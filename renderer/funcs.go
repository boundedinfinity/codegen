package renderer

import (
	ct "boundedinfinity/codegen/codegen_type"
	"boundedinfinity/codegen/renderer/dumper"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/boundedinfinity/go-commoner/caser"
	"github.com/boundedinfinity/go-commoner/optioner"
	"github.com/gertd/go-pluralize"
)

func dumpJson(obj any) string {
	return dumper.New().Dump(obj)
}

func (t *Renderer) resolveSchemaNs(schema ct.RenderNamespace) string {
	return schema.Namespace().SchemaNs
}

func (t *Renderer) resolveSchema(schema ct.CodeGenType) ct.CodeGenType {
	switch c := schema.(type) {
	case *ct.CodeGenTypeArray:
		return c.Items
	case *ct.CodeGenTypeRef:
		return c.Resolved
	default:
		return schema
	}
}

func (t *Renderer) singular(s string) string {
	return pluralize.NewClient().Singular(s)
}

func (t *Renderer) plural(s string) string {
	return pluralize.NewClient().Plural(s)
}

func (t *Renderer) pathRel(a, b string) (string, error) {
	r, err := filepath.Rel(a, b)

	if err != nil {
		return r, err
	}

	if !strings.HasPrefix(r, "..") && !strings.HasPrefix(r, "/") {
		r = "./" + r
	}

	return r, err
}

func (t *Renderer) pathBase(s string) string {
	p := s
	p = filepath.Base(p)

	// if p == "." {
	// 	p = ""
	// }

	return p
}

func (t *Renderer) pathDir(s string) string {
	p := s
	p = filepath.Dir(p)

	// if p == "." {
	// 	p = ""
	// }

	return p
}

func (t *Renderer) camel(s any) string {
	return caser.KebabToCamel(a2s(s))
}

func (t *Renderer) pascal(s any) string {
	return caser.KebabToPascal(a2s(s))
}

func (t *Renderer) snake(s any) string {
	return caser.KebabToSnake(a2s(s))
}

func (t *Renderer) defined(s any) bool {
	return a2s(s) != ""
}

func (t *Renderer) empty(s any) bool {
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
