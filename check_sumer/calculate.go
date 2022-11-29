package check_sumer

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"

	"github.com/boundedinfinity/go-commoner/pather"
)

type CheckSumCalculation struct {
	Path     string
	CheckSum string
}

func (t *CheckSumer) CalculatePath(path string) ([]CheckSumCalculation, error) {
	var results []CheckSumCalculation

	if ok, err := pather.IsDir(path); err != nil {
		return results, err
	} else {
		if ok {
			if subs, err := t.CalculateDir(path); err != nil {
				return results, err
			} else {
				results = append(results, subs...)
			}
		} else {
			if sub, err := t.CalculateFile(path); err != nil {
				return results, err
			} else {
				results = append(results, sub)
			}
		}

	}

	return results, nil
}

func (t *CheckSumer) CalculateDir(path string) ([]CheckSumCalculation, error) {
	var results []CheckSumCalculation

	if ok, err := pather.IsDir(path); err != nil {
		return results, err
	} else {
		if !ok {
			return results, nil
		}
	}

	subPaths, err := pather.GetPaths(path)

	if err != nil {
		return results, err
	}

	for _, subPath := range subPaths {
		if subs, err := t.CalculatePath(subPath); err != nil {
			return results, err
		} else {
			results = append(results, subs...)
		}
	}

	return results, nil
}

func (t *CheckSumer) CalculateFile(path string) (CheckSumCalculation, error) {
	var result CheckSumCalculation

	if err := pather.IsFileErr(path); err != nil {
		return result, err
	}

	result.Path = path
	file, err := os.Open(result.Path)

	if err != nil {
		return result, err
	}

	defer file.Close()

	hash := md5.New()

	_, err = io.Copy(hash, file)

	if err != nil {
		return result, err
	}

	result.CheckSum = hex.EncodeToString(hash.Sum(nil))

	return result, nil
}
