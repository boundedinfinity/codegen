package main

import (
	"boundedinfinity/codegen/generator"
	"boundedinfinity/codegen/loader"
	"fmt"
	"os"
)

func main() {
	// schemaPath := "../codegen-templates/spec.bi.yaml"
	schemaPath := os.Args[1]

	l := loader.New()

	if err := l.FromPath(schemaPath); err != nil {
		os.Exit(handleError(err))
	}

	g := generator.New(l.Output)

	if err := g.Generate(); err != nil {
		os.Exit(handleError(err))
	}
}

func handleError(err error) int {
	fmt.Printf("%v", err.Error())
	return 1
}
