package util

import (
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
