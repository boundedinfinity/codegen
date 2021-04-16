package loader

import (
	"boundedinfinity/codegen/model"
	"fmt"
	"path"
	"strings"
)

func (t *Loader) addMappedType(k, v string) {
	kc := fmt.Sprintf("%v[]", k)
	vc := fmt.Sprintf("[]%v", v)
	t.typeMap[k] = v
	t.typeMap[kc] = vc
}

func (t *Loader) getMappedType(k string) (string, bool) {
	v, ok := t.typeMap[k]

	if ok {
		return v, true
	}

	for x, v := range t.typeMap {
		if strings.HasSuffix(x, k) {
			return v, true
		}
	}

	return "", false
}

func (t *Loader) getMappedNamespace(k string) (string, bool) {
	var f bool
	ns := model.NAMESPACE_UNKNOWN
	_, ok := t.rtypeMap[k]

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
		if !strings.HasPrefix(ns, t.spec.Name) {
			ns = model.NAMESPACE_BUILTIN
		}
	}

	return ns, f
}

func (t *Loader) currentNamespace() string {
	var name string

	name = t.spec.Name
	name = path.Join(name, path.Join(t.stack.S()...))

	return name
}

func (t *Loader) currentTypeName(typ model.BiSpecType) string {
	var name string

	name = t.currentNamespace()
	name = path.Join(name, typ.Name)

	return name
}

func (t *Loader) relativeNamespace(ns string) string {
	var name string

	name = ns
	name = strings.TrimPrefix(name, t.spec.Name)

	return name
}
