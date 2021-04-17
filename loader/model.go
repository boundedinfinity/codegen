package loader

import (
	"boundedinfinity/codegen/model"
	"fmt"
	"path"
)

func (t *Loader) specModel2genModel(ns model.BiOutput_Model_Namespace, styp model.BiInput_Model) (model.BiOutput_Model, error) {
	gtyp := model.BiOutput_Model{
		Name:       styp.Name,
		Type:       styp.Type,
		Namespace:  ns.Name,
		Imports:    make([]string, 0),
		Properties: make([]model.BiOutput_TypeProperty, 0),
		Templates:  make([]model.BiOutput_Template, 0),
	}

	return gtyp, nil
}

func (t *Loader) genModelImports(gtyp *model.BiOutput_Model) error {
	if gtyp.Properties != nil {
		ptyps := make(map[string]bool)

		for _, prop := range gtyp.Properties {
			if prop.Namespace != model.NAMESPACE_BUILTIN && gtyp.Namespace != prop.Namespace {
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

func (t *Loader) specProperty2genProperty(ns model.BiOutput_Model_Namespace, gtyp model.BiOutput_Model, sprop model.BiInput_Model_Property) (model.BiOutput_TypeProperty, error) {
	gprop := model.BiOutput_TypeProperty{
		Name: sprop.Name,
	}

	if typ, ok := t.getMappedType(ns.Name, sprop.Type); ok {
		gprop.Type = typ
	} else {
		return gprop, fmt.Errorf("type not found %v", sprop.Type)
	}

	if tns, ok := t.getMappedNamespace(sprop.Type); ok {
		if tns != model.NAMESPACE_BUILTIN {
			gprop.Namespace = path.Dir(tns)
		} else {
			gprop.Namespace = tns
		}
	} else {
		return gprop, fmt.Errorf("namespace not found %v", sprop.Type)
	}

	return gprop, nil
}
