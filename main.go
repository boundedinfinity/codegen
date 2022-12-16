package main

import (
	"boundedinfinity/codegen/system"
	"fmt"
	"os"
)

func main() {
	paths := os.Args[1:]

	s, err := system.New()

	if err != nil {
		handleError(err)
		return
	}

	if err := s.Process(paths...); err != nil {
		handleError(err)
		return
	}

}

func handleError(err error) {
	fmt.Println(err.Error())
}
