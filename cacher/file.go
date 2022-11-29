package cacher

import (
	"path/filepath"

	"github.com/boundedinfinity/go-commoner/pather"
	"github.com/boundedinfinity/go-commoner/slicer"
	"github.com/boundedinfinity/go-urischemer"
)

func (t Cacher) cacheFilePath(group string, path string) error {
	ok, err := pather.IsDir(path)

	if err != nil {
		return err
	}

	if ok {
		if err := t.cacheFileDir(group, path); err != nil {
			return err
		}
	} else {
		if err := t.cacheFileFile(group, path); err != nil {
			return err
		}
	}

	return nil
}

func (t Cacher) cacheFileDir(group string, root string) error {
	paths, err := pather.GetPaths(root)

	if err != nil {
		return err
	}

	paths = slicer.Filter(paths, func(p string) bool {
		return p != root
	})

	for _, path := range paths {
		if err := t.cacheFilePath(group, path); err != nil {
			return err
		}
	}

	return nil
}

func (t Cacher) cacheFileFile(group string, path string) error {
	var data CachedData

	data.SourceUrl = urischemer.Combine(urischemer.File, path)

	if clean, err := urischemer.Clean(data.SourceUrl); err != nil {
		return err
	} else {
		data.SourceUrl = clean
	}

	data.DestPath = filepath.Clean(path)

	if r, err := t.checkSumer.Read(data.DestPath); err != nil {
		return err
	} else {
		data.DestCheckSumPath = r.Path
		data.DestCheckSum = r.CheckSum
		data.SourceCheckSum = r.CheckSum
		data.SourceCheckSumPath = r.Path
	}

	if r, err := t.checkSumer.CalculateFile(data.DestPath); err != nil {
		return err
	} else {
		data.CalculateCheckSum = r.CheckSum
	}

	if !t.groupMap.Has(group) {
		t.groupMap[group] = make([]*CachedData, 0)
	}

	t.groupMap[group] = append(t.groupMap[group], &data)

	if !t.sourceMap.Has(data.SourceUrl) {
		t.sourceMap[data.SourceUrl] = &data
	} else {
		// TODO
	}

	if !t.destMap.Has(data.DestPath) {
		t.destMap[data.DestPath] = &data
	} else {
		// TODO
	}

	return nil
}
