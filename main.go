package main

import (
	"boundedinfinity/codegen/system"
	"fmt"
	"os"
)

//go:generate enumer -standalone=true -package=uritype -name=UriType -items=file,http,https

func main() {
	// schemaPath := "../codegen-templates/spec.bi.yaml"
	schemaPaths := os.Args[1:]

	l := system.New()

	if err := l.Load(schemaPaths...); err != nil {
		handleError(err)
	}

	// if err := g.Generate(); err != nil {
	// 	os.Exit(handleError(err))
	// }
}

func handleError(err error) int {
	fmt.Printf("%v\n", err.Error())
	return 1
}
