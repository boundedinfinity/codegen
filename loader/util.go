package loader

import (
	"boundedinfinity/codegen/model"
	"fmt"
	"path"
	"strings"
)

func (t *Loader) splitDescription(s string) []string {
	var d []string
	var splitChar string

	if s == "" {
		return d
	}

	if t.input.Info.DescriptionSplitCharacter != "" {
		splitChar = t.input.Info.DescriptionSplitCharacter
	} else {
		splitChar = model.DEFAULT_DESCRIPTION_SPLIT_CHARACTER
	}

	s2 := strings.TrimSuffix(s, splitChar)
	d = strings.Split(s2, splitChar)

	return d
}

func (t *Loader) report(format string, a ...interface{}) {
	stack := strings.Join(t.reportStack.S(), ".")
	reportFormat := fmt.Sprintf("%v: %v\n", stack, format)
	fmt.Printf(reportFormat, a...)
}

func (t *Loader) rootNamespace() string {
	return t.input.Name
}

func (t *Loader) currentNamespace() string {
	var name string

	name = t.rootNamespace()
	name = path.Join(name, path.Join(t.namespaceStack.S()...))

	return name
}

func (t *Loader) isBuiltInType(typ string) bool {
	v := strings.TrimSuffix(typ, model.COLLECTION_SUFFIX)

	for k := range t.builtInTypeMap {
		if k == v {
			return true
		}
	}

	return false
}

func (t *Loader) isCustomType(typ string) bool {
	_, ok := t.customTypeMap[typ]
	return ok
}

func (t *Loader) absoluteNamespace(typ string) string {
	if t.isBuiltInType(typ) {
		return model.NAMESPACE_BUILTIN
	}

	var ns string

	if strings.HasPrefix(typ, "$") {
		ns = typ
	}

	if strings.HasPrefix(typ, "/") {
		ns = path.Join(t.rootNamespace(), typ)
	} else {
		ns = path.Join(t.currentNamespace(), typ)
	}

	return ns
}

func (t *Loader) relativeNamespace(name string) string {
	var ns string

	ns = t.absoluteNamespace(name)
	ns = strings.TrimPrefix(ns, t.rootNamespace())

	return ns
}
