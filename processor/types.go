package processor

import (
	"boundedinfinity/codegen/model"
	"errors"
	"fmt"
)

func (t *Processor) processTypes() error {
	for _, project := range t.projects {
		for _, typ := range project.Types {
			if _, ok := t.typeIdMap[typ.TypeId().Get()]; ok {
				return ErrCodeGenTypeSchemaIdDuplicateFn(typ)
			}

			t.typeIdMap[typ.TypeId().Get()] = typ
			t.combined.Types = append(t.combined.Types, typ)
		}
	}

	return nil
}

func (t *Processor) resolveType(typ model.CodeGenType) (model.CodeGenType, error) {
	var found model.CodeGenType
	var err error

	switch obj := typ.(type) {
	case *model.CodeGenRef:
		if obj.Ref.Empty() {
			err = errors.New("invalid reference")
			break
		}

		if target, ok := t.typeIdMap[obj.Ref.Get()]; !ok {
			err = fmt.Errorf("invalid reference %v", obj.Ref.Get())
		} else {
			found = target
		}
	case *model.CodeGenBoolean, *model.CodeGenEnum, *model.CodeGenFloat,
		*model.CodeGenInteger, *model.CodeGenString, *model.CodeGenArray,
		*model.CodeGenObject:
		found = typ
	default:
		err = ErrCodeGenUnsupportedTypeFn(typ)
	}

	return found, err
}

func (t *Processor) checkCombinedTypes() error {
	for _, typ := range t.combined.Types {
		if err := t.checkType(typ); err != nil {
			return err
		}
	}

	return nil
}

func (t *Processor) checkType(typ model.CodeGenType) error {
	switch obj := typ.(type) {
	case *model.CodeGenRef:
		if obj.Ref.Empty() {
			return errors.New("invalid reference")
		}

		if _, ok := t.typeIdMap[obj.Ref.Get()]; !ok {
			return fmt.Errorf("invalid reference %v", obj.Ref.Get())
		}
	case *model.CodeGenArray:
		return t.checkType(obj.Items)
	case *model.CodeGenObject:
		for _, prop := range obj.Properties {
			if err := t.checkType(prop); err != nil {
				return err
			}
		}
	case *model.CodeGenBoolean, *model.CodeGenEnum, *model.CodeGenFloat,
		*model.CodeGenInteger, *model.CodeGenString:
		// Nothing to do
	default:
		return ErrCodeGenUnsupportedTypeFn(typ)
	}

	return nil
}
