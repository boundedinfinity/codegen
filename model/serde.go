package model

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/boundedinfinity/go-commoner/errorer"
)

var (
	ErrInvalidMarshalCodeGenObject   = errorer.New("invalid code gen object")
	ErrCodeGenTypeTypeIdMissing      = errorer.New("type missing")
	ErrCodeGenTypeTypeIdNotSupported = errorer.New("type not supported")
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
	Type string `json:"type"`
}

func UnmarshalCodeGenType(data []byte) (CodeGenSchema, error) {
	descrim := descriminator{}
	var v CodeGenSchema

	if err := json.Unmarshal(data, &descrim); err != nil {
		return nil, err
	}

	if descrim.Type == "" {
		return nil, ErrCodeGenTypeTypeIdMissing
	}

	var err error

	switch descrim.Type {
	case CodeGenString{}.Schema():
		var obj CodeGenString

		if err = json.Unmarshal(data, &obj); err != nil {
			err = errCodeGenTypeUnmarshalFn(err)
		} else {
			v = &obj
		}
	case CodeGenInteger{}.Schema():
		var obj CodeGenInteger

		if err = json.Unmarshal(data, &obj); err != nil {
			err = errCodeGenTypeUnmarshalFn(err)
		} else {
			v = &obj
		}
	case CodeGenFloat{}.Schema():
		var obj CodeGenFloat

		if err = json.Unmarshal(data, &obj); err != nil {
			err = errCodeGenTypeUnmarshalFn(err)
		} else {
			v = &obj
		}
	case CodeGenBoolean{}.Schema():
		var obj CodeGenBoolean

		if err = json.Unmarshal(data, &obj); err != nil {
			err = errCodeGenTypeUnmarshalFn(err)
		} else {
			v = &obj
		}
	case CodeGenArray{}.Schema():
		var obj CodeGenArray

		if err = json.Unmarshal(data, &obj); err != nil {
			err = errCodeGenTypeUnmarshalFn(err)
		} else {
			v = &obj
		}
	case CodeGenObject{}.Schema():
		var obj CodeGenObject

		if err = json.Unmarshal(data, &obj); err != nil {
			err = errCodeGenTypeUnmarshalFn(err)
		} else {
			v = &obj
		}
	case CodeGenRef{}.Schema():
		var obj CodeGenRef

		if err = json.Unmarshal(data, &obj); err != nil {
			err = errCodeGenTypeUnmarshalFn(err)
		} else {
			v = &obj
		}
	default:
		err = fmt.Errorf("%v : %w", descrim.Type, ErrCodeGenTypeTypeIdNotSupported)
	}

	return v, err
}
