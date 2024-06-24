package generator_test

import (
	"boundedinfinity/codegen/generator"
	"boundedinfinity/codegen/model"
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/pather"
	"github.com/stretchr/testify/assert"
)

func Test_Generate(t *testing.T) {
	qname := "test-output/codegen/schema/util/name-50"

	tcs := []struct {
		name            string
		lang            string
		input           model.CodeGenType
		caserConversion string
		expected        map[string]string
		err             error
		newErr          error
	}{
		{
			name:            "case 1",
			lang:            "go",
			caserConversion: "kebab-to-pascal",
			input: model.NewString().
				WithQName(qname).
				WithName(pather.Paths.Base(qname)).
				WithPackage(pather.Paths.Dir(qname)).
				WithMax(50).WithMin(1).WithRegex(".*"),
			expected: map[string]string{},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			gen, err := generator.New(tc.lang, tc.caserConversion)
			assert.ErrorIs(tt, err, tc.newErr)

			// actual, err := gen.GenerateType(tc.input)
			actual, err := gen.WriteType(tc.input)
			assert.ErrorIs(tt, err, tc.err)
			assert.ElementsMatch(tt, tc.expected, actual)
		})
	}
}
