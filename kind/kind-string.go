package kind

import (
	"boundedinfinity/codegen/errorer"
	"boundedinfinity/codegen/kind/name"
	"errors"
	"regexp"

	"github.com/boundedinfinity/go-commoner/functional/optioner"
)

// //////////////////////////////////////////////////////////////////////////
// String Kind
// //////////////////////////////////////////////////////////////////////////

var _ Kind = &StringKind{}

type StringKind struct {
	KindCommon
	Min        optioner.Option[int]
	Max        optioner.Option[int]
	Regex      optioner.Option[string]
	StartsWith optioner.Option[string]
	EndsWith   optioner.Option[string]
	Contains   optioner.Option[string]
}

func (this StringKind) KindName() name.KindName {
	return name.String
}

var (
	ErrStringKindMinGreaterThanMax   = errorer.New("min is greater than max")
	errStringKindMinGreaterThanMaxFn = errorer.ValueFnf(ErrStringKindMinGreaterThanMax, "%v max, %v min")
	ErrStringKindMinNegative         = errorer.New("min is negative number")
	errStringKindMinNegativeFn       = errorer.ValueFn(ErrStringKindMinNegative)
	ErrStringKindMaxNegative         = errorer.New("max is negative number")
	errStringKindMaxNegativeFn       = errorer.ValueFn(ErrStringKindMaxNegative)
	ErrStringKindInvalidRegex        = errorer.New("invalid regex")
	errStringKindInvalidRegexFn      = errorer.ValueFnf(ErrStringKindInvalidRegex, "%v : %w")
)

func (this StringKind) Validate(config ValidatorConfig) error {
	var errs []error

	if this.Min.Defined() && this.Min.Get() < 0 {
		errs = append(errs, errStringKindMinNegativeFn(this.Min.Get()))
	}

	if this.Max.Defined() && this.Max.Get() < 0 {
		errs = append(errs, errStringKindMaxNegativeFn(this.Max.Get()))
	}

	if this.Max.Defined() && this.Min.Defined() && this.Min.Get() > this.Max.Get() {
		errs = append(errs, errStringKindMinGreaterThanMaxFn(this.Max.Get(), this.Min.Get()))
	}

	if this.Regex.Defined() {
		_, err := regexp.Compile(this.Regex.Get())

		if err != nil {
			errs = append(errs, errStringKindInvalidRegexFn(this.Regex.Get(), err))
		}
	}

	return errors.Join(errs...)
}

// //////////////////////////////////////////////////////////////////////////
// String Kind Builder
// //////////////////////////////////////////////////////////////////////////

var _ kindBuilder[stringKindBuilder] = &stringKindBuilder{}

type stringKindBuilder struct {
	kind *StringKind
}

func (this *stringKindBuilder) Done() Kind {
	return *this.kind
}

func (this *stringKindBuilder) Name(v string) *stringKindBuilder {
	this.kind.KindCommon.Name = optioner.OfZero(v)
	return this
}

func (this *stringKindBuilder) Type(v string) *stringKindBuilder {
	this.kind.KindCommon.Type = optioner.OfZero(v)
	return this
}

func (this *stringKindBuilder) Min(v int) *stringKindBuilder {
	this.kind.Min = optioner.Some(v)
	return this
}

func (this *stringKindBuilder) Max(v int) *stringKindBuilder {
	this.kind.Max = optioner.Some(v)
	return this
}

func (this *stringKindBuilder) Regex(v string) *stringKindBuilder {
	this.kind.Regex = optioner.OfZero(v)
	return this
}

func (this *stringKindBuilder) StartsWith(v string) *stringKindBuilder {
	this.kind.StartsWith = optioner.OfZero(v)
	return this
}

func (this *stringKindBuilder) EndsWith(v string) *stringKindBuilder {
	this.kind.EndsWith = optioner.OfZero(v)
	return this
}

func (this *stringKindBuilder) Contains(v string) *stringKindBuilder {
	this.kind.Contains = optioner.OfZero(v)
	return this
}
