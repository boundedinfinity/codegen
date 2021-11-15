package system

import (
	"boundedinfinity/codegen/model"
	"boundedinfinity/codegen/util"
	"io/ioutil"
	"path/filepath"

	"github.com/sethgrid/pester"
)

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
