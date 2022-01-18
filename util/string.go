package util

import (
	"path"
	"strings"
	"unicode"

	"github.com/boundedinfinity/caser"
)

func ReplaceMap(v string, m map[string]string) string {
	for o, n := range m {
		v = strings.ReplaceAll(v, o, n)
	}

	return v
}

func Normalize(v string) string {
	var o strings.Builder

	for i, r := range v {
		if i == 0 {
			if unicode.IsLetter(r) || r == '_' {
				o.WriteRune(r)
			} else {
				o.WriteRune('_')
			}
		} else {
			if unicode.IsLetter(r) || unicode.IsNumber(r) || r == '_' {
				o.WriteRune(r)
			} else {
				o.WriteRune('_')
			}
		}
	}

	return o.String()
}

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
	return caser.PhraseToCamel(v)
}

func PascalCase(v string) string {
	return caser.PhraseToCamel(v)
}

func PathBase(v string) string {
	return path.Base(v)
}

func PathDir(v string) string {
	return path.Dir(v)
}

func PathDirBase(v string) string {
	o := v

	if strings.Contains(v, "/") {
		o = path.Base(path.Dir(v))
	} else {
		o = v
	}

	return o
}

// func SameNamespace(a, b model.OutputModel) bool {
// 	aNs := path.Dir(a.FullName)
// 	bNs := path.Dir(b.FullName)
// 	return aNs == bNs
// }

// func ToJson(v interface{}) string {
// 	j, _ := json.MarshalIndent(v, "", "    ")
// 	return string(j)
// }
