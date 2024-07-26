package model_test

import (
	"boundedinfinity/codegen/model"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Marshal_String(t *testing.T) {
	tcs := []struct {
		name     string
		input    *model.CodeGenString
		expected string
		err      error
	}{
		{
			name:  "Serialize boolean",
			input: model.BuildString().Build(),
			err:   nil,
			expected: `{
		        "base-type": "string"
		    }`,
		},
		{
			name: "Serialize boolean",
			input: model.BuildString().
				Name("A_STRING").
				Description("an object description").
				Build(),
			err: nil,
			expected: `{
		        "base-type": "string",
                "name": "A_STRING",
                "description": "an object description"
		    }`,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			bs, err := json.MarshalIndent(&tc.input, "", "    ")
			actual := string(bs)

			assert.ErrorIs(tt, err, tc.err)
			assert.JSONEqf(tt, tc.expected, actual, actual)
		})
	}
}
