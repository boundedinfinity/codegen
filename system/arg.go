package system

import (
	"github.com/boundedinfinity/go-commoner/optioner"
)

var (
	DEFAULT_WORKDIR_NAME = "codegen"
)

type Arg func(*System)

func OutputDir(v string) Arg {
	return func(t *System) {
		t.outputDir = optioner.Some(v)
	}
}
