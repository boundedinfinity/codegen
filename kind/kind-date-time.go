package kind

import (
	"boundedinfinity/codegen/errorer"
	"boundedinfinity/codegen/kind/name"
	"errors"

	"github.com/boundedinfinity/go-commoner/functional/optioner"
	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
	"github.com/boundedinfinity/rfc3339date"
)

// //////////////////////////////////////////////////////////////////////////
// DateTime Kind
// //////////////////////////////////////////////////////////////////////////

var _ Kind = &DateTimeKind{}

type DateTimeKind struct {
	KindCommon
	Before optioner.Option[rfc3339date.Rfc3339DateTime]
	After  optioner.Option[rfc3339date.Rfc3339DateTime]
	OneOf  optioner.Option[[]rfc3339date.Rfc3339DateTime]
	NoneOf optioner.Option[[]rfc3339date.Rfc3339DateTime]
}

func (this DateTimeKind) KindName() name.KindName {
	return name.DateTime
}

var (
	ErrDateTimeKindBeforeLaterThanAfter   = errorer.New("before is later then after")
	errDateTimeKindBeforeLaterThanAfterFn = errorer.ValueFnf(ErrDateTimeKindBeforeLaterThanAfter, "%v before, %v after")
	ErrDateTimeKindOneOfOverlapsNoneOf    = errorer.New("one-of overlaps with none-of")
	errDateTimeKindOneOfOverlapsNoneOfFn  = errorer.ValueFnf(ErrDateTimeKindOneOfOverlapsNoneOf, "overlapping values [%v]")
)

func (this DateTimeKind) Validate(config ValidatorConfig) error {
	var errs []error

	if this.OneOf.Defined() && this.NoneOf.Defined() {
		var overlaps []rfc3339date.Rfc3339DateTime

		for _, oneOf := range this.OneOf.Get() {
			for _, noneOf := range this.NoneOf.Get() {
				if oneOf == noneOf {
					overlaps = append(overlaps, oneOf)
				}
			}
		}

		if len(overlaps) > 0 {
			value := stringer.Join(", ", overlaps...)
			errs = append(errs, errDateTimeKindOneOfOverlapsNoneOfFn(value))
		}
	}

	return errors.Join(errs...)
}

// //////////////////////////////////////////////////////////////////////////
// DateTime Kind Builder
// //////////////////////////////////////////////////////////////////////////

var _ kindBuilder[dateTimeKindBuilder] = &dateTimeKindBuilder{}

type dateTimeKindBuilder struct {
	kind *DateTimeKind
}

func (this *dateTimeKindBuilder) Done() Kind {
	return *this.kind
}

func (this *dateTimeKindBuilder) Name(v string) *dateTimeKindBuilder {
	this.kind.KindCommon.Name = optioner.OfZero(v)
	return this
}

func (this *dateTimeKindBuilder) Type(v string) *dateTimeKindBuilder {
	this.kind.KindCommon.Type = optioner.OfZero(v)
	return this
}

func (this *dateTimeKindBuilder) Before(v rfc3339date.Rfc3339DateTime) *dateTimeKindBuilder {
	this.kind.Before = optioner.Some(v)
	return this
}

func (this *dateTimeKindBuilder) After(v rfc3339date.Rfc3339DateTime) *dateTimeKindBuilder {
	this.kind.After = optioner.Some(v)
	return this
}

func (this *dateTimeKindBuilder) OneOf(v []rfc3339date.Rfc3339DateTime) *dateTimeKindBuilder {
	this.kind.OneOf = optioner.OfSlice(v)
	return this
}

func (this *dateTimeKindBuilder) NoneOf(v []rfc3339date.Rfc3339DateTime) *dateTimeKindBuilder {
	this.kind.NoneOf = optioner.OfSlice(v)
	return this
}

// //////////////////////////////////////////////////////////////////////////
// Date Kind
// //////////////////////////////////////////////////////////////////////////

var _ Kind = &DateKind{}

type DateKind struct {
	KindCommon
	Before optioner.Option[rfc3339date.Rfc3339Date]
	After  optioner.Option[rfc3339date.Rfc3339Date]
	OneOf  optioner.Option[[]rfc3339date.Rfc3339Date]
	NoneOf optioner.Option[[]rfc3339date.Rfc3339Date]
}

func (this DateKind) KindName() name.KindName {
	return name.Date
}

var (
	ErrDateKindBeforeLaterThanAfter   = errorer.New("before is later then after")
	errDateKindBeforeLaterThanAfterFn = errorer.ValueFnf(ErrDateKindBeforeLaterThanAfter, "%v before, %v after")
	ErrDateKindOneOfOverlapsNoneOf    = errorer.New("one-of overlaps with none-of")
	errDateKindOneOfOverlapsNoneOfFn  = errorer.ValueFnf(ErrDateKindOneOfOverlapsNoneOf, "overlapping values [%v]")
)

