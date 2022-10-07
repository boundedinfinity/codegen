package main

import (
	"boundedinfinity/codegen/system"
	"fmt"
	"os"
)

func main() {
	schemaPaths := os.Args[1:]

	s := system.New()

	if err := s.Process(schemaPaths...); err != nil {
		os.Exit(handleError(err))
	}

	// if err := s.Generate(); err != nil {
	// 	os.Exit(handleError(err))
	// }
}

func handleError(err error) int {
	fmt.Printf("%v\n", err.Error())
	return 1
}
