package loader

import (
	"boundedinfinity/codegen/model"
	"boundedinfinity/codegen/util"
	"fmt"
	"path"
	"strings"
)

func (t *Loader) report(format string, a ...interface{}) {
	stack := strings.Join(t.reportStack.S(), ".")
	reportFormat := fmt.Sprintf("%v: %v\n", stack, format)
	fmt.Printf(reportFormat, a...)
}

func (t *Loader) wrapError(err error) error {
	stack := strings.Join(t.reportStack.S(), ".")
	return fmt.Errorf("%v: %w", stack, err)
}

func (t *Loader) addMappedType(k string, v TypeInfo) {
	kc := fmt.Sprintf("%v[]", k)
	vc := TypeInfo{
		InNamespace:  fmt.Sprintf("[]%v", v.InNamespace),
		OutNamespace: fmt.Sprintf("[]%v", v.OutNamespace),
		Namespace:    v.Namespace,
	}
	t.typeMap[k] = v
	t.typeMap[kc] = vc

	t.report("mapping %v -> %v", util.SummerySuffix(k, model.SUMMERY_SIZE), v)
	t.report("mapping %v -> %v", util.SummerySuffix(kc, model.SUMMERY_SIZE), vc)
}

func (t *Loader) getMappedType(ns, typ string) (string, bool) {
	v, ok := t.typeMap[typ]

	if ok {
		if ns == v.Namespace {
			return v.InNamespace, ok
		} else {
			return v.OutNamespace, ok
		}
	}

	for x, v := range t.typeMap {
		if strings.HasSuffix(x, typ) {
			if v.Namespace == ns {
				return v.InNamespace, true
			} else {
				return v.OutNamespace, true
			}
		}
	}

	return "", false
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
		if !strings.HasPrefix(ns, t.input.Name) {
			ns = model.NAMESPACE_BUILTIN
		}
	}

	return ns, f
}

func (t *Loader) currentNamespace() string {
	var name string

	name = t.input.Name
	name = path.Join(name, path.Join(t.modelStack.S()...))

	return name
}

func (t *Loader) currentModelName(m model.BiInput_Model) string {
	var name string

	name = t.currentNamespace()
	name = path.Join(name, m.Name)

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

	return name
}
