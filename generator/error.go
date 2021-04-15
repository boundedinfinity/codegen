package generator

import (
	"fmt"
	"strings"
)

func (t *Generator) generatorErr(err error, ps ...string) error {
	p := strings.Join(ps, ".")
	return fmt.Errorf("%v: %w", p, err)
}

func (t *Generator) generatorSchemaErr(err error, ps ...string) error {
	// tmp := []string{fmt.Sprintf("spec[%v]", t.specPath)}
	tmp := []string{fmt.Sprintf("spec[%v]", "X")}
	tmp = append(tmp, ps...)
	return t.generatorErr(err, tmp...)
}
