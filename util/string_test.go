package util_test

import (
	"boundedinfinity/codegen/util"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestString(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "String suite")
}

var _ = Describe("Normalize", func() {
	It("should work", func() {
		expected := "_23Te_st"
		actual := util.Normalize("123Te!st")
		Expect(actual).To(Equal(expected))
	})
})
