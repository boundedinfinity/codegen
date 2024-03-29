//************************************************************************************
//*                                                                                  *
//* ===== DO NOT EDIT =====                                                          *
//* Any change will be overwritten                                                   *
//* Generated by github.com/boundedinfinity/enumer                                   *
//*                                                                                  *
//************************************************************************************

package codegen_type_id

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/boundedinfinity/go-commoner/slicer" // v1.0.15
)

var (
	All = []CodgenTypeId{
		Ref,
		Array,
		Coordinate,
		CreditCardNumber,
		Date,
		Datetime,
		Duration,
		Email,
		Enum,
		Float,
		Integer,
		Ipv4,
		Ipv6,
		Mac,
		Object,
		Phone,
		SemanticVersion,
		String,
		Time,
		Uuid,
		Url,
		Path,
	}
)

func (t CodgenTypeId) String() string {
	return string(t)
}

func Parse(v string) (CodgenTypeId, error) {
	f, ok := slicer.FindFn(All, func(x CodgenTypeId) bool {
		return CodgenTypeId(v) == x
	})

	if !ok {
		return f, ErrorV(v)
	}

	return f, nil
}

func Is(s string) bool {
	return slicer.ContainsFn(All, func(v CodgenTypeId) bool {
		return string(v) == s
	})
}

var ErrInvalid = errors.New("invalid enumeration type")

func ErrorV(v string) error {
	return fmt.Errorf(
		"%w '%v', must be one of %v",
		ErrInvalid, v, slicer.Join(All, ","),
	)
}

func (t CodgenTypeId) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(t))
}

func (t *CodgenTypeId) UnmarshalJSON(data []byte) error {
	var v string

	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	e, err := Parse(v)

	if err != nil {
		return err
	}

	*t = e

	return nil
}

func (t CodgenTypeId) MarshalYAML() (interface{}, error) {
	return string(t), nil
}

func (t *CodgenTypeId) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var v string

	if err := unmarshal(&v); err != nil {
		return err
	}

	e, err := Parse(v)

	if err != nil {
		return err
	}

	*t = e

	return nil
}
