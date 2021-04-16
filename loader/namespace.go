package loader

import (
	"boundedinfinity/codegen/model"
	"fmt"
	"path"
)

func (t *Loader) processNamespace1(ns model.BiSpecNamespace) error {
	t.stack.Push(ns.Name)

	if ns.Types != nil {
		for _, typ := range ns.Types {
			var k string
			var v string

			k = t.currentTypeName(typ)

			v = t.currentNamespace()
			v = path.Base(v)
			v = fmt.Sprintf("%v.%v", v, path.Base(k))

			t.addMappedType(k, v)
		}
	}

	if ns.Namespaces != nil {
		for _, cns := range ns.Namespaces {
			if err := t.processNamespace1(cns); err != nil {
				return err
			}
		}
	}

	t.stack.Pop()
	return nil
}

func (t *Loader) processNamespace2(sns model.BiSpecNamespace, tmpls []model.BiSpecTemplate) (model.BiGenNamespace, error) {
	t.stack.Push(sns.Name)

	gns := model.BiGenNamespace{
		QualifiedName: t.currentNamespace(),
		Types:         make([]model.BiGenType, 0),
		Namespaces:    make([]model.BiGenNamespace, 0),
	}

	if sns.Types != nil {
		for _, styp := range sns.Types {
			gtyp, err := t.specType2genType(gns, styp)

			if err != nil {
				return gns, err
			}

			if styp.Properties != nil {
				for _, sprop := range styp.Properties {
					gprop, err := t.specProperty2genProperty(gns, gtyp, sprop)

					if err != nil {
						return gns, err
					}

					gtyp.Properties = append(gtyp.Properties, gprop)
				}
			}

			if err := t.genTypeImports(&gtyp); err != nil {
				return gns, err
			}

			if tmpls != nil {
				for _, stmpl := range tmpls {
					gtmpl, err := t.processTemplate(gtyp.Namespace, gtyp.Name, stmpl)

					if err != nil {
						return gns, err
					}

					gtyp.Templates = append(gtyp.Templates, gtmpl)
				}
			}

			gns.Types = append(gns.Types, gtyp)
		}
	}

	if sns.Namespaces != nil {
		for _, cns := range sns.Namespaces {
			if gen, err := t.processNamespace2(cns, tmpls); err != nil {
				return gen, err
			} else {
				gns.Namespaces = append(gns.Namespaces, gen)
			}
		}
	}

	t.stack.Pop()
	return gns, nil
}

func (t *Loader) inNamespace(ns model.BiSpecNamespace, typ model.BiSpecTypeProperty) bool {
	// if t.isBuiltInType(typ) {
	// 	return true
	// }

	// if strings.HasPrefix(typ.Type, ns.Name) {
	// 	return true
	// }

	return false
}

func (t *Loader) outOfNamespace(ns model.BiSpecNamespace, typ model.BiSpecTypeProperty) bool {
	return !t.inNamespace(ns, typ)
}

func (t *Loader) inlineType(ns model.BiSpecNamespace, typ model.BiSpecTypeProperty) string {
	var name string
	// var isCollection bool

	// name = typ.Type

	// if strings.HasSuffix(name, model.COLLECTION_SUFFIX) {
	// 	isCollection = true
	// 	name = strings.TrimSuffix(name, model.COLLECTION_SUFFIX)
	// }

	// if t.isUserDefinedType(typ) {
	// 	if t.inNamespace(ns, typ) {
	// 		name = path.Base(name)
	// 	} else {
	// 		name = fmt.Sprintf("%v.%v", path.Dir(name), path.Base(name))
	// 	}
	// }

	// if isCollection {
	// 	name = fmt.Sprintf("[]%v", name)
	// }

	return name
}
