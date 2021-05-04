package loader

import (
	"boundedinfinity/codegen/errutil"
	"errors"
	"fmt"
	"path"
	"strings"
)

func (t *Loader) splitDescription(s string) []string {
	var ss []string
	// var splitChar string

	// if s == "" {
	// 	return ss
	// }

	// if t.inputSpec.Info.DescriptionSplitCharacter != "" {
	// 	splitChar = t.inputSpec.Info.DescriptionSplitCharacter
	// } else {
	// 	splitChar = model.DEFAULT_DESCRIPTION_SPLIT_CHARACTER
	// }

	// s2 := strings.TrimSuffix(s, splitChar)
	// ss = strings.Split(s2, splitChar)

	return ss
}

func (t *Loader) reportErr(err error) {
	var path []string
	var cerr errutil.CodeGenError

	if errors.As(err, &cerr) {
		path = append(path, cerr.Path...)
	}

	t.reportf(path, err.Error())
}

func (t *Loader) reportf(path []string, format string, a ...interface{}) {
	stack := strings.Join(path, ".")
	reportFormat := fmt.Sprintf("%v: %v\n", stack, format)
	fmt.Printf(reportFormat, a...)
}

func (t *Loader) rootName() string {
	var name string

	// name = t.inputSpec.Name

	return name
}

func (t *Loader) appendNamespace(name string) string {
	t.namespaceStack.Push(name)
	return t.getNamespace()
}

func (t *Loader) getNamespace() string {
	var namespace string

	namespace = t.inputSpec.RootPackage()
	namespace = path.Join(namespace, path.Join(t.namespaceStack.S()...))

	return namespace
}
