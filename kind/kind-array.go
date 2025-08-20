package kind

import (
	"boundedinfinity/codegen/errorer"
	"boundedinfinity/codegen/kind/name"
	"errors"

	"github.com/boundedinfinity/go-commoner/functional/optioner"
)

// //////////////////////////////////////////////////////////////////////////
// Array Kind
// //////////////////////////////////////////////////////////////////////////

var _ Kind = &ArrayKind{}

type ArrayKind struct {
	KindCommon
	Items Kind
	Min   optioner.Option[int]
	Max   optioner.Option[int]
}

func (this ArrayKind) KindName() name.KindName {
	return name.Array
}

var (
	ErrArrayKindMinGreaterThanMax   = errorer.New("min is greater than max")
	errArrayKindMinGreaterThanMaxFn = errorer.ValueFnf(ErrIntegerKindMinGreaterThanMax, "%v max, %v min")
	ErrArrayKindItemsMissing        = errorer.New("items is missing")
)

func (this ArrayKind) Validate(config ValidatorConfig) error {
	var errs []error

	if this.Max.Defined() && this.Min.Defined() && this.Min.Get() > this.Max.Get() {
		errs = append(errs, errIntegerKindMinGreaterThanMaxFn(this.Max.Get(), this.Min.Get()))
	}

	if this.Items == nil {
		errs = append(errs, ErrArrayKindItemsMissing)
	}

	errs = append(errs, this.Items.Validate(ValidatorConfig{}))

	return errors.Join(errs...)
}

// //////////////////////////////////////////////////////////////////////////
// Array Kind Builder
// //////////////////////////////////////////////////////////////////////////

var _ kindBuilder[arrayKindBuilder] = &arrayKindBuilder{}

type arrayKindBuilder struct {
	kind *ArrayKind
}

func (this *arrayKindBuilder) Done() Kind {
	return *this.kind
}

func (this *arrayKindBuilder) Name(v string) *arrayKindBuilder {
	this.kind.KindCommon.Name = optioner.Some(v)
	return this
}

func (this *arrayKindBuilder) Type(v string) *arrayKindBuilder {
	this.kind.KindCommon.Type = optioner.Some(v)
	return this
}

func (this *arrayKindBuilder) Items(v Kind) *arrayKindBuilder {
	this.kind.Items = v
	return this
}

func (this *arrayKindBuilder) Min(v int) *arrayKindBuilder {
	this.kind.Min = optioner.Some(v)
	return this
}

func (this *arrayKindBuilder) Max(v int) *arrayKindBuilder {
	this.kind.Max = optioner.Some(v)
	return this
}
