package model_test

import (
	"boundedinfinity/codegen/loader"
	"boundedinfinity/codegen/util"
	"fmt"
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBcd(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Model suite")
}

var _ = Describe("Smoke Test", func() {
	It("Deserialize", func() {
		d, _ := os.Getwd()
		fmt.Printf("dir: %v", d)
		loader := loader.New()
		err := loader.FromPath("/Users/bbabb200/dev/github.comcast.com/rubicon-specer/specification2/spec.yaml")

		Expect(err).To(BeNil())

		err = util.MarshalIndentToFile("output.json", loader.Output, "", "    ")

		Expect(err).To(BeNil())
	})
})
