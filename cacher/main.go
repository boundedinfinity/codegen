package cacher

import (
	"boundedinfinity/codegen/check_sumer"
	"io/fs"

	o "github.com/boundedinfinity/go-commoner/optioner"
	"github.com/boundedinfinity/go-commoner/optioner/mapper"
)

type Cacher struct {
	cacheDir          string
	checkSumFileMode  fs.FileMode
	overwriteModified o.Option[bool]
	checkSumer        *check_sumer.CheckSumer
	copyToDest        o.Option[bool]
	listFiles         []string
	sourceMap         mapper.Mapper[string, *CachedData]
	destMap           mapper.Mapper[string, *CachedData]
	groupMap          mapper.Mapper[string, []*CachedData]
}

func New(opts ...Arg) (*Cacher, error) {
	cacher := &Cacher{
		listFiles: make([]string, 0),
		sourceMap: make(mapper.Mapper[string, *CachedData]),
		destMap:   make(mapper.Mapper[string, *CachedData]),
		groupMap:  make(mapper.Mapper[string, []*CachedData]),
	}

	for _, opt := range opts {
		opt(cacher)
	}

	if err := cacher.init(); err != nil {
		return nil, err
	}

	return cacher, nil
}
