package kind

import (
	"boundedinfinity/codegen/errorer"
	"boundedinfinity/codegen/kind/name"
	"errors"

	"github.com/boundedinfinity/go-commoner/functional/optioner"
)

// //////////////////////////////////////////////////////////////////////////
// Ref Kind
// //////////////////////////////////////////////////////////////////////////

var _ Kind = &RefKind{}

type RefKind struct {
	KindCommon
	Ref Kind
}

func (this RefKind) KindName() name.KindName {
	return name.Ref
}

var (
	ErrRefKindMissing   = errorer.New("ref missing")
	ErrRefKindInvalid   = errorer.New("ref is invalid")
	errRefKindInvalidFn = errorer.ValueFn(ErrRefKindInvalid)
)

func (this RefKind) Validate(config ValidatorConfig) error {
	var errs []error

	if this.Ref == nil {
		errs = append(errs, ErrRefKindMissing)
	}

	return errors.Join(errs...)
}

// //////////////////////////////////////////////////////////////////////////
// Ref Kind Builder
// //////////////////////////////////////////////////////////////////////////

var _ kindBuilder[refKindBuilder] = &refKindBuilder{}

type refKindBuilder struct {
	kind *RefKind
}

func (this *refKindBuilder) Done() Kind {
	return *this.kind
}

func (this *refKindBuilder) Name(v string) *refKindBuilder {
	this.kind.KindCommon.Name = optioner.OfZero(v)
	return this
}

func (this *refKindBuilder) Ref(v Kind) *refKindBuilder {
	this.kind.Ref = v
	return this
}
