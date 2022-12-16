package cacher

import (
	"github.com/boundedinfinity/go-commoner/slicer"
	"github.com/boundedinfinity/go-urischemer"
)

type CachedData struct {
	RootUri            string
	SourceUri          string
	SourceCheckSum     string
	SourceCheckSumPath string
	DestPath           string
	DestCheckSum       string
	DestCheckSumPath   string
	CalculatedCheckSum string
}

func (t Cacher) Cache(urls ...string) error {
	for _, rootUri := range slicer.Dedup(urls) {
		if t.source2Data.Has(rootUri) {
			continue
		}

		scheme, path, err := urischemer.Break(rootUri)

		if err != nil {
			return err
		}

		switch scheme {
		case urischemer.File:
			if err := t.cacheFilePath(rootUri, path); err != nil {
				return err
			}
		case urischemer.Http, urischemer.Https:
			if err := t.cacheHttpPath(rootUri, path); err != nil {
				return err
			}
		default:
			return urischemer.ErrUriSchemeNotFoundv(scheme)
		}
	}

	return nil
}
