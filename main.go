package main

import (
	"boundedinfinity/codegen/generator"
	"fmt"
	"os"
)

func main() {
	// config := "../openapi-parser-templates/go/server/echo/handlebars/project.yml"
	// config := "../openapi-parser-templates/go/server/echo/go/project.yml"
	// schemaPath := "../codegen-templates/openapi.yaml"
	schemaPath := "../codegen-templates/spec.bi.yaml"
	// schemaPath := ""

	// fmt.Println("=================================================================")
	// fmt.Println("=================================================================")

	g := generator.New(schemaPath)

	if err := g.Generate(); err != nil {
		os.Exit(handleError(err))
	}
}

func handleError(err error) int {
	fmt.Printf("%v", err.Error())
	return 1
}
