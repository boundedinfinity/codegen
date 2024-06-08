package processor

import (
	"boundedinfinity/codegen/model"
	"errors"
	"fmt"
)

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
