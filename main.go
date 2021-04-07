package main

import (
	"boundedinfinity/codegen/generator"
	"boundedinfinity/codegen/loader"
	"fmt"
	"os"
)

func main() {
	// config := "../openapi-parser-templates/go/server/echo/handlebars/project.yml"
	// config := "../openapi-parser-templates/go/server/echo/go/project.yml"
	schemaPath := "../codegen-templates/openapi.yaml"

	// fmt.Println("=================================================================")
	// fmt.Println("=================================================================")

	ldr := loader.New()
	rc, err := ldr.Load(schemaPath)

	if err != nil {
		os.Exit(handleError(err))
	}

	// fmt.Print(util.Jdump(rc))

	g := generator.New()

	if err := g.Generate(rc); err != nil {
		os.Exit(handleError(err))
	}
}

func handleError(err error) int {
	fmt.Printf("%v", err.Error())
	return 1
}
