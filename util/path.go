package util

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/boundedinfinity/optional"
)

func FileExt(v string) (string, error) {
	ext := path.Ext(v)
	ext = strings.ReplaceAll(ext, ".", "")

	if ext == "" {
		return "", fmt.Errorf("no extention")
	}

	return ext, nil
}

func DirEnsure(v string) error {
	ok, err := DirExists(v)

	if err != nil {
		return err
	}

	if !ok {
		d := filepath.Dir(v)
		if err := os.MkdirAll(d, 0755); err != nil {
			return err
		}
	}

	return nil
}

func DirExists(v string) (bool, error) {
	d := filepath.Dir(v)
	return PathExists(d)
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

func AbsFromDirPath(r, p string) (string, error) {
	if filepath.IsAbs(p) {
		return p, nil
	}

	a := filepath.Join(r, p)

	if x, err := filepath.Abs(a); err != nil {
		return "", err
	} else {
		a = x
	}

	return a, nil
}

func AbsFromFilePath(r, p string) (string, error) {
	d := filepath.Dir(r)
	return AbsFromDirPath(d, p)
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
