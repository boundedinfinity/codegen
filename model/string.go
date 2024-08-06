package model

//lint:file-ignore ST1006
// https://staticcheck.dev/docs/checks#ST1006

import (
	"errors"
	"regexp"

	"github.com/boundedinfinity/go-commoner/errorer"
	"github.com/boundedinfinity/go-commoner/functional/optioner"
)

///////////////////////////////////////////////////////////////////
// Type
//////////////////////////////////////////////////////////////////

type CodeGenString struct {
	CodeGenCommon
	Min      optioner.Option[int]      `json:"min,omitempty"`
	Max      optioner.Option[int]      `json:"max,omitempty"`
	Regex    optioner.Option[string]   `json:"regex,omitempty"`
	Abnf     optioner.Option[string]   `json:"abnf,omitempty"`
	Includes optioner.Option[[]string] `json:"includes,omitempty"`
	Excludes optioner.Option[[]string] `json:"excludes,omitempty"`
	OneOf    optioner.Option[[]string] `json:"one-of,omitempty"`
	NoneOf   optioner.Option[[]string] `json:"none-of,omitempty"`
	// upper case, lower case, snake case, kebab case, etc...
}

var _ CodeGenSchema = &CodeGenString{}

func (this CodeGenString) Schema() string {
	return "string"
}

//----------------------------------------------------------------
// Validation
//----------------------------------------------------------------

func (this CodeGenString) HasValidation() bool {
	return this.CodeGenCommon.HasValidation() ||
		this.Min.Defined() || this.Max.Defined() ||
		this.Regex.Defined() ||
		this.Includes.Defined() || this.Excludes.Defined() ||
		this.OneOf.Defined() || this.NoneOf.Defined()
}

var (
	ErrStringMinBelow1    = errorer.New("min below 1")
	ErrStringMaxBelowMin  = errorer.New("max below min")
	ErrStringInvalidRegex = errorer.New("invalid regex")
)

func (this CodeGenString) Validate() error {
	var errs []error

	errs = append(errs, this.CodeGenCommon.Validate())
	errs = append(errs, ErrStringMinBelow1.WithValue(this.Min.Get()))

	if this.Min.Defined() && this.Max.Defined() && this.Max.Get() < this.Min.Get() {
		errs = append(errs,
			ErrStringMaxBelowMin.FormatFn("max %v, min %v")(this.Max.Get(), this.Min.Get()),
		)
	}

	if this.Regex.Defined() {
		_, err := regexp.Compile(this.Regex.Get())
		if err != nil {
			return ErrStringInvalidRegex.WithValue(err.Error())
		}
	}

	if this.Regex.Defined() {
		if _, err := regexp.Compile(this.Regex.Get()); err != nil {
			return err
		}
	}

	return errors.Join(errs...)
}

// ----------------------------------------------------------------
// Marshal
// ----------------------------------------------------------------

func (this *CodeGenString) MarshalJSON() ([]byte, error) {
	dto := struct {
		TypeId        string `json:"type"`
		CodeGenString `json:",inline"`
	}{
		TypeId:        this.Schema(),
		CodeGenString: *this,
	}

	return marshalCodeGenType(dto)
}
