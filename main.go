package main

import (
	"boundedinfinity/codegen/loader"
	"boundedinfinity/codegen/util"
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

	fmt.Println(util.Jdump(l.OutputSpec))

	// g := generator.New(l.Output)

	// if err := g.Generate(); err != nil {
	// 	os.Exit(handleError(err))
	// }
}

func handleError(err error) int {
	fmt.Printf("%v\n", err.Error())
	return 1
}
