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

func (t Cacher) Cache(urls ...string) ([]CachedData, error) {
	var results []CachedData

	for _, sourceUrl := range urls {
		scheme, path, err := urischemer.Break(sourceUrl)

		if err != nil {
			return results, err
		}

		switch scheme {
		case urischemer.File:
			if rs, err := t.cacheFilePath(path); err != nil {
				return results, err
			} else {
				results = append(results, rs...)
			}
		case urischemer.Http, urischemer.Https:
			if rs, err := t.cacheHttpPath(path); err != nil {
				return results, err
			} else {
				results = append(results, rs...)
			}
		default:
			return results, urischemer.ErrUriSchemeNotFoundv(scheme)
		}
	}

	return results, nil
}
