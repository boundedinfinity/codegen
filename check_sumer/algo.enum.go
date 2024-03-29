//************************************************************************************
//*                                                                                  *
//* ===== DO NOT EDIT =====                                                          *
//* Any change will be overwritten                                                   *
//* Generated by github.com/boundedinfinity/enumer                                   *
//*                                                                                  *
//************************************************************************************

package check_sumer

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/boundedinfinity/go-commoner/slicer" // v1.0.15
)

var (
	All = []CheckSumAlgo{
		Md5,
		Sha1,
		Sha256,
	}
)

func (t CheckSumAlgo) String() string {
	return string(t)
}

func Parse(v string) (CheckSumAlgo, error) {
	f, ok := slicer.FindFn(All, func(x CheckSumAlgo) bool {
		return CheckSumAlgo(v) == x
	})

	if !ok {
		return f, ErrorV(v)
	}

	return f, nil
}

func Is(s string) bool {
	return slicer.ContainsFn(All, func(v CheckSumAlgo) bool {
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

func (t CheckSumAlgo) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(t))
}

func (t *CheckSumAlgo) UnmarshalJSON(data []byte) error {
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

func (t CheckSumAlgo) MarshalYAML() (interface{}, error) {
	return string(t), nil
}

func (t *CheckSumAlgo) UnmarshalYAML(unmarshal func(interface{}) error) error {
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
