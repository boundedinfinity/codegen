package util

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/boundedinfinity/optional"
)

type RenderExt string

const (
	GoTmplExt     RenderExt = ".gotmpl"
	HandlebarsExt RenderExt = ".handlebars"
)

var (
	RenderExts = []RenderExt{
		GoTmplExt,
		HandlebarsExt,
	}
)

func TrimTemplateExtO(s optional.StringOptional) optional.StringOptional {
	if s.IsDefined() {
		return optional.NewStringValue(TrimTemplateExt(s.Get()))
	}

	return s
}

func TrimTemplateExt(s string) string {
	o := s
	ext := filepath.Ext(o)

	for _, x := range RenderExts {
		if ext == string(x) {
			o = strings.TrimSuffix(o, string(x))
		}
	}

	return o
}

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

	if err := MarshalIndentToFile(aopCtx, v, "", "    "); err != nil {
		return err
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
