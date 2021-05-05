package loader

import (
	"boundedinfinity/codegen/model"
	"encoding/json"
	"fmt"
)

var (
	jsonTypeMap = map[string]string{
		"string":       "string",
		"number":       "number",
		"integer":      "number",
		"biginteger":   "number",
		"smallinteger": "number",
		"float":        "number",
		"smallfloat":   "number",
		"bigfloat":     "number",
		"boolean":      "boolean",
	}
)

type ExampleExtractor func(example interface{}) (interface{}, error)

func json2Str(example interface{}) (interface{}, error) {
	if example == nil {
		return model.JSON_DEFAULT_STRING, nil
	}

	var v string
	str := fmt.Sprintf(`"%v"`, example)
	bs := []byte(str)

	if err := json.Unmarshal(bs, &v); err != nil {
		return v, err
	}

	return v, nil
}

func json2Boolean(example interface{}) (interface{}, error) {
	if example == nil {
		return false, nil
	}

	var v bool
	str := fmt.Sprintf(`"%v"`, example)
	bs := []byte(str)

	if err := json.Unmarshal(bs, &v); err != nil {
		return v, err
	}

	return v, nil
}

func json2Int64(example interface{}) (interface{}, error) {
	if example == nil {
		return model.JSON_DEFAULT_NUMBER, nil
	}

	var v int64
	str := fmt.Sprintf("%v", example)
	bs := []byte(str)

	if err := json.Unmarshal(bs, &v); err != nil {
		return v, err
	}

	return v, nil
}

func json2Float64(example interface{}) (interface{}, error) {
	if example == nil {
		return model.JSON_DEFAULT_NUMBER, nil
	}

	var v float64
	str := fmt.Sprintf("%v", example)
	bs := []byte(str)

	if err := json.Unmarshal(bs, &v); err != nil {
		return v, err
	}

	return v, nil
}