func (this DateKind) Validate(config ValidatorConfig) error {
	var errs []error

	if this.OneOf.Defined() && this.NoneOf.Defined() {
		var overlaps []rfc3339date.Rfc3339Date

		for _, oneOf := range this.OneOf.Get() {
			for _, noneOf := range this.NoneOf.Get() {
				if oneOf == noneOf {
					overlaps = append(overlaps, oneOf)
				}
			}
		}

		if len(overlaps) > 0 {
			value := stringer.Join(", ", overlaps...)
			errs = append(errs, errDateKindOneOfOverlapsNoneOfFn(value))
		}
	}

	return errors.Join(errs...)
}

// //////////////////////////////////////////////////////////////////////////
// Date Kind Builder
// //////////////////////////////////////////////////////////////////////////

var _ kindBuilder[dateKindBuilder] = &dateKindBuilder{}

type dateKindBuilder struct {
	kind *DateKind
}

func (this *dateKindBuilder) Done() Kind {
	return *this.kind
}

func (this *dateKindBuilder) Name(v string) *dateKindBuilder {
	this.kind.KindCommon.Name = optioner.OfZero(v)
	return this
}

func (this *dateKindBuilder) Type(v string) *dateKindBuilder {
	this.kind.KindCommon.Type = optioner.OfZero(v)
	return this
}

func (this *dateKindBuilder) Before(v rfc3339date.Rfc3339Date) *dateKindBuilder {
	this.kind.Before = optioner.Some(v)
	return this
}

func (this *dateKindBuilder) After(v rfc3339date.Rfc3339Date) *dateKindBuilder {
	this.kind.After = optioner.Some(v)
	return this
}

func (this *dateKindBuilder) OneOf(v []rfc3339date.Rfc3339Date) *dateKindBuilder {
	this.kind.OneOf = optioner.OfSlice(v)
	return this
}

func (this *dateKindBuilder) NoneOf(v []rfc3339date.Rfc3339Date) *dateKindBuilder {
	this.kind.NoneOf = optioner.OfSlice(v)
	return this
}

// //////////////////////////////////////////////////////////////////////////
// Time Kind
// //////////////////////////////////////////////////////////////////////////

var _ Kind = &TimeKind{}

type TimeKind struct {
	KindCommon
	Before optioner.Option[rfc3339date.Rfc3339Time]
	After  optioner.Option[rfc3339date.Rfc3339Time]
	OneOf  optioner.Option[[]rfc3339date.Rfc3339Time]
	NoneOf optioner.Option[[]rfc3339date.Rfc3339Time]
}

func (this TimeKind) KindName() name.KindName {
	return name.Time
}

var (
	ErrTimeKindBeforeLaterThanAfter   = errorer.New("before is later then after")
	errTimeKindBeforeLaterThanAfterFn = errorer.ValueFnf(ErrTimeKindBeforeLaterThanAfter, "%v before, %v after")
	ErrTimeKindOneOfOverlapsNoneOf    = errorer.New("one-of overlaps with none-of")
	errTimeKindOneOfOverlapsNoneOfFn  = errorer.ValueFnf(ErrTimeKindOneOfOverlapsNoneOf, "overlapping values [%v]")
)

func (this TimeKind) Validate(config ValidatorConfig) error {
	var errs []error

	if this.OneOf.Defined() && this.NoneOf.Defined() {
		var overlaps []rfc3339date.Rfc3339Time

		for _, oneOf := range this.OneOf.Get() {
			for _, noneOf := range this.NoneOf.Get() {
				if oneOf == noneOf {
					overlaps = append(overlaps, oneOf)
				}
			}
		}

		if len(overlaps) > 0 {
			value := stringer.Join(", ", overlaps...)
			errs = append(errs, errTimeKindOneOfOverlapsNoneOfFn(value))
		}
	}

	return errors.Join(errs...)
}

// //////////////////////////////////////////////////////////////////////////
// Time Kind Builder
// //////////////////////////////////////////////////////////////////////////

var _ kindBuilder[timeKindBuilder] = &timeKindBuilder{}

type timeKindBuilder struct {
	kind *TimeKind
}

func (this *timeKindBuilder) Done() Kind {
	return *this.kind
}

func (this *timeKindBuilder) Name(v string) *timeKindBuilder {
	this.kind.KindCommon.Name = optioner.OfZero(v)
	return this
}

func (this *timeKindBuilder) Type(v string) *timeKindBuilder {
	this.kind.KindCommon.Type = optioner.OfZero(v)
	return this
}

func (this *timeKindBuilder) Before(v rfc3339date.Rfc3339Time) *timeKindBuilder {
	this.kind.Before = optioner.Some(v)
	return this
}

func (this *timeKindBuilder) After(v rfc3339date.Rfc3339Time) *timeKindBuilder {
	this.kind.After = optioner.Some(v)
	return this
}

func (this *timeKindBuilder) OneOf(v []rfc3339date.Rfc3339Time) *timeKindBuilder {
	this.kind.OneOf = optioner.OfSlice(v)
	return this
}

func (this *timeKindBuilder) NoneOf(v []rfc3339date.Rfc3339Time) *timeKindBuilder {
	this.kind.NoneOf = optioner.OfSlice(v)
	return this
}
