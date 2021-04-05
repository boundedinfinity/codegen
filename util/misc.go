package util

import (
	"encoding/json"
	"fmt"
	"runtime"
	"strings"

	"github.com/blang/semver/v4"
)

func Jdump(v interface{}) string {
	bs, _ := json.MarshalIndent(v, "", "    ")

	return string(bs)
}

func GetVersion() (string, error) {
	var o string

	o = runtime.Version()
	o = strings.ReplaceAll(o, "go", "")
	sv, err := semver.Make(o)

	if err != nil {
		return o, err
	}

	o = fmt.Sprintf("%v.%v", sv.Major, sv.Minor)

	return o, nil
}

func StrSliceMap(i []string, fn func(string) string) []string {
	var o []string

	for _, x := range i {
		o = append(o, fn(x))
	}

	return o
}
