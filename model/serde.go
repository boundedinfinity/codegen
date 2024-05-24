package model

import (
	"encoding/json"
	"errors"
)

var ErrInvalidMarshalCodeGenObject = errors.New("invalid code gen object")

type typedDto struct {
	TypeId string `json:"type-id"`
	Value  any    `json:"value"`
}

type unmarshalWrapper struct {
	TypeId string          `json:"type-id"`
	Value  json.RawMessage `json:"value"`
}

var ErrCodeGenTypeTypeIdMissing = errors.New("type-id missing")
var ErrCodeGenTypeUnmarshal = errors.New("unmarshal error")

func errCodeGenTypeUnmarshalFn(err error) error {
	return errors.Join(ErrCodeGenTypeUnmarshal, err)
}

func UnmarshalCodeGenObject(data []byte) (CodeGenType, error) {
	wrapper := unmarshalWrapper{}
	var v CodeGenType

	if err := json.Unmarshal(data, &wrapper); err != nil {
		return nil, err
	}

	if wrapper.TypeId == "" {
		return nil, ErrCodeGenTypeTypeIdMissing

	}

	switch wrapper.TypeId {
	case "string":
		var obj CodeGenString

		if err := json.Unmarshal(wrapper.Value, &obj); err != nil {
			return nil, errCodeGenTypeUnmarshalFn(err)
		}

		v = &obj
	case "integer":
		var obj CodeGenInteger

		if err := json.Unmarshal(wrapper.Value, &obj); err != nil {
			return nil, errCodeGenTypeUnmarshalFn(err)
		}

		v = &obj
	case "float":
		var obj CodeGenFloat

		if err := json.Unmarshal(wrapper.Value, &obj); err != nil {
			return nil, errCodeGenTypeUnmarshalFn(err)
		}

		v = &obj
	case "boolean":
		var obj CodeGenBoolean

		if err := json.Unmarshal(wrapper.Value, &obj); err != nil {
			return nil, errCodeGenTypeUnmarshalFn(err)
		}

		v = &obj
	case "array":
		var obj CodeGenArray

		if err := json.Unmarshal(wrapper.Value, &obj); err != nil {
			return nil, errCodeGenTypeUnmarshalFn(err)
		}

		v = &obj
	default:
	}

	return v, nil
}
