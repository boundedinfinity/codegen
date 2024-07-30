package generator_test

import (
	"boundedinfinity/codegen/generator"
	"boundedinfinity/codegen/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Generate(t *testing.T) {
	tcs := []struct {
		name     string
		lang     string
		input    model.CodeGenType
		expected map[string]string
		err      error
		newErr   error
	}{
		{
			name: "string 01",
			lang: "go",
			input: model.BuildString().
				Id("test-output/codegen/schema/util/string-01").
				Max(50).
				Min(1).
				Regex(".*").
				Build(),
			expected: map[string]string{},
		},
		{
			name: "integer 01",
			lang: "go",
			input: model.BuildInteger().
				Id("test-output/codegen/schema/util/integer-01").
				Min(2).
				Max(100).
				MultipleOf(5).
				Positive().
				Build(),
			expected: map[string]string{},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			gen, err := generator.New(tc.lang)
			assert.ErrorIs(tt, err, tc.newErr)

			// actual, err := gen.GenerateType(tc.input)
			_, err = gen.WriteType(tc.input)
			assert.ErrorIs(tt, err, tc.err)
			// assert.Equal(tt, tc.expected, actual)
		})
	}
}
