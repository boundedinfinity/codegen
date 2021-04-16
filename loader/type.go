package loader

import (
	"boundedinfinity/codegen/model"
)

func (t *Loader) specType2genType(ns model.BiGenNamespace, styp model.BiSpecType) (model.BiGenType, error) {
	gtyp := model.BiGenType{
		Name:       styp.Name,
		Type:       styp.Type,
		Namespace:  ns.QualifiedName,
		Imports:    make([]string, 0),
		Properties: make([]model.BiGenTypeProperty, 0),
		Templates:  make([]model.BiGenTemplate, 0),
	}

	// rns := path.Join(ns.QualifiedName, gtyp.Name)

	// if qns, ok := t.getMappedType(rns); ok {
	// 	gtyp.Type = qns
	// } else {
	// 	return gtyp, fmt.Errorf("not found %v", rns)
	// }

	return gtyp, nil
}

func (t *Loader) genTypeImports(gtyp *model.BiGenType) error {
	if gtyp.Properties != nil {
		ptyps := make(map[string]bool)

		for _, prop := range gtyp.Properties {
			if prop.Namespace != model.NAMESPACE_BUILTIN {
				if ok := ptyps[prop.Namespace]; !ok {
					ptyps[prop.Namespace] = true
				}
			}
		}

		for k := range ptyps {
			gtyp.Imports = append(gtyp.Imports, k)
		}
	}

	return nil
}
