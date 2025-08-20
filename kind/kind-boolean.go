package kind

import (
	"boundedinfinity/codegen/kind/name"

	"github.com/boundedinfinity/go-commoner/functional/optioner"
)

// //////////////////////////////////////////////////////////////////////////
// Boolean Kind
// //////////////////////////////////////////////////////////////////////////

var _ Kind = &BooleanKind{}

type BooleanKind struct {
	KindCommon
}

func (this BooleanKind) KindName() name.KindName {
	return name.Boolean
}

func (this BooleanKind) Validate(config ValidatorConfig) error {
	return nil
}

// //////////////////////////////////////////////////////////////////////////
// Boolean Kind Builder
// //////////////////////////////////////////////////////////////////////////

var _ kindBuilder[booleanKindBuilder] = &booleanKindBuilder{}

type booleanKindBuilder struct {
	kind *BooleanKind
}

func (this *booleanKindBuilder) Done() Kind {
	return *this.kind
}

func (this *booleanKindBuilder) Name(v string) *booleanKindBuilder {
	this.kind.KindCommon.Name = optioner.Some(v)
	return this
}

func (this *booleanKindBuilder) Type(v string) *booleanKindBuilder {
	this.kind.KindCommon.Type = optioner.Some(v)
	return this
}
