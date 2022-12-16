package cacher

import (
	"boundedinfinity/codegen/util"
	"path/filepath"

	"github.com/boundedinfinity/go-commoner/pather"
	"github.com/boundedinfinity/go-commoner/slicer"
	"github.com/boundedinfinity/go-urischemer"
)

func (t Cacher) cacheFilePath(rootUri, path string) error {
	paths := make([]string, 0)

	ok, err := pather.IsDir(path)

	if err != nil {
		return err
	}

	if ok {
		ps, err := pather.GetFiles(path)

		if err != nil {
			return err
		}

		paths = append(paths, ps...)
	}

	paths = slicer.Filter(paths, func(p string) bool {
		return p != path
	})

	paths = slicer.Dedup(paths)

	for _, path := range paths {
		if err := t.cacheFileFile(rootUri, path); err != nil {
			return err
		}
	}

	return nil
}

func (t Cacher) cacheFileFile(rootUri, path string) error {
	data := CachedData{
		RootUri: rootUri,
	}

	data.SourceUri = urischemer.Combine(urischemer.File, path)

	if clean, err := urischemer.Clean(data.SourceUri); err != nil {
		return err
	} else {
		data.SourceUri = clean
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
		data.CalculatedCheckSum = r.CheckSum
	}

	t.source2Data[data.SourceUri] = &data
	t.dest2Data[data.DestPath] = &data
	t.dest2Orig[data.DestPath] = data.DestPath
	util.MapListAdd(t.orig2Dest, data.RootUri, data.DestPath)
	util.MapListAdd(t.orig2Data, data.RootUri, &data)

	return nil
}
