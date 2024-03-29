package dumper

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

type dumper struct {
	indent  string
	comment string
}

type errJson struct {
	Err string `json:"error" yaml:"error"`
}

func (t *dumper) Dump(v any) string {
	name := reflect.TypeOf(v).Name()
	var out string

	if bs, err := json.MarshalIndent(v, "// ", t.indent); err != nil {
		errBs, _ := json.MarshalIndent(errJson{Err: err.Error()}, "", t.indent)
		out = string(errBs)
	} else {
		out = string(bs)
	}

	out = fmt.Sprintf("// %v\n// ", name) + out

	return out
}

func New(args ...Arg) *dumper {
	t := &dumper{}

	for _, arg := range args {
		arg(t)
	}

	if t.indent == "" {
		Indent(4)(t)
	}

	if t.comment == "" {
		Comment("//")(t)
	}

	return t
}

type Arg func(*dumper)

func Indent(v int) Arg {
	return func(t *dumper) {
		t.indent = strings.Repeat(" ", v)
	}
}

func Comment(v string) Arg {
	return func(t *dumper) {
		t.comment = v
	}
}
