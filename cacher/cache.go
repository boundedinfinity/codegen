package cacher

import (
	"path/filepath"

	"github.com/boundedinfinity/go-commoner/optioner"
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
	return t.CacheWithBase(group, optioner.None[string](), urls...)
}

func (t Cacher) CacheWithBase(group string, dir optioner.Option[string], urls ...string) error {
	for _, sourceUrl := range urls {
		scheme, path, err := urischemer.Break(sourceUrl)

		if err != nil {
			return err
		}

		if scheme == urischemer.File && !filepath.IsAbs(path) && dir.Defined() {
			rel := path
			path = filepath.Join(dir.Get(), rel)
			path, err = filepath.Abs(path)

			if err != nil {
				return err
			}
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
