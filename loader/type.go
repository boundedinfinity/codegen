package loader

import (
	"encoding/json"
	"fmt"
)

type ExampleExtractor func(example interface{}) (interface{}, error)

func json2Interface(example, v interface{}, format string) error {
	if example == nil {
		return nil
	}

	str := fmt.Sprintf(format, example)
	bs := []byte(str)

	if err := json.Unmarshal(bs, v); err != nil {
		return err
	}

	return nil
}

func json2Str(example interface{}) (interface{}, error) {
	var v string
	return v, json2Interface(example, &v, `"%v"`)
}

func json2Boolean(example interface{}) (interface{}, error) {
	var v bool
	return v, json2Interface(example, &v, `"%v"`)
}

func json2Int64(example interface{}) (interface{}, error) {
	var v int64
	return v, json2Interface(example, &v, "%v")
}

func json2Float64(example interface{}) (interface{}, error) {
	var v float64
	return v, json2Interface(example, &v, "%v")
}
