package loader

import (
	"boundedinfinity/codegen/model"
	"fmt"
	"path"
)

func (t *Loader) specOperation2genOperation(ns model.BiOutput_Operation_Namespace, styp model.BiInput_Operation) (model.BiOutput_Operation, error) {
	gtyp := model.BiOutput_Operation{
		Name:      styp.Name,
		Namespace: ns.Name,
		Inputs:    make([]model.BiOutput_TypeProperty, 0),
		Outputs:   make([]model.BiOutput_TypeProperty, 0),
		Templates: make([]model.BiOutput_Template, 0),
	}

	return gtyp, nil
}

func (t *Loader) genOperationImports(gtyp *model.BiOutput_Operation) error {
	ptyps := make(map[string]bool)

	if gtyp.Inputs != nil {
		for _, input := range gtyp.Inputs {
			if input.Namespace != model.NAMESPACE_BUILTIN && input.Namespace != gtyp.Namespace {
				if ok := ptyps[input.Namespace]; !ok {
					ptyps[input.Namespace] = true
				}
			}
		}
	}

	for k := range ptyps {
		gtyp.Imports = append(gtyp.Imports, k)
	}

	return nil
}

func (t *Loader) specIO2genProperty(ns model.BiOutput_Operation_Namespace, gtyp model.BiOutput_Operation, io string) (model.BiOutput_TypeProperty, error) {
	gprop := model.BiOutput_TypeProperty{}

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
