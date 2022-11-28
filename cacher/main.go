package cacher

import (
	"boundedinfinity/codegen/check_sumer"
	"io/fs"

	o "github.com/boundedinfinity/go-commoner/optioner"
)

type Cacher struct {
	cacheDir          string
	checkSumFileMode  fs.FileMode
	overwriteModified o.Option[bool]
	checkSumer        *check_sumer.CheckSumer
	copyToDest        o.Option[bool]
	listFiles         []string
}

func New(opts ...Arg) (*Cacher, error) {
	cacher := &Cacher{
		listFiles: make([]string, 0),
	}

	for _, opt := range opts {
		opt(cacher)
	}

	if err := cacher.init(); err != nil {
		return nil, err
	}

	return cacher, nil
}
