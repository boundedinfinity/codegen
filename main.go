package main

import (
	"boundedinfinity/codegen/generator"
	"boundedinfinity/codegen/loader"
	"fmt"
	"os"
)

func main() {
	// schemaPath := "../codegen-templates/spec.bi.yaml"
	schemaPaths := os.Args[1:]

	l := loader.New()

	if err := l.FromPath(schemaPaths); err != nil {
		os.Exit(handleError(err))
	}

	g := generator.New(l.Output)

	if err := g.Generate(); err != nil {
		os.Exit(handleError(err))
	}
}

func handleError(err error) int {
	fmt.Printf("%v\n", err.Error())
	return 1
}
