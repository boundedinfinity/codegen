package generator_test

import (
	"boundedinfinity/codegen/generator"
	"boundedinfinity/codegen/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Generate(t *testing.T) {
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
			name:            "string 01",
			lang:            "go",
			caserConversion: "kebab-to-pascal",
			input:           model.BuildString().Name("test-output/codegen/schema/util/string-01").Build(),
			expected:        map[string]string{},
		},
		{
			name:            "string 02",
			lang:            "go",
			caserConversion: "kebab-to-pascal",
			input: model.BuildString().
				Name("test-output/codegen/schema/util/string-02").
				Max(50).Min(1).Regex(".*").Build(),
			expected: map[string]string{},
		},
		{
			name:            "integer 01",
			lang:            "go",
			caserConversion: "kebab-to-pascal",
			input:           model.BuildInteger().Name("test-output/codegen/schema/util/integer-01").Build(),
			expected:        map[string]string{},
		},
		{
			name:            "integer 02",
			lang:            "go",
			caserConversion: "kebab-to-pascal",
			input: model.BuildInteger().
				QName("test-output/codegen/schema/util/integer-02").
				MultipleOf(5).Build(),
			expected: map[string]string{},
		},
		{
			name:            "integer 03",
			lang:            "go",
			caserConversion: "kebab-to-pascal",
			input: model.BuildInteger().
				QName("test-output/codegen/schema/util/integer-03").
				Ranges(model.BuildRange[int]().Max(10).Min(1).Build()).Build(),
			expected: map[string]string{},
		},
		{
			name:            "integer 04",
			lang:            "go",
			caserConversion: "kebab-to-pascal",
			input: model.BuildInteger().
				QName("test-output/codegen/schema/util/integer-04").
				MultipleOf(5).
				Ranges(model.BuildRange[int]().Max(10).Min(1).Build()).
				Build(),
			expected: map[string]string{},
		},
		{
			name:            "integer 05",
			lang:            "go",
			caserConversion: "kebab-to-pascal",
			input: model.BuildInteger().
				QName("test-output/codegen/schema/util/integer-05").
				MultipleOf(5).
				Ranges(model.BuildRange[int]().Max(10).Min(1).Build()).
				Ranges(model.BuildRange[int]().ExclusiveMax(20).ExclusiveMin(15).Build()).
				Build(),
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
			assert.Equal(tt, tc.expected, actual)
		})
	}
}
