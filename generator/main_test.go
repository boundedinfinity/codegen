package generator_test

import (
	"testing"

	jsmodel "github.com/boundedinfinity/go-jsonschema/model"
)

func createString() jsmodel.JsonSchemaString {
	// m := jsmodel.NewString("https://www.boundedinfinity.com/schema/string-1")
	m := jsmodel.JsonSchemaString{}
	return m
}

func Test_Generate_String(t *testing.T) {
	// input := createString()

	// assert.True(t, actual.Id.Defined())
}

// var _ = Describe("Generate Schema", func() {
// 	It("should work", func() {
// 		// n := optioner.NewStringValue("a Name")
// 		n := o.None[string]()
// 		s := model.JsonSchema{
// 			Id:         o.Some("https://www.boundedinfinity.com/schema/x/even-numbers"),
// 			Schema:     dialect,
// 			Type:       objecttype.NewObjectTypeValue(objecttype.Integer),
// 			MultipleOf: optioner.NewFloat64Value(2),
// 		}
// 		g, err := generator.New(map[string]string{
// 			"https://www.boundedinfinity.com/schema": "github.com/boundedinfinity",
// 		})
// 		Expect(err).To(BeNil())

// 		err = g.GenerateSchema(lang_ext.Go, n, s)
// 		Expect(err).To(BeNil())
// 	})
// })
