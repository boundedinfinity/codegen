package model_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestTypes(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Model suite")
}

// var (
// 	stringJson1 = `type: string
// name: AString
// description: A string description
// example: a string
//     `

// 	stringExpected1 = model.StringTypeDescriptor{
// 		Name:        "AString",
// 		Type:        model.SchemaType_String,
// 		Description: "A string description",
// 		Example:     "a string",
// 	}

// 	longJson1 = `type: long
// name: ALong
// description: A long description
// example: 10
//     `

// 	longExpected1 = model.LongTypeDescriptor{
// 		Name:        "ALong",
// 		Type:        model.SchemaType_Long,
// 		Description: "A long description",
// 		Example:     10,
// 	}

// 	doubleJson1 = `type: double
// name: ADouble
// description: A double description
// example: 10.1
//     `

// 	doubleExpected1 = model.DoubleTypeDescriptor{
// 		Name:        "ADouble",
// 		Type:        model.SchemaType_Double,
// 		Description: "A double description",
// 		Example:     10.1,
// 	}

// 	arrayStringJson1 = `type: array
// name: AStringArray
// items: string
// description: A string array description
// example:
//     -   a
//     -   b
//     `

// 	arrayStringExpected1 = model.ArrayStringTypeDescriptor{
// 		Name:        "AStringArray",
// 		Type:        model.SchemaType_Array,
// 		Items:       model.SchemaType_String,
// 		Description: "A string array description",
// 		Example:     []string{"a", "b"},
// 	}

// 	arrayLongJson1 = `type: array
// name: ALongArray
// items: long
// description: A long array description
// example:
//     -   1
//     -   2
//     `

// 	arrayLongExpected1 = model.ArrayLongTypeDescriptor{
// 		Name:        "ALongArray",
// 		Type:        model.SchemaType_Array,
// 		Items:       model.SchemaType_Long,
// 		Description: "A long array description",
// 		Example:     []int32{1, 2},
// 	}
// )

// var _ = Describe("Smoke Test", func() {
// 	It("Umarshal StringSchemaTypeDescriptor", func() {
// 		input := stringJson1
// 		expected1 := stringExpected1
// 		var actual1 model.TypeDescriptor

// 		err := yaml.Unmarshal([]byte(input), &actual1)
// 		Expect(err).To(BeNil())
// 		Expect(actual1.Name).To(Equal(expected1.Name))
// 		Expect(actual1.Type).To(Equal(expected1.Type))

// 		var actual2 model.StringTypeDescriptor

// 		err = yaml.Unmarshal([]byte(input), &actual2)
// 		Expect(err).To(BeNil())
// 		Expect(actual2.Name).To(Equal(expected1.Name))
// 		Expect(actual2.Type).To(Equal(expected1.Type))
// 		Expect(actual2.Example).To(Equal(expected1.Example))
// 	})

// 	It("Umarshal LongSchemaTypeDescriptor", func() {
// 		input := longJson1
// 		expected1 := longExpected1
// 		var actual1 model.TypeDescriptor

// 		err := yaml.Unmarshal([]byte(input), &actual1)
// 		Expect(err).To(BeNil())
// 		Expect(actual1.Name).To(Equal(expected1.Name))
// 		Expect(actual1.Type).To(Equal(expected1.Type))

// 		var actual2 model.LongTypeDescriptor

// 		err = yaml.Unmarshal([]byte(input), &actual2)
// 		Expect(err).To(BeNil())
// 		Expect(actual2.Name).To(Equal(expected1.Name))
// 		Expect(actual2.Type).To(Equal(expected1.Type))
// 		Expect(actual2.Example).To(Equal(expected1.Example))
// 	})

// 	It("Umarshal DoubleSchemaTypeDescriptor", func() {
// 		input := doubleJson1
// 		expected1 := doubleExpected1
// 		var actual1 model.TypeDescriptor

// 		err := yaml.Unmarshal([]byte(input), &actual1)
// 		Expect(err).To(BeNil())
// 		Expect(actual1.Name).To(Equal(expected1.Name))
// 		Expect(actual1.Type).To(Equal(expected1.Type))

// 		var actual2 model.DoubleTypeDescriptor

// 		err = yaml.Unmarshal([]byte(input), &actual2)
// 		Expect(err).To(BeNil())
// 		Expect(actual2.Name).To(Equal(expected1.Name))
// 		Expect(actual2.Type).To(Equal(expected1.Type))
// 		Expect(actual2.Example).To(Equal(expected1.Example))
// 	})

// 	It("Umarshal ArrayStringSchemaTypeDescriptor", func() {
// 		input := arrayStringJson1
// 		expected1 := arrayStringExpected1
// 		var actual1 model.TypeDescriptor

// 		err := yaml.Unmarshal([]byte(input), &actual1)
// 		Expect(err).To(BeNil())
// 		Expect(actual1.Name).To(Equal(expected1.Name))
// 		Expect(actual1.Type).To(Equal(expected1.Type))

// 		var actual2 model.ArrayStringTypeDescriptor

// 		err = yaml.Unmarshal([]byte(input), &actual2)
// 		Expect(err).To(BeNil())
// 		Expect(actual2.Name).To(Equal(expected1.Name))
// 		Expect(actual2.Type).To(Equal(expected1.Type))
// 		Expect(actual2.Example).To(Equal(expected1.Example))
// 		Expect(actual2.Items).To(Equal(expected1.Items))
// 	})

// 	It("Umarshal ArrayLongSchemaTypeDescriptor", func() {
// 		input := arrayLongJson1
// 		expected1 := arrayLongExpected1
// 		var actual1 model.TypeDescriptor

// 		err := yaml.Unmarshal([]byte(input), &actual1)
// 		Expect(err).To(BeNil())
// 		Expect(actual1.Name).To(Equal(expected1.Name))
// 		Expect(actual1.Type).To(Equal(expected1.Type))

// 		var actual2 model.ArrayLongTypeDescriptor

// 		err = yaml.Unmarshal([]byte(input), &actual2)
// 		Expect(err).To(BeNil())
// 		Expect(actual2.Name).To(Equal(expected1.Name))
// 		Expect(actual2.Type).To(Equal(expected1.Type))
// 		Expect(actual2.Example).To(Equal(expected1.Example))
// 		Expect(actual2.Items).To(Equal(expected1.Items))
// 	})
// })
