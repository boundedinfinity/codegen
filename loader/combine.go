package loader

import (
	"boundedinfinity/codegen/model"
	"boundedinfinity/codegen/util"
)

func (t *Loader) combine(inputPath string) error {
	var input model.InputFile

	if err := util.UnmarshalFromFile(inputPath, &input); err != nil {
		return err
	}

	for k, v := range input.Info.Primitives {
		if _, ok := t.dependencies[k]; !ok {
			return t.ErrInvalidPrimitive(k)
		}

		if p, ok := t.primitiveMap[k]; ok {
			if p != "" {
				return t.ErrDuplicatePrimitive(k)
			} else {
				t.primitiveMap[k] = v
			}
		} else {
			return t.ErrInvalidPrimitive(k)
		}
	}

	for _, m := range input.Specification.Models {
		if _, ok := t.inputModels[m.Name]; ok {
			return t.ErrDuplicateType(m.Name)
		} else {
			t.inputModels[m.Name] = m
		}
	}

	return nil
}
