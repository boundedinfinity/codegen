package generator

import (
	jmodel "github.com/boundedinfinity/go-jsonschema/model"
)

func (t *Generator) PackageBase(s jmodel.JsonSchema) string {
	// result := pl.New[string]().Append(path.Base).RunSingle(s.Id)
	// return result
	return ""
}

// func (t *Generator) schema2Primtive(s string) string {
// 	p, _ := t.spec.Info.Primitives[s]
// 	return p
// }

// func ifeq(v1, v2 string) bool {
// 	return v1 == v2
// }

// func path2CamelCase(v string) string {
// 	ss := strings.Split(v, "/")
// 	s := strings.Join(ss, " ")
// 	s = strings.Title(s)
// 	s = strings.ReplaceAll(s, " ", "")
// 	return s
// }

// func filename(v string) string {
// 	var fn string
// 	fn = v
// 	fn = filepath.Base(fn)
// 	fn = strings.TrimSuffix(fn, filepath.Ext(fn))
// 	return fn
// }
