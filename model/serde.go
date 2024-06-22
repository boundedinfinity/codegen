package model

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/boundedinfinity/go-commoner/errorer"
)

var (
	ErrInvalidMarshalCodeGenObject   = errorer.New("invalid code gen object")
	ErrCodeGenTypeTypeIdMissing      = errorer.New("codegen-id missing")
	ErrCodeGenTypeTypeIdNotSupported = errorer.New("codegen-id not supported")
	ErrCodeGenTypeUnmarshal          = errorer.New("unmarshal error")

	errCodeGenTypeUnmarshalFn = func(err error) error {
		return errors.Join(err, ErrCodeGenTypeUnmarshal)
	}
)

type descriminator struct {
	CodeGenId string `json:"codegen-id"`
}

func UnmarshalCodeGenType(data []byte) (CodeGenType, error) {
	descrim := descriminator{}
	var v CodeGenType

	if err := json.Unmarshal(data, &descrim); err != nil {
		return nil, err
	}

	if descrim.CodeGenId == "" {
		return nil, ErrCodeGenTypeTypeIdMissing
	}

	var err error

	switch descrim.CodeGenId {
	case CodeGenString{}.CodeGenId():
		var obj CodeGenString

		if err = json.Unmarshal(data, &obj); err != nil {
			err = errCodeGenTypeUnmarshalFn(err)
		} else {
			v = &obj
		}
	case CodeGenInteger{}.CodeGenId():
		var obj CodeGenInteger

		if err = json.Unmarshal(data, &obj); err != nil {
			err = errCodeGenTypeUnmarshalFn(err)
		} else {
			v = &obj
		}
	case CodeGenFloat{}.CodeGenId():
		var obj CodeGenFloat

		if err = json.Unmarshal(data, &obj); err != nil {
			err = errCodeGenTypeUnmarshalFn(err)
		} else {
			v = &obj
		}
	case CodeGenBoolean{}.CodeGenId():
		var obj CodeGenBoolean

		if err = json.Unmarshal(data, &obj); err != nil {
			err = errCodeGenTypeUnmarshalFn(err)
		} else {
			v = &obj
		}
	case CodeGenArray{}.CodeGenId():
		var obj CodeGenArray

		if err = json.Unmarshal(data, &obj); err != nil {
			err = errCodeGenTypeUnmarshalFn(err)
		} else {
			v = &obj
		}
	case CodeGenObject{}.CodeGenId():
		var obj CodeGenObject

		if err = json.Unmarshal(data, &obj); err != nil {
			err = errCodeGenTypeUnmarshalFn(err)
		} else {
			v = &obj
		}
	case CodeGenRef{}.CodeGenId():
		var obj CodeGenRef

		if err = json.Unmarshal(data, &obj); err != nil {
			err = errCodeGenTypeUnmarshalFn(err)
		} else {
			v = &obj
		}
	default:
		err = fmt.Errorf("%v : %w", descrim.CodeGenId, ErrCodeGenTypeTypeIdNotSupported)
	}

	return v, err
}
