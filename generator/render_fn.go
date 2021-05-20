package generator

import (
	"boundedinfinity/codegen/model"
	"boundedinfinity/codegen/util"
	"fmt"
	"path"
	"path/filepath"
	"strings"
)

func (t *Generator) langType(m *model.OutputModel) string {
	typ := "<unkown type>"

	switch m.Type {
	case model.SchemaType_String, model.SchemaType_Int, model.SchemaType_Long, model.SchemaType_Float, model.SchemaType_Double, model.SchemaType_Boolean:
		typ = t.schema2Primtive(m.Type.String())
	case model.SchemaType_Ref:
		if m.Imported {
			n := util.UcFirst(util.NsBase(m.Ref))
			r := util.NsBase(util.NsDir(m.Ref))
			typ = fmt.Sprintf("%v.%v", r, n)
		} else {
			typ = util.UcFirst(path.Base(m.Ref))
		}
	}

	return typ
}

func (t *Generator) schema2Primtive(s string) string {
	p, _ := t.spec.Info.Primitives[s]
	return p
}

func ifeq(v1, v2 string) bool {
	return v1 == v2
}

func path2CamelCase(v string) string {
	ss := strings.Split(v, "/")
	s := strings.Join(ss, " ")
	s = strings.Title(s)
	s = strings.ReplaceAll(s, " ", "")
	return s
}

func filename(v string) string {
	var fn string
	fn = v
	fn = filepath.Base(fn)
	fn = strings.TrimSuffix(fn, filepath.Ext(fn))
	return fn
}
