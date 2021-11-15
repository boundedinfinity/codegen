package util

import (
	"boundedinfinity/codegen/uritype"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/boundedinfinity/optional"
)

var (
	fileUriPrefix = fmt.Sprintf("%v://", uritype.File.String())
)

func Uri2Path(v string) string {
	return strings.ReplaceAll(v, fileUriPrefix, "")
}

func Path2Uri(v string) string {
	return fmt.Sprintf("%v%v", fileUriPrefix, v)
}

func DirEnsure(v string) error {
	ok, err := PathExists(v)

	if err != nil {
		return err
	}

	if !ok {
		if err := os.MkdirAll(v, 0755); err != nil {
			return err
		}
	}

	return nil
}

func PathExists(v string) (bool, error) {
	_, err := os.Stat(v)

	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		} else {
			return false, err
		}
	}

	return true, nil
}

func IsFile(p string) (bool, error) {
	info, err := os.Stat(p)

	if err != nil {
		return false, err
	}

	return info.Mode().IsRegular(), nil
}

func IsDir(p string) (bool, error) {
	ok, err := IsFile(p)

	if err != nil {
		return false, err
	}

	return !ok, nil
}

func FileSearch(v optional.StringOptional, rs ...optional.StringOptional) (string, bool) {
	if v.IsEmpty() {
		return "", false
	}

	abs, err := filepath.Abs(v.Get())

	if err == nil {
		ok, err := PathExists(abs)

		if err == nil && ok {
			return abs, true
		}
	}

	if rs != nil {
		for _, r := range rs {
			if r.IsEmpty() {
				continue
			}

			p := filepath.Join(r.Get(), v.Get())
			abs, err := filepath.Abs(p)

			if err != nil {
				continue
			}

			ok, err := PathExists(abs)

			if err == nil && ok {
				return abs, true
			}
		}
	}

	return "", false
}
