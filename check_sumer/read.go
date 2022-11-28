package check_sumer

import (
	"io/ioutil"

	"github.com/boundedinfinity/go-commoner/pather"
)

type CheckSumRead struct {
	Path     string
	CheckSum string
}

func (t *CheckSumer) Read(path string) (CheckSumRead, error) {
	var result CheckSumRead

	result.Path = t.calcDestPath(path)
	ok, err := pather.PathExistsErr(result.Path)

	if err != nil {
		return result, err
	}

	if ok {
		if bs, err := ioutil.ReadFile(result.Path); err != nil {
			return result, err
		} else {
			result.CheckSum = string(bs)
		}
	}

	return result, nil
}
