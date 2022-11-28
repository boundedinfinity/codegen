package main

import (
	"boundedinfinity/codegen/system"
	"os"
)

func main() {
	schemaPaths := os.Args[1:]

	s, err := system.New()

	if err != nil {
		handleError(err)
	}

	if err := s.LoadUri(schemaPaths...); err != nil {
		handleError(err)
	}

	if err := s.Check(); err != nil {
		handleError(err)
	}

	// if err := s.Generate(); err != nil {
	// 	os.Exit(handleError(err))
	// }
}

func handleError(err error) {
	panic(err)
}
