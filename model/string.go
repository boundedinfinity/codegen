package model

import (
	"regexp"

	"github.com/boundedinfinity/go-commoner/errorer"
	"github.com/boundedinfinity/go-commoner/functional/optioner"
)

///////////////////////////////////////////////////////////////////
// Type
//////////////////////////////////////////////////////////////////

type CodeGenString struct {
	codeGenCommon
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

var _ CodeGenType = &CodeGenString{}

func (t CodeGenString) GetType() string {
	return "string"
}

//----------------------------------------------------------------
// Validation
//----------------------------------------------------------------

func (t CodeGenString) HasValidation() bool {
	return t.codeGenCommon.HasValidation() || t.Min.Defined() || t.Max.Defined() || t.Regex.Defined() ||
		t.Includes.Defined() || t.Excludes.Defined() || t.OneOf.Defined() || t.NoneOf.Defined()
}

var (
	ErrStringMinBelow1    = errorer.New("min below 1")
	ErrStringMaxBelowMin  = errorer.New("max below min")
	ErrStringInvalidRegex = errorer.New("invalid regex")
)

func (t CodeGenString) Validate() error {
	if err := t.codeGenCommon.Validate(); err != nil {
		return err
	}

	if t.Min.Defined() && t.Min.Get() < 1 {
		return ErrStringMinBelow1.WithValue(t.Min.Get())
	}

	if t.Min.Defined() && t.Max.Defined() {
		if t.Max.Get() < t.Min.Get() {
			return ErrStringMaxBelowMin.FormatFn("max %v, min %v")(t.Max.Get(), t.Min.Get())
		}
	}

	if t.Regex.Defined() {
		_, err := regexp.Compile(t.Regex.Get())
		if err != nil {
			return ErrStringInvalidRegex.WithValue(err.Error())
		}
	}

	if t.Regex.Defined() {
		if _, err := regexp.Compile(t.Regex.Get()); err != nil {
			return err
		}
	}

	return nil
}

// ----------------------------------------------------------------
// Marshal
// ----------------------------------------------------------------

func (t *CodeGenString) MarshalJSON() ([]byte, error) {
	dto := struct {
		TypeId        string `json:"type"`
		CodeGenString `json:",inline"`
	}{
		TypeId:        t.GetType(),
		CodeGenString: *t,
	}

	return marshalCodeGenType(dto)
}

//----------------------------------------------------------------
// Builder
//----------------------------------------------------------------

func BuildString() StringBuilder {
	return &codeGenStringBuilder{}
}

type codeGenStringBuilder struct {
	obj CodeGenString
}

var _ StringBuilder = &codeGenStringBuilder{}

// Build implements StringBuilder.
func (t *codeGenStringBuilder) Build() *CodeGenString {
	return &t.obj
}

// Abnf implements StringBuilder.
func (t *codeGenStringBuilder) Abnf(v string) StringBuilder {
	return setO(t, &t.obj.Abnf, v)
}

// Description implements StringBuilder.
func (t *codeGenStringBuilder) Description(v string) StringBuilder {
	return setO(t, &t.obj.Description, v)
}

// Excludes implements StringBuilder.
func (t *codeGenStringBuilder) Excludes(v ...string) StringBuilder {
	return setO(t, &t.obj.Excludes, v)
}

// Includes implements StringBuilder.
func (t *codeGenStringBuilder) Includes(v ...string) StringBuilder {
	return setO(t, &t.obj.Includes, v)
}

// Max implements StringBuilder.
func (t *codeGenStringBuilder) Max(v int) StringBuilder {
	return setO(t, &t.obj.Max, v)
}

// Min implements StringBuilder.
func (t *codeGenStringBuilder) Min(v int) StringBuilder {
	return setO(t, &t.obj.Min, v)
}

// Name implements StringBuilder.
func (t *codeGenStringBuilder) Name(v string) StringBuilder {
	return setO(t, &t.obj.Name, v)
}

// NoneOf implements StringBuilder.
func (t *codeGenStringBuilder) NoneOf(v ...string) StringBuilder {
	return setO(t, &t.obj.NoneOf, v)
}

// OneOf implements StringBuilder.
func (t *codeGenStringBuilder) OneOf(v ...string) StringBuilder {
	return setO(t, &t.obj.OneOf, v)
}

// Package implements StringBuilder.
func (t *codeGenStringBuilder) Package(v string) StringBuilder {
	return setO(t, &t.obj.Package, v)
}

// QName implements StringBuilder.
func (t *codeGenStringBuilder) QName(v string) StringBuilder {
	panic("unimplemented")
}

// Ref implements StringBuilder.
func (t *codeGenStringBuilder) Ref() RefBuilder {
	panic("unimplemented")
}

// Regex implements StringBuilder.
func (t *codeGenStringBuilder) Regex(v string) StringBuilder {
	return setO(t, &t.obj.Regex, v)
}

// Required implements StringBuilder.
func (t *codeGenStringBuilder) Required(v bool) StringBuilder {
	return setO(t, &t.obj.Required, v)
}
