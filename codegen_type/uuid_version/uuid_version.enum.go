//************************************************************************************
//*                                                                                  *
//* ===== DO NOT EDIT =====                                                          *
//* Any change will be overwritten                                                   *
//* Generated by github.com/boundedinfinity/enumer                                   *
//*                                                                                  *
//************************************************************************************

package uuid_version

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/boundedinfinity/go-commoner/slicer"
)

var (
	All = []UuidVersion{
		V1,
		V2,
		V3,
		V4,
		V5,
	}
)

func (t UuidVersion) String() string {
	return string(t)
}

func Parse(v string) (UuidVersion, error) {
	f, ok := slicer.FindFn(All, func(x UuidVersion) bool {
		return UuidVersion(v) == x
	})

	if !ok {
		return f, ErrorV(v)
	}

	return f, nil
}

func Is(s string) bool {
	return slicer.ContainsFn(All, func(v UuidVersion) bool {
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

func (t UuidVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(t))
}

func (t *UuidVersion) UnmarshalJSON(data []byte) error {
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

func (t UuidVersion) MarshalYAML() (interface{}, error) {
	return string(t), nil
}

func (t *UuidVersion) UnmarshalYAML(unmarshal func(interface{}) error) error {
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
