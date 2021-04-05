package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func UnmarshalFromFile(p string, v interface{}) error {
	wrapErr := func(err error) error {
		return fmt.Errorf("unmarshal error: %w", err)
	}

	if p == "" {
		return wrapErr(errors.New("path empty"))
	}

	if v == nil {
		return wrapErr(errors.New("no  type"))
	}

	bs, err := ioutil.ReadFile(p)

	if err != nil {
		return wrapErr(err)
	}

	ext, err := FileExt(p)

	if err != nil {
		return wrapErr(err)
	}

	switch ext {
	case "json":
		if err := json.Unmarshal(bs, v); err != nil {
			return wrapErr(err)
		}
	case "yaml":
		if err := yaml.Unmarshal(bs, v); err != nil {
			return wrapErr(err)
		}
	case "yml":
		if err := yaml.Unmarshal(bs, v); err != nil {
			return wrapErr(err)
		}
	default:
		return wrapErr(fmt.Errorf("unmarshal error: unsupported extention %v", ext))
	}

	return nil
}
