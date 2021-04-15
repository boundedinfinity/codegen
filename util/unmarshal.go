package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

func MarshalIndentToFile(p string, v interface{}, prefix, indent string) error {
	wrapErr := func(err error) error {
		return fmt.Errorf("marshal error: %w", err)
	}

	if p == "" {
		return wrapErr(errors.New("path empty"))
	}

	if v == nil {
		return wrapErr(errors.New("no  type"))
	}

	ext := filepath.Ext(p)

	var bs []byte

	switch ext {
	case ".json":
		if x, err := json.MarshalIndent(v, prefix, indent); err != nil {
			return wrapErr(err)
		} else {
			bs = x
		}
	case ".yaml":
		if x, err := yaml.Marshal(v); err != nil {
			return wrapErr(err)
		} else {
			bs = x
		}
	case ".yml":
		if x, err := yaml.Marshal(v); err != nil {
			return wrapErr(err)
		} else {
			bs = x
		}
	default:
		return wrapErr(fmt.Errorf("unmarshal error: unsupported extention %v", ext))
	}

	if err := DirEnsure(p); err != nil {
		return wrapErr(err)
	}

	if err := ioutil.WriteFile(p, bs, 0755); err != nil {
		return wrapErr(err)
	}

	return nil
}

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

	ext := filepath.Ext(p)

	switch ext {
	case ".json":
		if err := json.Unmarshal(bs, v); err != nil {
			return wrapErr(err)
		}
	case ".yaml":
		if err := yaml.Unmarshal(bs, v); err != nil {
			return wrapErr(err)
		}
	case ".yml":
		if err := yaml.Unmarshal(bs, v); err != nil {
			return wrapErr(err)
		}
	default:
		return wrapErr(fmt.Errorf("unmarshal error: unsupported extention %v", ext))
	}

	return nil
}
