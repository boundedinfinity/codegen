package loader

import (
	"boundedinfinity/codegen/model"
	"fmt"
	"path"
)

func (t *Loader) specProperty2genProperty(ns model.BiGenNamespace, gtyp model.BiGenType, sprop model.BiSpecTypeProperty) (model.BiGenTypeProperty, error) {
	gprop := model.BiGenTypeProperty{
		Name: sprop.Name,
	}

	if typ, ok := t.getMappedType(sprop.Type); ok {
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
