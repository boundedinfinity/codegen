package processor

import (
	"boundedinfinity/codegen/model"
)

func (this *Processor) processTypes() error {
	for _, project := range this.projects {
		for _, typ := range project.Types {
			if _, ok := this.TypeMap[typ.Common().Id.Get()]; ok {
				return ErrCodeGenTypeSchemaIdDuplicateFn(typ)
			}

			this.TypeMap[typ.Common().Id.Get()] = typ
			this.Combined.Types = append(this.Combined.Types, typ)
		}
	}

	for _, typ := range this.Combined.Types {
		if err := this.checkType(typ); err != nil {
			return err
		}
	}

	for _, typ := range this.Combined.Types {
		switch typ.(type) {
		case *model.CodeGenRef:
		case *model.CodeGenArray:
		case *model.CodeGenObject:
		default:
			processType(this.Combined, typ)
		}
	}

	for _, typ := range this.Combined.Types {
		switch typ.(type) {
		case *model.CodeGenRef:
			processType(this.Combined, typ)
		default:
		}
	}

	for _, typ := range this.Combined.Types {
		switch typ.(type) {
		case *model.CodeGenArray:
			processType(this.Combined, typ)
		case *model.CodeGenObject:
			processType(this.Combined, typ)
		default:
		}
	}

	return nil
}

func processType(project model.CodeGenProject, typ model.CodeGenSchema) {
	switch rtyp := typ.(type) {
	case *model.CodeGenObject:
		model.EnsureName(typ)
		model.EnsurePackage(typ)
		model.EnsureJsonName(typ)
		model.EnsureYamlName(typ)
		model.EnsureSqlName(typ)

		for _, prop := range rtyp.Properties.Get() {
			processType(project, prop)
		}
	case *model.CodeGenArray:
		processType(project, rtyp.Items.Get())

		model.EnsureName(typ)
		model.EnsurePackage(typ)
		model.EnsureJsonName(typ)
		model.EnsureYamlName(typ)
		model.EnsureSqlName(typ)
	default:
		model.EnsureName(typ)
		model.EnsurePackage(typ)
		model.EnsureJsonName(typ)
		model.EnsureYamlName(typ)
		model.EnsureSqlName(typ)
	}
}

func (this *Processor) resolveType(typ model.CodeGenSchema) (model.CodeGenSchema, error) {
	var found model.CodeGenSchema
	var err error

	switch obj := typ.(type) {
	case *model.CodeGenRef:
		if obj.Ref.Empty() {
			err = model.ErrRefEmpty
			break
		}

		if target, ok := this.TypeMap[obj.Ref.Get()]; !ok {
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

func (this *Processor) checkType(typ model.CodeGenSchema) error {
	var err error
	var found model.CodeGenSchema

	switch obj := typ.(type) {
	case *model.CodeGenRef:
		found, err = this.resolveType(typ)
		obj.Resolved = found
	case *model.CodeGenArray:
		_, err = this.resolveType(obj.Items.Get())
	case *model.CodeGenObject:
		for _, prop := range obj.Properties.Get() {
			err = this.checkType(prop)

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
