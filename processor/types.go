package processor

import (
	"boundedinfinity/codegen/model"

	"github.com/boundedinfinity/go-commoner/idiomatic/pather"
	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
)

func (t *Processor) processTypes() error {
	for _, project := range t.projects {
		for _, typ := range project.Types {
			if _, ok := t.typeIdMap[typ.Common().Id.Get()]; ok {
				return ErrCodeGenTypeSchemaIdDuplicateFn(typ)
			}

			t.typeIdMap[typ.Common().Id.Get()] = typ
			t.combined.Types = append(t.combined.Types, typ)
		}
	}

	for _, typ := range t.combined.Types {
		if err := t.checkType(typ); err != nil {
			return err
		}
	}

	translations := map[string]string{}

	for _, typ := range t.combined.Types {
		if err := t.processTypePackageAndQualifiedName(typ, translations); err != nil {
			return err
		}
	}

	return nil
}

func (t *Processor) processTypePackageAndQualifiedName(typ model.CodeGenType, translations map[string]string) error {
	var qualifedName string

	if typ.Common().Package.Defined() {
		qualifedName = typ.Common().Package.Get()
	} else {
		if t.combined.Package.Defined() {
			qualifedName = t.combined.Package.Get()
		}

		qualifedName = pather.Paths.Join(qualifedName, typ.Common().Id.Get())
	}

	for from, to := range translations {
		qualifedName = stringer.Replace(qualifedName, to, from)
	}

	// packageName := pather.Paths.Dir(qualifedName)
	// typ.Common().QName_ = optioner.Some(qualifedName)
	// typ.Common().Package = optioner.Some(packageName)

	return nil
}

func (t *Processor) resolveType(typ model.CodeGenType) (model.CodeGenType, error) {
	var found model.CodeGenType
	var err error

	switch obj := typ.(type) {
	case *model.CodeGenRef:
		if obj.Ref.Empty() {
			err = model.ErrRefEmpty
			break
		}

		if target, ok := t.typeIdMap[obj.Ref.Get()]; !ok {
			err = model.ErrRefNotFound.WithValue(obj.Ref.Get())
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
	var err error
	var found model.CodeGenType

	switch obj := typ.(type) {
	case *model.CodeGenRef:
		found, err = t.resolveType(typ)
		obj.Resolved = found
	case *model.CodeGenArray:
		_, err = t.resolveType(obj.Items.Get())
	case *model.CodeGenObject:
		for _, prop := range obj.Properties.Get() {
			err = t.checkType(prop)

			if err != nil {
				break
			}
		}
	case *model.CodeGenBoolean, *model.CodeGenEnum, *model.CodeGenFloat,
		*model.CodeGenInteger, *model.CodeGenString:
		// Nothing to do
	default:
		return ErrCodeGenUnsupportedTypeFn(typ)
	}

	return err
}
