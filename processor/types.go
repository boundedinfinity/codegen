package processor

import (
	"boundedinfinity/codegen/model"

	"github.com/boundedinfinity/go-commoner/functional/optioner"
	"github.com/boundedinfinity/go-commoner/idiomatic/caser"
	"github.com/boundedinfinity/go-commoner/idiomatic/pather"
)

func (t *Processor) processTypes() error {
	for _, project := range t.projects {
		for _, typ := range project.Types {
			if _, ok := t.TypeMap[typ.Common().Id.Get()]; ok {
				return ErrCodeGenTypeSchemaIdDuplicateFn(typ)
			}

			t.TypeMap[typ.Common().Id.Get()] = typ
			t.Combined.Types = append(t.Combined.Types, typ)
		}
	}

	for _, typ := range t.Combined.Types {
		if err := t.checkType(typ); err != nil {
			return err
		}
	}

	for _, typ := range t.Combined.Types {
		switch typ.(type) {
		case *model.CodeGenRef:
		case *model.CodeGenArray:
		case *model.CodeGenObject:
		default:
			processType(t.Combined, typ)
		}
	}

	for _, typ := range t.Combined.Types {
		switch typ.(type) {
		case *model.CodeGenRef:
			processType(t.Combined, typ)
		default:
		}
	}

	for _, typ := range t.Combined.Types {
		switch typ.(type) {
		case *model.CodeGenArray:
			processType(t.Combined, typ)
		case *model.CodeGenObject:
			processType(t.Combined, typ)
		default:
		}
	}

	return nil
}

func processType(project model.CodeGenProject, typ model.CodeGenType) {
	switch rtyp := typ.(type) {
	case *model.CodeGenObject:
		for _, prop := range rtyp.Properties.Get() {
			processType(project, prop)
		}
	case *model.CodeGenArray:
		processType(project, rtyp.Items.Get())
	default:
		if typ.Common().Name.Empty() {
			name := typ.Common().Id.Get()
			name = pather.Paths.Base(name)
			name = caser.KebabToPascal(name)
			typ.Common().Name = optioner.Some(name)
		}

		if typ.Common().Package.Empty() {
			pkg := typ.Common().Id.Get()
			pkg = pather.Paths.Dir(pkg)
			pkg = pather.Paths.Base(pkg)
			typ.Common().Package = optioner.Some(pkg)
		}

		if typ.Common().ImportPath.Empty() {
			pkg := typ.Common().Id.Get()
			pkg = pather.Paths.Dir(pkg)
			pkg = pather.Paths.Join(project.Package.Get(), pkg)
			typ.Common().ImportPath = optioner.Some(pkg)
		}

		if typ.Common().JsonName.Empty() {
			name := typ.Common().Id.Get()
			name = pather.Paths.Base(name)
			typ.Common().JsonName = optioner.Some(name)
		}

		if typ.Common().YamlName.Empty() {
			name := typ.Common().Id.Get()
			name = pather.Paths.Base(name)
			typ.Common().YamlName = optioner.Some(name)
		}

		if typ.Common().SqlName.Empty() {
			name := typ.Common().Id.Get()
			name = pather.Paths.Base(name)
			name = caser.KebabToSnake(name)
			typ.Common().SqlName = optioner.Some(name)
		}
	}
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

		if target, ok := t.TypeMap[obj.Ref.Get()]; !ok {
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
