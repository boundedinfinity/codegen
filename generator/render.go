package generator

import (
	"boundedinfinity/codegen/model"
	"boundedinfinity/codegen/util"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func (t *Generator) renderFile(tmpl model.OutputTemplate, v interface{}) error {
	fmt.Printf("rendering: %v\n", tmpl.Output)

	if tmpl.Input == "" {
		return fmt.Errorf("render error: no template path")
	}

	if tmpl.Output == "" {
		return fmt.Errorf("render error: no output path")
	}

	if v == nil {
		return fmt.Errorf("render error: no type")
	}

	atp, err := filepath.Abs(tmpl.Input)

	if err != nil {
		return fmt.Errorf("render err: %w", err)
	}

	bs, err := ioutil.ReadFile(atp)

	if err != nil {
		return fmt.Errorf("render err: %w", err)
	}

	ext := filepath.Ext(tmpl.Input)
	ext = strings.TrimPrefix(ext, ".")
	aop, err := filepath.Abs(tmpl.Output)

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
	case string(model.TemplateExt_Gotmpl):
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
