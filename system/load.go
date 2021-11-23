package system

import (
	"boundedinfinity/codegen/model"
	"boundedinfinity/codegen/uritype"
	"boundedinfinity/codegen/util"
	"fmt"
	"io/fs"
	"path/filepath"
)

func (t *System) processUri(uri string) error {
	if _, ok := t.sourceInfo[uri]; ok {
		fmt.Printf("already processed: %v", uri)
		return nil
	}

	info := model.SourceInfo{
		SourceUri: uri,
	}

	if err := t.detectUriType(uri, &info.UriType); err != nil {
		return err
	}

	switch info.UriType {
	case uritype.File:
		localPath := util.Uri2Path(uri)
		file, err := util.IsFile(localPath)

		if err != nil {
			return err
		}

		if file {
			info.LocalPath = localPath
		} else {
			fn := func(path string, d fs.DirEntry, err error) error {
				if err != nil {
					return err
				}

				if d.IsDir() {
					return nil
				}

				localUri := util.Path2Uri(path)

				if err := t.processUri(localUri); err != nil {
					return err
				}

				return nil
			}

			if err := filepath.WalkDir(localPath, fn); err != nil {
				return err
			}
		}
	case uritype.Http, uritype.Https:
		cache, err := t.cacheUrl(uri)

		if err != nil {
			return err
		}

		info.LocalPath = cache
	default:
		return uritype.ErrUriTypeInvalid
	}

	if err := t.detectMimeType(uri, &info.MimeType); err != nil {
		return err
	}

	t.sourceInfo[uri] = &info

	return nil
}
