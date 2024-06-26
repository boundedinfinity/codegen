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
			input: model.NewString().
				WithQName("test-output/codegen/schema/util/string-01"),
			expected: map[string]string{},
		},
		{
			name:            "string 02",
			lang:            "go",
			caserConversion: "kebab-to-pascal",
			input: model.NewString().
				WithQName("test-output/codegen/schema/util/string-02").
				WithMax(50).WithMin(1).WithRegex(".*"),
			expected: map[string]string{},
		},
		{
			name:            "integer 01",
			lang:            "go",
			caserConversion: "kebab-to-pascal",
			input: model.NewInteger().
				WithQName("test-output/codegen/schema/util/integer-01"),
			expected: map[string]string{},
		},
		{
			name:            "integer 02",
			lang:            "go",
			caserConversion: "kebab-to-pascal",
			input: model.NewInteger().
				WithQName("test-output/codegen/schema/util/integer-02").
				WithMultipleOf(5),
			expected: map[string]string{},
		},
		{
			name:            "integer 03",
			lang:            "go",
			caserConversion: "kebab-to-pascal",
			input: model.NewInteger().
				WithQName("test-output/codegen/schema/util/integer-03").
				WithRange(model.NewRange[int]().WithMax(10).WithMin(1)),
			expected: map[string]string{},
		},
		{
			name:            "integer 04",
			lang:            "go",
			caserConversion: "kebab-to-pascal",
			input: model.NewInteger().
				WithQName("test-output/codegen/schema/util/integer-04").
				WithMultipleOf(5).
				WithRange(model.NewRange[int]().WithMax(10).WithMin(1)),
			expected: map[string]string{},
		},
		{
			name:            "integer 05",
			lang:            "go",
			caserConversion: "kebab-to-pascal",
			input: model.NewInteger().
				WithQName("test-output/codegen/schema/util/integer-05").
				WithMultipleOf(5).
				WithRange(model.NewRange[int]().WithMax(10).WithMin(1)).
				WithRange(model.NewRange[int]().WithExclusiveMax(20).WithExclusiveMin(15)),
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
