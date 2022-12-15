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
	orig2Dest         mapper.Mapper[string, []string]
	dest2Orig         mapper.Mapper[string, string]
	orig2Data         mapper.Mapper[string, []*CachedData]
	source2Data       mapper.Mapper[string, *CachedData]
	dest2Data         mapper.Mapper[string, *CachedData]
}

func New(opts ...Arg) (*Cacher, error) {
	cacher := &Cacher{
		listFiles:   make([]string, 0),
		orig2Dest:   make(mapper.Mapper[string, []string]),
		dest2Orig:   make(mapper.Mapper[string, string]),
		orig2Data:   make(mapper.Mapper[string, []*CachedData]),
		source2Data: make(mapper.Mapper[string, *CachedData]),
		dest2Data:   make(mapper.Mapper[string, *CachedData]),
	}

	for _, opt := range opts {
		opt(cacher)
	}

	if err := cacher.init(); err != nil {
		return nil, err
	}

	return cacher, nil
}
