package template_type

import (
	"fmt"
	"strings"

	"github.com/boundedinfinity/go-commoner/slicer"
)

func FromUrl(uri string) (TemplateType, error) {
	typ, ok := slicer.FindFn(All, func(x TemplateType) bool {
		term := fmt.Sprintf(".%v.", x)
		return strings.Contains(uri, term)
	})

	if !ok {
		return typ, fmt.Errorf("invalid template type %v", uri)
	}

	return typ, nil
}
