package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func RenderFile(input, output string, v interface{}) error {
	if input == "" {
		return fmt.Errorf("render error: no template path")
	}

	if output == "" {
		return fmt.Errorf("render error: no output path")
	}

	if v == nil {
		return fmt.Errorf("render error: no type")
	}

	atp, err := filepath.Abs(input)

	if err != nil {
		return fmt.Errorf("render err: %w", err)
	}

	bs, err := ioutil.ReadFile(atp)

	if err != nil {
		return fmt.Errorf("render err: %w", err)
	}

	ext, err := FileExt(input)

	if err != nil {
		return fmt.Errorf("render err: %w", err)
	}

	aop, err := filepath.Abs(output)

	if err != nil {
		return fmt.Errorf("render err: %w", err)
	}

	if err := DirEnsure(aop); err != nil {
		return fmt.Errorf("render err: %w", err)
	}

	aopCtx := fmt.Sprintf("%v.json", aop)

	if ctxBs, err := json.MarshalIndent(v, "", "    "); err != nil {
		return err
	} else {
		if err := ioutil.WriteFile(aopCtx, ctxBs, 0755); err != nil {
			return fmt.Errorf("render error: %w", err)
		}
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

	if err := ioutil.WriteFile(aop, []byte(o), 0755); err != nil {
		return fmt.Errorf("render error: %w", err)
	}

	return nil
}
