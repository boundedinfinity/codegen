package generator

import (
	"path/filepath"
	"strings"
)

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
