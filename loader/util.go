package loader

import (
	"boundedinfinity/codegen/model"
	"boundedinfinity/codegen/util"
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

func (t *Loader) wrapError(err error) error {
	stack := strings.Join(t.reportStack.S(), ".")
	return fmt.Errorf("%v: %w", stack, err)
}

func (t *Loader) addUserMappedType(bname string) {
	ns := t.currentNamespace()
	qname := path.Join(ns, bname)

	iname := qname
	iname = path.Dir(iname)
	iname = path.Base(iname)
	iname = fmt.Sprintf("%v.%v", iname, bname)

	t.addMappedType(qname, TypeInfo{
		BaseName:   bname,
		QName:      qname,
		ImportName: iname,
		Namespace:  ns,
		BuiltIn:    false,
	})
}

func (t *Loader) addMappedType(k string, v TypeInfo) {
	kc := fmt.Sprintf("%v[]", k)
	vc := TypeInfo{
		BaseName:   fmt.Sprintf("[]%v", v.BaseName),
		ImportName: fmt.Sprintf("[]%v", v.ImportName),
		QName:      fmt.Sprintf("[]%v", v.QName),
		Namespace:  v.Namespace,
		BuiltIn:    v.BuiltIn,
	}
	t.typeMap[k] = v
	t.typeMap[kc] = vc

	t.report("mapping %v -> %v", util.SummerySuffix(k, model.SUMMERY_SIZE), v.ImportName)
	t.report("mapping %v -> %v", util.SummerySuffix(kc, model.SUMMERY_SIZE), vc.ImportName)
}

func (t *Loader) getMappedType(typ string) (TypeInfo, bool) {
	if tf, ok := t.typeMap[typ]; ok {
		return tf, ok
	}

	n1 := path.Join(t.rootPackage(), typ)

	if tf, ok := t.typeMap[n1]; ok {
		return tf, ok
	}

	n2 := path.Join(t.currentNamespace(), typ)

	if tf, ok := t.typeMap[n2]; ok {
		return tf, ok
	}

	return TypeInfo{}, false
}

func (t *Loader) getMappedNamespace(k string) (string, bool) {
	var f bool
	ns := model.NAMESPACE_UNKNOWN
	_, ok := t.typeMap[k]

	if ok {
		ns = k
	} else {
		for x := range t.typeMap {
			if strings.HasSuffix(x, k) {
				ns = x
				break
			}
		}
	}

	if ns != model.NAMESPACE_UNKNOWN {
		f = true
		if !strings.HasPrefix(ns, t.rootPackage()) {
			ns = model.NAMESPACE_BUILTIN
		}
	}

	return ns, f
}

func (t *Loader) rootPackage() string {
	return t.input.Name
}

func (t *Loader) currentNamespace() string {
	var name string

	name = t.rootPackage()
	name = path.Join(name, path.Join(t.modelStack.S()...))

	return name
}

func (t *Loader) currentOperationName(m model.BiInput_Operation) string {
	var name string

	name = t.currentNamespace()
	name = path.Join(name, m.Name)

	return name
}

func (t *Loader) relativeNamespace(ns string) string {
	var name string

	name = ns
	name = strings.TrimPrefix(name, t.Output.Name)
	name = strings.TrimPrefix(name, "/")

	return name
}
