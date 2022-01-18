package system

import (
	"boundedinfinity/codegen/model"
	"boundedinfinity/codegen/uritype"
	"boundedinfinity/codegen/util"
	"fmt"
	"io/fs"
	"io/ioutil"
	"path/filepath"

	"github.com/sethgrid/pester"
)

func (t *System) processUri(source *model.SourceInfo) error {
	if _, ok := t.sourceInfo[source.SourceUri]; ok {
		fmt.Printf("already processed: %v", source.SourceUri)
		return nil
	}

	if err := t.detectUriType(source.SourceUri, &source.UriType); err != nil {
		return err
	}

	switch source.UriType {
	case uritype.File:
		localPath := util.Uri2Path(source.SourceUri)
		file, err := util.IsFile(localPath)

		if err != nil {
			return err
		}

		if file {
			source.LocalPath = localPath
		} else {
			if err := filepath.WalkDir(localPath, t.processInputDir(source.InputUri)); err != nil {
				return err
			}
		}
	case uritype.Http, uritype.Https:
		cache, err := t.cacheUrl(source.SourceUri)

		if err != nil {
			return err
		}

		source.LocalPath = cache
	default:
		return uritype.ErrUriTypeInvalid
	}

	if err := t.detectMimeType(source.SourceUri, &source.MimeType); err != nil {
		return err
	}

	t.sourceInfo[source.SourceUri] = source

	return nil
}

func (t System) processInputDir(uri string) fs.WalkDirFunc {
	return func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		source := model.SourceInfo{
			InputUri:  uri,
			SourceUri: util.Path2Uri(path),
		}

		if err := t.processUri(&source); err != nil {
			return err
		}

		return nil
	}
}

func (t System) cacheUrl(uri string) (string, error) {
	if err := util.DirEnsure(model.CACHE_DIR); err != nil {
		return "", err
	}

	filename := filepath.Base(uri)
	localPath := filepath.Join(model.CACHE_DIR, filename)
	exist, err := util.PathExists(localPath)

	if err != nil {
		return localPath, err
	}

	if exist {
		return localPath, nil
	}

	resp, err := pester.Get(uri)

	if err != nil {
		return localPath, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return localPath, err
	}

	if err := ioutil.WriteFile(localPath, body, 0755); err != nil {
		return localPath, err
	}

	return localPath, nil
}

func (t System) cacheFile(source *model.SourceInfo) error {
	return nil
}
