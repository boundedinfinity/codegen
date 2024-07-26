package model

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/boundedinfinity/go-commoner/errorer"
)

var (
	ErrInvalidMarshalCodeGenObject   = errorer.New("invalid code gen object")
	ErrCodeGenTypeTypeIdMissing      = errorer.New("base-type missing")
	ErrCodeGenTypeTypeIdNotSupported = errorer.New("base-type not supported")
	ErrCodeGenTypeUnmarshal          = errorer.New("unmarshal error")

	errCodeGenTypeUnmarshalFn = func(err error) error {
		return errors.Join(err, ErrCodeGenTypeUnmarshal)
	}
)

func marshalCodeGenType(v any) ([]byte, error) {
	data, err := json.Marshal(v)

	if err != nil {
		return data, err
	}

	var temp map[string]any

	if err := json.Unmarshal(data, &temp); err != nil {
		return data, err
	}

	for k, v := range temp {
		if v == nil {
			delete(temp, k)
		}
	}

	return json.Marshal(temp)
}

type descriminator struct {
	CodeGenId string `json:"base-type"`
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
	case CodeGenString{}.GetType():
		var obj CodeGenString

		if err = json.Unmarshal(data, &obj); err != nil {
			err = errCodeGenTypeUnmarshalFn(err)
		} else {
			v = &obj
		}
	case CodeGenInteger{}.GetType():
		var obj CodeGenInteger

		if err = json.Unmarshal(data, &obj); err != nil {
			err = errCodeGenTypeUnmarshalFn(err)
		} else {
			v = &obj
		}
	case CodeGenFloat{}.GetType():
		var obj CodeGenFloat

		if err = json.Unmarshal(data, &obj); err != nil {
			err = errCodeGenTypeUnmarshalFn(err)
		} else {
			v = &obj
		}
	case CodeGenBoolean{}.GetType():
		var obj CodeGenBoolean

		if err = json.Unmarshal(data, &obj); err != nil {
			err = errCodeGenTypeUnmarshalFn(err)
		} else {
			v = &obj
		}
	case CodeGenArray{}.GetType():
		var obj CodeGenArray

		if err = json.Unmarshal(data, &obj); err != nil {
			err = errCodeGenTypeUnmarshalFn(err)
		} else {
			v = &obj
		}
	case CodeGenObject{}.GetType():
		var obj CodeGenObject

		if err = json.Unmarshal(data, &obj); err != nil {
			err = errCodeGenTypeUnmarshalFn(err)
		} else {
			v = &obj
		}
	case CodeGenRef{}.GetType():
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
