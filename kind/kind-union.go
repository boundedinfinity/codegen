package kind

import (
	"boundedinfinity/codegen/errorer"
	"boundedinfinity/codegen/kind/name"
	"errors"

	"github.com/boundedinfinity/go-commoner/functional/optioner"
)

// //////////////////////////////////////////////////////////////////////////
// Union Kind
// //////////////////////////////////////////////////////////////////////////

var _ Kind = &UnionKind{}

type UnionKind struct {
	KindCommon
	Refs []Kind
}

func (this UnionKind) KindName() name.KindName {
	return name.Union
}

var (
	ErrUnionKindMissing   = errorer.New("union missing")
	ErrUnionKindInvalid   = errorer.New("union is invalid")
	errUnionKindInvalidFn = errorer.ValueFn(ErrUnionKindInvalid)
)

func (this UnionKind) Validate(config ValidatorConfig) error {
	var errs []error

	if this.Refs == nil {
		errs = append(errs, ErrUnionKindMissing)
	}

	for _, ref := range this.Refs {
		if ref == nil {
			errs = append(errs, ErrUnionKindMissing)
		}
	}

	return errors.Join(errs...)
}

// //////////////////////////////////////////////////////////////////////////
// Union Kind Builder
// //////////////////////////////////////////////////////////////////////////

var _ kindBuilder[unionKindBuilder] = &unionKindBuilder{}

type unionKindBuilder struct {
	kind *UnionKind
}

func (this *unionKindBuilder) Done() Kind {
	return *this.kind
}

func (this *unionKindBuilder) Name(v string) *unionKindBuilder {
	this.kind.KindCommon.Name = optioner.OfZero(v)
	return this
}

func (this *unionKindBuilder) Refs(v []Kind) *unionKindBuilder {
	this.kind.Refs = append(this.kind.Refs, v...)
	return this
}

func (this *unionKindBuilder) Ref(v Kind) *unionKindBuilder {
	return this.Refs([]Kind{v})
}
