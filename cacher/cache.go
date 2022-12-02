package cacher

import (
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

func (t Cacher) Cache(group string, urls ...string) error {
	for _, sourceUrl := range urls {
		scheme, path, err := urischemer.Break(sourceUrl)

		if err != nil {
			return err
		}

		switch scheme {
		case urischemer.File:
			if err := t.cacheFilePath(group, path); err != nil {
				return err
			}
		case urischemer.Http, urischemer.Https:
			if err := t.cacheHttpPath(group, path); err != nil {
				return err
			}
		default:
			return urischemer.ErrUriSchemeNotFoundv(scheme)
		}
	}

	return nil
}
