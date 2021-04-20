package generator

import (
	"encoding/json"
	"fmt"
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

func to_json(v interface{}) string {
	var s string
	indent_spaces := strings.Repeat(" ", 4)
	if bs, err := json.MarshalIndent(v, "", indent_spaces); err != nil {
		s = string(bs)
	}

	return s
}
