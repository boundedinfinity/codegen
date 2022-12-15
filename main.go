package main

import (
	"boundedinfinity/codegen/system"
	"fmt"
	"os"
)

func main() {
	uris := os.Args[1:]

	s, err := system.New()

	if err != nil {
		handleError(err)
		return
	}

	if err := s.Load(uris...); err != nil {
		handleError(err)
		return
	}

	if err := s.ProcessTemplates(); err != nil {
		handleError(err)
		return
	}

	if err := s.Generate(); err != nil {
		handleError(err)
		return
	}
}

func handleError(err error) {
	fmt.Println(err.Error())
}
