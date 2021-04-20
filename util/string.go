package util

import (
	"path"
	"strings"

	"github.com/ozgio/strutil"
)

func Uc(v string) string {
	return strings.ToUpper(v)
}

func UcFirst(v string) string {
	f := string(v[0])
	r := string(v[1:])
	return strings.ToUpper(f) + r
}

func Lc(v string) string {
	return strings.ToLower(v)
}

func LcFirst(v string) string {
	f := string(v[0])
	r := string(v[1:])
	return strings.ToLower(f) + r
}

func CamelCase(v string) string {
	return LcFirst(strutil.ToCamelCase(v))
}

func BasePath(v string) string {
	return path.Base(v)
}
