package cacher

import (
	"boundedinfinity/codegen/model"
	"path/filepath"

	"github.com/boundedinfinity/go-commoner/pather"
	"github.com/boundedinfinity/go-commoner/slicer"
	"github.com/boundedinfinity/go-urischemer"
)

func (t Cacher) cacheFilePath(orig, root string) error {
	paths := make([]string, 0)

	ok, err := pather.IsDir(root)

	if err != nil {
		return err
	}

	if ok {
		ps, err := pather.GetFiles(root)

		if err != nil {
			return err
		}

		paths = append(paths, ps...)
	}

	paths = slicer.Filter(paths, func(p string) bool {
		return p != root
	})

	paths = slicer.Dedup(paths)

	for _, path := range paths {
		if err := t.cacheFileFile(orig, path); err != nil {
			return err
		}
	}

	return nil
}

func (t Cacher) cacheFileFile(orig, path string) error {
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

	t.source2Data[data.SourceUrl] = &data
	t.dest2Data[data.DestPath] = &data
	t.dest2Orig[data.DestPath] = orig
	model.MapListAdd(t.orig2Dest, orig, data.DestPath)
	model.MapListAdd(t.orig2Data, orig, &data)

	return nil
}
