package main

import (
	"boundedinfinity/codegen/loader"
	"boundedinfinity/codegen/util"
	"fmt"
	"os"
)

func main() {
	// config := "../openapi-parser-templates/go/server/echo/handlebars/project.yml"
	// config := "../openapi-parser-templates/go/server/echo/go/project.yml"
	schemaPath := "../openapi-parser-templates/openapi.yaml"

	fmt.Println("=================================================================")
	fmt.Println("=================================================================")

	ldr := loader.New()
	rctx, err := ldr.Load(schemaPath, "")

	if err != nil {
		os.Exit(handleError(err))
	}

	fmt.Print(util.Jdump(rctx))

	// g := generator.New()

	// if err := g.Generate(rctx); err != nil {
	// 	os.Exit(handleError(err))
	// }
}

func handleError(err error) int {
	fmt.Printf("%v", err.Error())
	return 1
}
