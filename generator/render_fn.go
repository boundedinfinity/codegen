package generator

import (
	"fmt"
	"path"
	"path/filepath"
	"strings"

	"github.com/ozgio/strutil"
)

func ifeq(v1, v2 string) bool {
	return v1 == v2
}

func uc(v string) string {
	return strings.ToUpper(v)
}

func ucFirst(v string) string {
	f := string(v[0])
	r := string(v[1:])
	return strings.ToUpper(f) + r
}

func path2CamelCase(v string) string {
	ss := strings.Split(v, "/")
	s := strings.Join(ss, " ")
	s = strings.Title(s)
	s = strings.ReplaceAll(s, " ", "")
	return s
}

func camelCase(v string) string {
	return lcFirst(strutil.ToCamelCase(v))
}

func lc(v string) string {
	return strings.ToLower(v)
}

func lcFirst(v string) string {
	f := string(v[0])
	r := string(v[1:])
	return strings.ToLower(f) + r
}

func basePath(v string) string {
	return path.Base(v)
}

func filename(v string) string {
	var fn string
	fn = v
	fn = filepath.Base(fn)
	fn = strings.TrimSuffix(fn, filepath.Ext(fn))
	return fn
}

// func goModelComment(m model.BiOutput_Model) string {

// 	return strings.Join(ls2, "\n")
// }

func linePrefix(v, p string) string {
	var ls2 []string
	var l2 string

	l2 = strings.TrimSuffix(v, "\n")
	ls1 := strings.Split(l2, "\n")

	for _, l := range ls1 {
		ls2 = append(ls2, fmt.Sprintf("%v %v", p, l))
	}
	return strings.Join(ls2, "\n")
}
