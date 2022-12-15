package cacher

import (
	"github.com/boundedinfinity/go-commoner/environmenter"
	"github.com/boundedinfinity/go-commoner/slicer"
	"github.com/boundedinfinity/go-urischemer"
)

type CachedData struct {
	SourceUrl          string
	SourceCheckSum     string
	SourceCheckSumPath string
	DestPath           string
	DestCheckSum       string
	DestCheckSumPath   string
	CalculateCheckSum  string
}

func (t Cacher) Cache(urls ...string) error {
	rUrls := slicer.Map(urls, environmenter.Sub)

	for _, sourceUrl := range slicer.Dedup(rUrls) {
		if t.source2Data.Has(sourceUrl) {
			continue
		}

		scheme, path, err := urischemer.Break(sourceUrl)

		if err != nil {
			return err
		}

		switch scheme {
		case urischemer.File:
			if err := t.cacheFilePath(sourceUrl, path); err != nil {
				return err
			}
		case urischemer.Http, urischemer.Https:
			if err := t.cacheHttpPath(sourceUrl, path); err != nil {
				return err
			}
		default:
			return urischemer.ErrUriSchemeNotFoundv(scheme)
		}
	}

	return nil
}
