package kind

import (
	"boundedinfinity/codegen/errorer"
	"boundedinfinity/codegen/kind/name"
	"errors"

	"github.com/boundedinfinity/go-commoner/functional/optioner"
	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
)

// //////////////////////////////////////////////////////////////////////////
// Integer Kind
// //////////////////////////////////////////////////////////////////////////

var _ Kind = &IntegerKind{}

type IntegerKind struct {
	KindCommon
	Min        optioner.Option[int]
	Max        optioner.Option[int]
	MultipleOf optioner.Option[int]
	Positive   optioner.Option[bool]
	Negative   optioner.Option[bool]
	OneOf      optioner.Option[[]int]
	NoneOf     optioner.Option[[]int]
}

func (this IntegerKind) KindName() name.KindName {
	return name.Integer
}

var (
	ErrIntegerKindMinGreaterThanMax     = errorer.New("min is greater than max")
	errIntegerKindMinGreaterThanMaxFn   = errorer.ValueFnf(ErrIntegerKindMinGreaterThanMax, "%v max, %v min")
	ErrIntegerKindPositiveAndNegative   = errorer.New("positive and negative at the same time")
	ErrIntegerKindMultipleOfZero        = errorer.New("multiple-of is zero")
	ErrIntegerKindOneOfOverlapsNoneOf   = errorer.New("one-of overlaps with none-of")
	errIntegerKindOneOfOverlapsNoneOfFn = errorer.ValueFnf(ErrIntegerKindOneOfOverlapsNoneOf, "overlapping values [%v]")
)

func (this IntegerKind) Validate(config ValidatorConfig) error {
	var errs []error

	if this.Max.Defined() && this.Min.Defined() && this.Min.Get() > this.Max.Get() {
		errs = append(errs, errIntegerKindMinGreaterThanMaxFn(this.Max.Get(), this.Min.Get()))
	}

	if this.Positive.Defined() && this.Negative.Defined() {
		errs = append(errs, ErrIntegerKindPositiveAndNegative)
	}

	if this.MultipleOf.Defined() && this.MultipleOf.Get() == 0 {
		errs = append(errs, ErrIntegerKindMultipleOfZero)
	}

	if this.OneOf.Defined() && this.NoneOf.Defined() {
		var overlaps []int

		for _, oneOf := range this.OneOf.Get() {
			for _, noneOf := range this.NoneOf.Get() {
				if oneOf == noneOf {
					overlaps = append(overlaps, oneOf)
				}
			}
		}

		if len(overlaps) > 0 {
			value := stringer.Join(", ", overlaps...)
			errs = append(errs, errIntegerKindOneOfOverlapsNoneOfFn(value))
		}
	}

	return errors.Join(errs...)
}

// //////////////////////////////////////////////////////////////////////////
// Integer Kind Builder
// //////////////////////////////////////////////////////////////////////////

var _ kindBuilder[integerKindBuilder] = &integerKindBuilder{}

type integerKindBuilder struct {
	kind *IntegerKind
}

func (this *integerKindBuilder) Done() Kind {
	return *this.kind
}

func (this *integerKindBuilder) Name(v string) *integerKindBuilder {
	this.kind.KindCommon.Name = optioner.OfZero(v)
	return this
}

func (this *integerKindBuilder) Type(v string) *integerKindBuilder {
	this.kind.KindCommon.Type = optioner.OfZero(v)
	return this
}

func (this *integerKindBuilder) Min(v int) *integerKindBuilder {
	this.kind.Min = optioner.Some(v)
	return this
}

func (this *integerKindBuilder) Max(v int) *integerKindBuilder {
	this.kind.Max = optioner.Some(v)
	return this
}

func (this *integerKindBuilder) MultipleOf(v int) *integerKindBuilder {
	this.kind.MultipleOf = optioner.Some(v)
	return this
}

func (this *integerKindBuilder) Positive(v bool) *integerKindBuilder {
	this.kind.Positive = optioner.Some(v)
	return this
}

func (this *integerKindBuilder) Negative(v bool) *integerKindBuilder {
	this.kind.Negative = optioner.Some(v)
	return this
}

func (this *integerKindBuilder) OneOf(v []int) *integerKindBuilder {
	this.kind.OneOf = optioner.OfSlice(v)
	return this
}

