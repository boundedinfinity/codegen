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
