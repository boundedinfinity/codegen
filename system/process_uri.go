package system

import (
	"boundedinfinity/codegen/model"
	"boundedinfinity/codegen/uri_scheme"
	"fmt"

	"github.com/boundedinfinity/go-marshaler"
)

func (t *System) processUri(uri string) error {
	path, scheme, err := uri_scheme.PathErr(uri)

	if err != nil {
		return err
	}

	source := model.SourceInfo{
		InputUri: uri,
		Path:     path,
		Scheme:   scheme,
	}

	if _, ok := t.sourceInfo[path]; ok {
		fmt.Printf("already processed: %v", path)
		return nil
	}

	m := map[string]jsonschema.Sch

	marshaler.UnmarshalFromPath()

	t.sourceInfo[source.SourceUri] = &source

	return nil
}

// func (t System) processInputDir(uri string) fs.WalkDirFunc {
// 	return func(path string, d fs.DirEntry, err error) error {
// 		if err != nil {
// 			return err
// 		}

// 		if d.IsDir() {
// 			return nil
// 		}

// 		if err := t.processUri(util.Path2Uri(path)); err != nil {
// 			return err
// 		}

// 		return nil
// 	}
// }

// func (t System) cacheUrl(uri string) (string, error) {
// 	if err := pather.DirEnsure(model.CACHE_DIR); err != nil {
// 		return "", err
// 	}

// 	filename := filepath.Base(uri)
// 	localPath := filepath.Join(model.CACHE_DIR, filename)
// 	exist, err := pather.PathExistsErr(localPath)

// 	if err != nil {
// 		return localPath, err
// 	}

// 	if exist {
// 		return localPath, nil
// 	}

// 	resp, err := pester.Get(uri)

// 	if err != nil {
// 		return localPath, err
// 	}

// 	defer resp.Body.Close()

// 	body, err := ioutil.ReadAll(resp.Body)

// 	if err != nil {
// 		return localPath, err
// 	}

// 	if err := ioutil.WriteFile(localPath, body, 0755); err != nil {
// 		return localPath, err
// 	}

// 	return localPath, nil
// }

// func (t System) cacheFile(source *model.SourceInfo) error {
// 	return nil
// }
