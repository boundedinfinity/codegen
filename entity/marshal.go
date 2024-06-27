package entity

import (
	"encoding/json"

	"gopkg.in/yaml.v3"
)

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

func ToJsonSchema(e Marshalable) ([]byte, error) {
	return marshal(e, json.Marshal)
}

func ToJsonSchemaIndent(e Marshalable) ([]byte, error) {
	return marshal(e, func(data any) ([]byte, error) {
		return json.MarshalIndent(data, "", "    ")
	})
}
