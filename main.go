package main

import (
	"boundedinfinity/codegen/system"
	"os"
)

func main() {
	uris := os.Args[1:]

	s, err := system.New(
		system.OutputDir("/tmp/codegen"),
	)

	if err != nil {
		handleError(err)
	}

	if err := s.Load(uris...); err != nil {
		handleError(err)
	}

	if err := s.ProcessTemplates(); err != nil {
		handleError(err)
	}

	if err := s.Generate(); err != nil {
		handleError(err)
	}
}

func handleError(err error) {
	panic(err)
}