func (this *integerKindBuilder) NoneOf(v []int) *integerKindBuilder {
	this.kind.NoneOf = optioner.OfSlice(v)
	return this
}

// //////////////////////////////////////////////////////////////////////////
// Float Kind
// //////////////////////////////////////////////////////////////////////////

var _ Kind = &FloatKind{}

type FloatKind struct {
	KindCommon
	Min        optioner.Option[float64]
	Max        optioner.Option[float64]
	MultipleOf optioner.Option[float64]
	Positive   optioner.Option[bool]
	Negative   optioner.Option[bool]
	OneOf      optioner.Option[[]float64]
	NoneOf     optioner.Option[[]float64]
}

func (this FloatKind) KindName() name.KindName {
	return name.Float
}

var (
	ErrFloatKindMinGreaterThanMax     = errorer.New("min is greater than max")
	errFloatKindMinGreaterThanMaxFn   = errorer.ValueFnf(ErrFloatKindMinGreaterThanMax, "%v max, %v min")
	ErrFloatKindPositiveAndNegative   = errorer.New("positive and negative at the same time")
	ErrFloatKindMultipleOfZero        = errorer.New("multiple-of is zero")
	ErrFloatKindOneOfOverlapsNoneOf   = errorer.New("one-of overlaps with none-of")
	errFloatKindOneOfOverlapsNoneOfFn = errorer.ValueFnf(ErrFloatKindOneOfOverlapsNoneOf, "overlapping values [%v]")
)

func (this FloatKind) Validate(config ValidatorConfig) error {
	var errs []error

	if this.Max.Defined() && this.Min.Defined() && this.Min.Get() > this.Max.Get() {
		errs = append(errs, errFloatKindMinGreaterThanMaxFn(this.Max.Get(), this.Min.Get()))
	}

	if this.Positive.Defined() && this.Negative.Defined() {
		errs = append(errs, ErrFloatKindPositiveAndNegative)
	}

	if this.MultipleOf.Defined() && this.MultipleOf.Get() == 0 {
		errs = append(errs, ErrFloatKindMultipleOfZero)
	}

	if this.OneOf.Defined() && this.NoneOf.Defined() {
		var overlaps []float64

		for _, oneOf := range this.OneOf.Get() {
			for _, noneOf := range this.NoneOf.Get() {
				if oneOf == noneOf {
					overlaps = append(overlaps, oneOf)
				}
			}
		}

		if len(overlaps) > 0 {
			value := stringer.Join(", ", overlaps...)
			errs = append(errs, errFloatKindOneOfOverlapsNoneOfFn(value))
		}
	}

	return errors.Join(errs...)
}

// //////////////////////////////////////////////////////////////////////////
// Float Kind Builder
// //////////////////////////////////////////////////////////////////////////

var _ kindBuilder[floatKindBuilder] = &floatKindBuilder{}

type floatKindBuilder struct {
	kind *FloatKind
}

func (this *floatKindBuilder) Done() Kind {
	return *this.kind
}

func (this *floatKindBuilder) Name(v string) *floatKindBuilder {
	this.kind.KindCommon.Name = optioner.OfZero(v)
	return this
}

func (this *floatKindBuilder) Type(v string) *floatKindBuilder {
	this.kind.KindCommon.Type = optioner.OfZero(v)
	return this
}

func (this *floatKindBuilder) Min(v float64) *floatKindBuilder {
	this.kind.Min = optioner.Some(v)
	return this
}

func (this *floatKindBuilder) Max(v float64) *floatKindBuilder {
	this.kind.Max = optioner.Some(v)
	return this
}

func (this *floatKindBuilder) MultipleOf(v float64) *floatKindBuilder {
	this.kind.MultipleOf = optioner.Some(v)
	return this
}

func (this *floatKindBuilder) Positive(v bool) *floatKindBuilder {
	this.kind.Positive = optioner.Some(v)
	return this
}

func (this *floatKindBuilder) Negative(v bool) *floatKindBuilder {
	this.kind.Negative = optioner.Some(v)
	return this
}

func (this *floatKindBuilder) OneOf(v []float64) *floatKindBuilder {
	this.kind.OneOf = optioner.OfSlice(v)
	return this
}

func (this *floatKindBuilder) NoneOf(v []float64) *floatKindBuilder {
	this.kind.NoneOf = optioner.OfSlice(v)
	return this
}
