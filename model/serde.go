package model

import (
	"encoding/json"
	"errors"
	"fmt"
)

var (
	ErrInvalidMarshalCodeGenObject   = errors.New("invalid code gen object")
	ErrCodeGenTypeTypeIdMissing      = errors.New("type-id missing")
	ErrCodeGenTypeTypeIdNotSupported = errors.New("type-id not supported")
	ErrCodeGenTypeUnmarshal          = errors.New("unmarshal error")

	errCodeGenTypeUnmarshalFn = func(err error) error {
		return errors.Join(err, ErrCodeGenTypeUnmarshal)
	}
)

type unmarshalDto struct {
	TypeId string          `json:"type-id"`
	Value  json.RawMessage `json:"value"`
}

func UnmarshalCodeGenType(data []byte) (CodeGenType, error) {
	wrapper := unmarshalDto{}
	var v CodeGenType

	if err := json.Unmarshal(data, &wrapper); err != nil {
		return nil, err
	}

	if wrapper.TypeId == "" {
		return nil, ErrCodeGenTypeTypeIdMissing
	}

	var err error

	switch wrapper.TypeId {
	case CodeGenString{}.TypeId():
		var obj CodeGenString

		if err = json.Unmarshal(wrapper.Value, &obj); err != nil {
			err = errCodeGenTypeUnmarshalFn(err)
		} else {
			v = &obj
		}
	case CodeGenInteger{}.TypeId():
		var obj CodeGenInteger

		if err = json.Unmarshal(wrapper.Value, &obj); err != nil {
			err = errCodeGenTypeUnmarshalFn(err)
		} else {
			v = &obj
		}
	case CodeGenFloat{}.TypeId():
		var obj CodeGenFloat

		if err = json.Unmarshal(wrapper.Value, &obj); err != nil {
			err = errCodeGenTypeUnmarshalFn(err)
		} else {
			v = &obj
		}
	case CodeGenBoolean{}.TypeId():
		var obj CodeGenBoolean

		if err = json.Unmarshal(data, &obj); err != nil {
			err = errCodeGenTypeUnmarshalFn(err)
		} else {
			v = &obj
		}
	case CodeGenArray{}.TypeId():
		var obj CodeGenArray

		if err = json.Unmarshal(wrapper.Value, &obj); err != nil {
			err = errCodeGenTypeUnmarshalFn(err)
		} else {
			v = &obj
		}
	case CodeGenObject{}.TypeId():
		var obj CodeGenObject

		if err = json.Unmarshal(wrapper.Value, &obj); err != nil {
			err = errCodeGenTypeUnmarshalFn(err)
		} else {
			v = &obj
		}
	default:
		err = fmt.Errorf("%v : %w", wrapper.TypeId, ErrCodeGenTypeTypeIdNotSupported)
	}

	return v, err
}
