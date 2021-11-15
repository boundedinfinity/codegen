//************************************************************************************
//*                                                                                  *
//* ===== DO NOT EDIT =====                                                          *
//* Any change will be overwritten                                                   *
//* Generated by github.com/boundedinfinity/enumer                                   *
//*                                                                                  *
//************************************************************************************

package template_type

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type TemplateType string
type TemplateTypes []TemplateType

func Slice(es ...TemplateType) TemplateTypes {
	var s TemplateTypes

	for _, e := range es {
		s = append(s, e)
	}

	return s
}

const (
	Model     TemplateType = "model"
	Namespace TemplateType = "namespace"
	Operation TemplateType = "operation"
)

var (
	All = TemplateTypes{
		Model,
		Namespace,
		Operation,
	}
)

func Is(v string) bool {
	return All.Is(v)
}

func Parse(v string) (TemplateType, error) {
	return All.Parse(v)
}

func Strings() []string {
	return All.Strings()
}

func (t TemplateType) String() string {
	return string(t)
}

var ErrTemplateTypeInvalid = errors.New("invalid enumeration type")

func Error(vs TemplateTypes, v string) error {
	return fmt.Errorf(
		"%w '%v', must be one of %v",
		ErrTemplateTypeInvalid, v, strings.Join(vs.Strings(), ","),
	)
}

func (t TemplateTypes) Strings() []string {
	var ss []string

	for _, v := range t {
		ss = append(ss, v.String())
	}

	return ss
}

func (t TemplateTypes) Parse(v string) (TemplateType, error) {
	var o TemplateType
	var f bool
	n := strings.ToLower(v)

	for _, e := range t {
		if strings.ToLower(e.String()) == n {
			o = e
			f = true
			break
		}
	}

	if !f {
		return o, Error(t, v)
	}

	return o, nil
}

func (t TemplateTypes) Is(v string) bool {
	var f bool

	for _, e := range t {
		if string(e) == v {
			f = true
			break
		}
	}

	return f
}

func (t TemplateTypes) Contains(v TemplateType) bool {
	for _, e := range t {
		if e == v {
			return true
		}
	}

	return false
}

func (t TemplateType) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(t))
}

func (t *TemplateType) UnmarshalJSON(data []byte) error {
	var s string

	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	e, err := Parse(s)

	if err != nil {
		return err
	}

	*t = e

	return nil
}
