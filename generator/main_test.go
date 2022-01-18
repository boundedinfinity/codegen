package generator_test

import (
	"boundedinfinity/codegen/generator"
	"boundedinfinity/codegen/lang_ext"
	"testing"

	"github.com/boundedinfinity/jsonschema"
	"github.com/boundedinfinity/jsonschema/objecttype"
	"github.com/boundedinfinity/optioner"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	dialect = optioner.NewStringValue("https://json-schema.org/draft/2020-12/schema")
)

func TestTypes(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Generator Suite")
}

var _ = Describe("Generate Schema", func() {
	It("should work", func() {
		// n := optioner.NewStringValue("a Name")
		n := optioner.NewStringEmpty()
		s := jsonschema.JsonSchmea{
			Id:         optioner.NewStringValue("https://www.boundedinfinity.com/schema/x/even-numbers"),
			Schema:     dialect,
			Type:       objecttype.NewObjectTypeValue(objecttype.Integer),
			MultipleOf: optioner.NewFloat64Value(2),
		}
		g, err := generator.New(map[string]string{
			"https://www.boundedinfinity.com/schema": "github.com/boundedinfinity",
		})
		Expect(err).To(BeNil())

		err = g.GenerateSchema(lang_ext.Go, n, s)
		Expect(err).To(BeNil())
	})
})
