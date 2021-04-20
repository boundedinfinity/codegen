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

func (t *Loader) modelJsonStructure(input model.BiInput_Model, output *model.BiOutput_Model) error {
	// m := make(map[string]interface{})

	// if input.Properties != nil {
	// 	for _, inputProperty := range input.Properties {
	// 		outputProperty := model.BiOutput_Property{}

	// 		if err := t.propertyJsonStructure(inputProperty, &outputProperty); err != nil {
	// 			return err
	// 		}

	// 		for pk, pv := range outputProperty.JsonStruture {
	// 			m[pk] = pv
	// 		}
	// 	}
	// }

	// output.JsonStruture = m
	return nil
}

func json2Str(example interface{}) (string, error) {
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

func json2Int64(example interface{}) (int64, error) {
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
