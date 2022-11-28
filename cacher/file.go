package cacher

import (
	"path/filepath"

	"github.com/boundedinfinity/go-commoner/pather"
	"github.com/boundedinfinity/go-commoner/slicer"
	"github.com/boundedinfinity/go-urischemer"
)

func (t Cacher) cacheFilePath(path string) ([]CachedData, error) {
	var results []CachedData

	ok, err := pather.IsDir(path)

	if err != nil {
		return results, err
	}

	if ok {
		if rs, err := t.cacheFileDir(path); err != nil {
			return results, err
		} else {
			results = append(results, rs...)
		}
	} else {
		if r, err := t.cacheFileFile(path); err != nil {
			return results, err
		} else {
			results = append(results, r)
		}
	}

	return results, nil
}

func (t Cacher) cacheFileDir(root string) ([]CachedData, error) {
	var results []CachedData
	paths, err := pather.GetPaths(root)

	if err != nil {
		return results, err
	}

	paths = slicer.Filter(paths, func(p string) bool {
		return p != root
	})

	for _, path := range paths {
		if rs, err := t.cacheFilePath(path); err != nil {
			return results, err
		} else {
			results = append(results, rs...)
		}
	}

	return results, nil
}

func (t Cacher) cacheFileFile(path string) (CachedData, error) {
	var data CachedData

	data.SourceUrl = urischemer.Combine(urischemer.File, path)

	if clean, err := urischemer.Clean(data.SourceUrl); err != nil {
		return data, err
	} else {
		data.SourceUrl = clean
	}

	data.DestPath = filepath.Clean(path)

	if r, err := t.checkSumer.Read(data.DestPath); err != nil {
		return data, err
	} else {
		data.DestCheckSumPath = r.Path
		data.DestCheckSum = r.CheckSum
		data.SourceCheckSum = r.CheckSum
		data.SourceCheckSumPath = r.Path
	}

	if r, err := t.checkSumer.CalculateFile(data.DestPath); err != nil {
		return data, err
	} else {
		data.CalculateCheckSum = r.CheckSum
	}

	return data, nil
}
