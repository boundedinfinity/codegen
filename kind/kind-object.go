package kind

import (
	"boundedinfinity/codegen/errorer"
	"boundedinfinity/codegen/kind/name"
	"errors"

	"github.com/boundedinfinity/go-commoner/functional/optioner"
	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
)

// //////////////////////////////////////////////////////////////////////////
// Object Kind
// //////////////////////////////////////////////////////////////////////////

var _ Kind = &ObjectKind{}

type ObjectKind struct {
	KindCommon
	Properties []Kind
}

func (this ObjectKind) KindName() name.KindName {
	return name.Object
}

var (
	ErrObjectKindConflictingNames   = errorer.New("property names conflicting")
	errObjectKindConflictingNamesFn = errorer.ValueFnf(ErrObjectKindConflictingNames, "conflicting name [%v]")
)

func (this ObjectKind) Validate(config ValidatorConfig) error {
	var errs []error
	nameCount := map[string]int{}

	for _, property := range this.Properties {
		switch kind := property.(type) {
		case *StringKind:
			if kind.KindCommon.Name.Defined() {
				nameCount[kind.KindCommon.Name.Get()]++
			}
		}
	}

	var duplicateNames []string

	for name, count := range nameCount {
		if count > 1 {
			duplicateNames = append(duplicateNames, name)
		}
	}

	if len(duplicateNames) > 0 {
		value := stringer.Join(", ", duplicateNames...)
		errs = append(errs, errObjectKindConflictingNamesFn(value))
	}

	for _, property := range this.Properties {
		errs = append(errs, property.Validate(ValidatorConfig{}))
	}

	return errors.Join(errs...)
}

// //////////////////////////////////////////////////////////////////////////
// Object Kind Builder
// //////////////////////////////////////////////////////////////////////////

var _ kindBuilder[objectKindBuilder] = &objectKindBuilder{}

type objectKindBuilder struct {
	kind *ObjectKind
}

func (this *objectKindBuilder) Done() Kind {
	return *this.kind
}

func (this *objectKindBuilder) Name(v string) *objectKindBuilder {
	this.kind.KindCommon.Name = optioner.OfZero(v)
	return this
}

func (this *objectKindBuilder) Type(v string) *objectKindBuilder {
	this.kind.KindCommon.Type = optioner.OfZero(v)
	return this
}

func (this *objectKindBuilder) Property(v Kind) *objectKindBuilder {
	this.kind.Properties = append(this.kind.Properties, v)
	return this
}

func (this *objectKindBuilder) Properties(v []Kind) *objectKindBuilder {
	this.kind.Properties = append(this.kind.Properties, v...)
	return this
}
