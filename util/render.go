package util

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func RenderFile(tp, op string, v interface{}) error {
	if tp == "" {
		return fmt.Errorf("render error: no template path")
	}

	if op == "" {
		return fmt.Errorf("render error: no output path")
	}

	if v == nil {
		return fmt.Errorf("render error: no type")
	}

	atp, err := filepath.Abs(tp)

	if err != nil {
		return fmt.Errorf("render err: %w", err)
	}

	bs, err := ioutil.ReadFile(atp)

	if err != nil {
		return fmt.Errorf("render err: %w", err)
	}

	ext, err := FileExt(tp)

	if err != nil {
		return fmt.Errorf("render err: %w", err)
	}

	var o string

	switch ext {
	case "handlebars":
		if o, err = renderHandlebars(string(bs), v); err != nil {
			return fmt.Errorf("render err: %w", err)
		}
	case "gotmpl":
		if o, err = renderGoTemplate(string(bs), v); err != nil {
			return fmt.Errorf("render err: %w", err)
		}
	default:
		return fmt.Errorf("render error: unsupported extention %v", ext)
	}

	aop, err := filepath.Abs(op)

	if err != nil {
		return fmt.Errorf("render err: %w", err)
	}

	if err := DirEnsure(aop); err != nil {
		return fmt.Errorf("render err: %w", err)
	}

	if err := ioutil.WriteFile(aop, []byte(o), 0755); err != nil {
		return fmt.Errorf("render error: %w", err)
	}

	return nil
}
