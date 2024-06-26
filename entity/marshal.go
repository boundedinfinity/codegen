package entity

import (
	"encoding/json"

	"gopkg.in/yaml.v3"
)

func (t *entityBase) toMap() (map[string]any, error) {
	data := map[string]any{}

	if t.entityType != noneType {
		data["entity-type"] = type2String[t.entityType]
	}

	if t.name != "" {
		data["name"] = t.name
	}

	if t.description != "" {
		data["description"] = t.description
	}

	if t.required {
		data["required"] = t.required
	}

	if t.defaultValue != nil {
		val, err := t.defaultValue.toMap()

		if err != nil {
			return data, err
		}

		data["default"] = val
	}

	return data, nil
}

func (t *entityBase) ToJson() ([]byte, error) {
	data, err := t.toMap()

	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}

func (t *entityBase) ToJsonIndent() ([]byte, error) {
	data, err := t.toMap()

	if err != nil {
		return nil, err
	}

	return json.MarshalIndent(data, "", "    ")
}

func (t *entityBase) ToYaml() ([]byte, error) {
	data, err := t.toMap()

	if err != nil {
		return nil, err
	}

	return yaml.Marshal(data)
}
