package check_sumer

import (
	"io/fs"

	"github.com/boundedinfinity/go-commoner/optioner"
)

type CheckSumer struct {
	fileMode  fs.FileMode
	ext       string
	overwrite optioner.Option[bool]
	sourceDir optioner.Option[string]
	destDir   optioner.Option[string]
	algo      optioner.Option[CheckSumAlgo]
}

func New(opts ...Arg) (*CheckSumer, error) {
	checkSummer := &CheckSumer{}

	for _, opt := range opts {
		opt(checkSummer)
	}

	if err := checkSummer.init(); err != nil {
		return checkSummer, nil
	}

	return checkSummer, nil
}
