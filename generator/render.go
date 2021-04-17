package generator

import (
	"boundedinfinity/codegen/model"
	"boundedinfinity/codegen/util"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func (t *Generator) renderFile(input, output string, v interface{}) error {
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

	ext := filepath.Ext(input)
	aop, err := filepath.Abs(output)

	if err != nil {
		return fmt.Errorf("render err: %w", err)
	}

	if err := util.DirEnsure(aop); err != nil {
		return fmt.Errorf("render err: %w", err)
	}

	if t.spec.Info.DumpContext {
		aopCtx := fmt.Sprintf("%v.json", aop)

		if err := util.MarshalIndentToFile(aopCtx, v, "", "    "); err != nil {
			return err
		}
	}

	var o string

	switch ext {
	case string(model.TemplateExt_Handlebars):
		if o, err = t.renderHandlebars(string(bs), v); err != nil {
			return fmt.Errorf("render err: %w", err)
		}
	case string(model.TemplateExt_GoTmpl):
		if o, err = t.renderGoTemplate(string(bs), v); err != nil {
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
