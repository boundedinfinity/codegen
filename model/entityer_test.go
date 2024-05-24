package model_test

import (
	"boundedinfinity/codegen/model"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Marshal_Jsoner(t *testing.T) {
	type Something struct {
		A string
		B int
		C float64
		D []string
		E []int
		F []float64
	}

	something := Something{
		A: "A",
		B: 0,
		D: []string{"X", "Y", "Z"},
	}

	jsoner := model.Entityer{}

	assert.Nil(t, jsoner.Value("a", &something.A))
	assert.Nil(t, jsoner.Value("b", &something.B))
	assert.Nil(t, jsoner.Value("d", &something.D))

	assert.Nil(t, jsoner.Fn("a", func() any { return something.A }))
	assert.Nil(t, jsoner.Fn("b", func() any { return something.B }))

	bs, err := jsoner.MarshalJSON()
	actual := string(bs)
	fmt.Println(actual)

	assert.Nil(t, err)
	// assert.JSONEq(t, "{}", actual)

	something.A = "AA"
	something.B = 11

	bs, err = jsoner.MarshalJSON()
	actual = string(bs)
	fmt.Println(actual)

	assert.Nil(t, err)
	// assert.JSONEq(t, "{}", actual)

	// tcs := []struct {
	// 	name     string
	// 	input    model.CodeGenType
	// 	expected string
	// 	err      error
	// }{}

	// for _, tc := range tcs {
	// 	t.Run(tc.name, func(tt *testing.T) {
	// 		// bs, err := model.MarshalCodeGenObject(tc.input)
	// 		bs, err := json.MarshalIndent(tc.input, "", "    ")
	// 		actual := string(bs)

	// 		if tc.err != nil {
	// 			assert.Equal(t, tc.err, err, "%v : %v", tc.name, actual)
	// 		} else {
	// 			assert.Nil(t, err, tc.name, actual)
	// 		}

	// 		assert.JSONEq(t, tc.expected, actual, "%v : %v", tc.name, actual)
	// 	})
	// }
}
