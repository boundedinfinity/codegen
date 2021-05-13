package model_test

import (
	"fmt"
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestLoader(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Model suite")
}

var _ = Describe("Smoke Test", func() {
	It("Deserialize", func() {
		d, _ := os.Getwd()
		fmt.Printf("dir: %v", d)
		// loader := loader.New()
		// err := loader.FromPath("")

		// Expect(err).To(BeNil())

		// err = util.MarshalIndentToFile("output.json", loader.Output, "", "    ")

		// Expect(err).To(BeNil())
	})
})
