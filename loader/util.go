package loader

import (
	"boundedinfinity/codegen/errutil"
	"boundedinfinity/codegen/model"
	"errors"
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

func (t *Loader) reportErr(err error) {
	var path []string
	var cerr errutil.CodeGenError

	if errors.As(err, &cerr) {
		path = append(path, cerr.Path...)
	}

	t.report(path, err.Error())
}

func (t *Loader) report(path []string, format string, a ...interface{}) {
	stack := strings.Join(path, ".")
	reportFormat := fmt.Sprintf("%v: %v\n", stack, format)
	fmt.Printf(reportFormat, a...)
}

func (t *Loader) rootNamespace() string {
	return t.input.Name
}

func (t *Loader) currentNamespace2() string {
	var name string

	name = t.rootNamespace()
	name = path.Join(name, path.Join(t.namespaceStack.S()...))

	return name
}
