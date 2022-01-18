package system

import (
	"boundedinfinity/codegen/util"
	"text/template"
)

var (
	functions = template.FuncMap{
		"uc_first": util.UcFirst,
		"uc":       util.Uc,
		"lc_first": util.LcFirst,
		"lc":       util.Lc,
	}
)
