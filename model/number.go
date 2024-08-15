package model

import (
	"github.com/boundedinfinity/go-commoner/errorer"
	"github.com/boundedinfinity/go-commoner/functional/optioner"
)

///////////////////////////////////////////////////////////////////
// number
///////////////////////////////////////////////////////////////////

type Number[T ~int | ~float64] struct {
	CodeGenCommon
	Min        optioner.Option[T]          `json:"min,omitempty"`
	Max        optioner.Option[T]          `json:"max,omitempty"`
	MultipleOf optioner.Option[T]          `json:"multiple-of,omitempty"`
	Ranges     optioner.Option[[]Range[T]] `json:"ranges,omitempty"`
	NoneOf     optioner.Option[[]T]        `json:"one-of,omitempty"`
	OneOf      optioner.Option[[]T]        `json:"some-of,omitempty"`
	Positive   optioner.Option[bool]       `json:"positive,omitempty"`
	Negative   optioner.Option[bool]       `json:"negative,omitempty"`
}

//----------------------------------------------------------------
// Validation
//----------------------------------------------------------------

var (
	ErrNumberMultipleOfBelow1                          = errorer.New("multiple of below 1")
	ErrNumberRange                                     = errorer.New("number range")
	ErrNumberRangePositiveAndNegativeMutuallyExclusive = errorer.New("positive and negative are multually exclusive")
)

func (this Number[T]) Validate() error {
	if err := this.CodeGenCommon.Validate(); err != nil {
		return err
	}

	if this.MultipleOf.Defined() && this.MultipleOf.Get() < 1 {
		return ErrNumberMultipleOfBelow1.WithValue(this.MultipleOf.Get())
	}

	if this.Positive.Defined() && this.Negative.Defined() {
		return ErrNumberRangePositiveAndNegativeMutuallyExclusive
	}

	for _, rng := range this.Ranges.Get() {
		if err := rng.Validate(); err != nil {
			return err
		}
	}

	return nil
}

func (this Number[T]) HasValidation() bool {
	return this.CodeGenCommon.HasValidation() || this.MultipleOf.Defined() || this.Ranges.Defined()
}

///////////////////////////////////////////////////////////////////
// Integer
///////////////////////////////////////////////////////////////////

func NewInteger() *CodeGenInteger {
	return &CodeGenInteger{}
}

type CodeGenInteger struct {
	Number[int]
}

var _ CodeGenSchema = &CodeGenInteger{}

func (_ CodeGenInteger) Schema() string {
	return "integer"
}

func (this *CodeGenInteger) MarshalJSON() ([]byte, error) {
	dto := struct {
		TypeId         string `json:"type"`
		CodeGenInteger `json:",inline"`
	}{
		TypeId:         this.Schema(),
		CodeGenInteger: *this,
	}

	return marshalCodeGenType(dto)
}

///////////////////////////////////////////////////////////////////
// Float
///////////////////////////////////////////////////////////////////

func NewFloat() *CodeGenFloat {
	return &CodeGenFloat{}
}

type CodeGenFloat struct {
	Number[float64]
	Precision optioner.Option[int]     `json:"precision,omitempty"`
	Tolerance optioner.Option[float64] `json:"tolerance,omitempty"`
}

var _ CodeGenSchema = &CodeGenFloat{}

func (_ CodeGenFloat) Schema() string {
	return "float"
}

func (this *CodeGenFloat) MarshalJSON() ([]byte, error) {
	dto := struct {
		TypeId       string `json:"type"`
		CodeGenFloat `json:",inline"`
	}{
		TypeId:       this.Schema(),
		CodeGenFloat: *this,
	}

	return marshalCodeGenType(dto)
}
