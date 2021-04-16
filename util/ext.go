package util

import (
	"boundedinfinity/codegen/model"
	"path/filepath"
	"strings"

	"github.com/boundedinfinity/optional"
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

	for _, x := range model.TemplateExts {
		if ext == string(x) {
			o = strings.TrimSuffix(o, string(x))
		}
	}

	return o
}

func GetTemplateExt(s string) model.TemplateExt {
	o := s
	ext := filepath.Ext(o)

	for _, rext := range model.TemplateExts {
		if ext == string(rext) {
			return rext
		}
	}

	return model.TemplateExt_Unkown
}
