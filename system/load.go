package system

import (
	"boundedinfinity/codegen/model"
	"boundedinfinity/codegen/uritype"
	"boundedinfinity/codegen/util"
	"io/fs"
	"path/filepath"
)

func (t *System) Load(uris ...string) error {
	if err := t.load1(uris...); err != nil {
		return err
	}

	if err := t.load2(); err != nil {
		return err
	}

	if err := t.load3(); err != nil {
		return err
	}

	return nil
}

func (t *System) load1(uris ...string) error {
	for _, uri := range uris {
		if _, ok := t.sourceInfo[uri]; ok {
			return model.ErrDuplicateSourceUriV(uri)
		}

		info := model.SourceInfo{
			SourceUri: uri,
		}

		t.sourceInfo[uri] = &info

		if err := t.detectUriType(uri, &info.UriType); err != nil {
			return err
		}
	}

	return nil
}

func (t *System) load2() error {
	for uri, info := range t.sourceInfo {
		switch info.UriType {
		case uritype.File:
			localPath := util.Uri2Path(uri)
			file, err := util.IsFile(localPath)

			if err != nil {
				return err
			}

			if file {
				t.sourceInfo[uri].LocalPath = localPath
			} else {
				fn := func(path string, d fs.DirEntry, err error) error {
					if err != nil {
						return err
					}

					if d.IsDir() {
						return nil
					}

					localUri := util.Path2Uri(path)

					if _, ok := t.sourceInfo[localUri]; ok {
						return model.ErrDuplicateSourceUriV(uri)
					}

					t.sourceInfo[localUri] = &model.SourceInfo{
						SourceUri: localUri,
						UriType:   uritype.File,
						LocalPath: path,
					}

					return nil
				}

				if err := filepath.WalkDir(localPath, fn); err != nil {
					return err
				}

				delete(t.sourceInfo, uri)
			}
		case uritype.Http, uritype.Https:
			cache, err := t.cacheUrl(uri)

			if err != nil {
				return err
			}

			t.sourceInfo[uri] = info
			t.sourceInfo[uri].LocalPath = cache
		default:
			return uritype.ErrUriTypeInvalid
		}
	}

	return nil
}

func (t *System) load3() error {
	for uri, info := range t.sourceInfo {
		if err := t.detectMimeType(uri, &info.MimeType); err != nil {
			return err
		}
	}

	return nil
}
