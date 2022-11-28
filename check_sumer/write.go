package check_sumer

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/boundedinfinity/go-commoner/pather"
)

type CheckSumWriteCalculation struct {
	Path         string
	CheckSumPath string
	CheckSum     string
}

func (t *CheckSumer) Write(path string) (CheckSumWriteCalculation, error) {
	var result CheckSumWriteCalculation

	calc, err := t.CalculateFile(path)

	if err != nil {
		return result, err
	}

	result.Path = path
	result.CheckSum = calc.CheckSum
	result.CheckSumPath = t.calcDestPath(result.Path)

	if pather.PathExists(result.CheckSumPath) && t.overwrite.Defined() && !t.overwrite.Get() {
		if t.overwrite.Defined() && !t.overwrite.Get() {
			return result, fmt.Errorf("%v already exists", result.CheckSumPath)
		} else {
			if err := os.Remove(result.CheckSumPath); err != nil {
				return result, fmt.Errorf("%v %w", result.CheckSumPath, err)
			}
		}
	}

	if err := ioutil.WriteFile(result.CheckSumPath, []byte(result.CheckSum), t.fileMode); err != nil {
		return result, fmt.Errorf("%v %w", result.CheckSumPath, err)
	}

	return result, nil
}
