package kind

import (
	"encoding/json"

	"gopkg.in/yaml.v3"
)

type JsonMarshaller struct {
}

func (this JsonMarshaller) Marshal(kind Kind) ([]byte, error) {
	var data []byte
	var err error

	switch kind.(type) {
	case stringEntity:

	}

	return data, err
}

type Marshalable interface {
	ToMap() (map[string]any, error)
	ToJson() ([]byte, error)
	ToJsonIndent() ([]byte, error)
	ToYaml() ([]byte, error)
}

func marshal(e Marshalable, marshal func(any) ([]byte, error)) ([]byte, error) {
	data, err := e.ToMap()

	if err != nil {
		return nil, err
	}

	return marshal(data)
}

func ToJson(e Marshalable) ([]byte, error) {
	return marshal(e, json.Marshal)
}

func ToJsonIndent(e Marshalable) ([]byte, error) {
	return marshal(e, func(data any) ([]byte, error) {
		return json.MarshalIndent(data, "", "    ")
	})
}

func ToYaml(e Marshalable) ([]byte, error) {
	return marshal(e, yaml.Marshal)
}

func ToJsonSchema(e Kind) ([]byte, error) {
	data, err := e.AsJsonSchema()

	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}

func ToJsonSchemaIndent(e Kind) ([]byte, error) {
	data, err := e.AsJsonSchema()

	if err != nil {
		return nil, err
	}

	return json.MarshalIndent(data, "", "    ")
}

// ErrEntityInvalidType = errorer.Errorf(
// 		"invalid entity type, must be one of %v",
// 		strings.Join(asString(validTypes...), ", "),
// 	)

// if this.entityType != noEntityType {
// 		data["entity-type"] = type2StringMap[this.entityType]
// 	}
