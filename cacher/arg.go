package cacher

import (
	"boundedinfinity/codegen/check_sumer"
	"io/fs"
	"os"
	"path/filepath"

	o "github.com/boundedinfinity/go-commoner/optioner"
	"github.com/boundedinfinity/go-commoner/pather"
	"github.com/boundedinfinity/go-commoner/slicer"
)

type Arg func(*Cacher)

func CacheDir(v string) Arg {
	return func(t *Cacher) {
		t.cacheDir = v
	}
}

func OverwriteModified(v bool) Arg {
	return func(t *Cacher) {
		t.overwriteModified = o.Some(v)
	}
}

func CheckSummer(v *check_sumer.CheckSumer) Arg {
	return func(t *Cacher) {
		t.checkSumer = v
	}
}

func CopyToDest(v bool) Arg {
	return func(t *Cacher) {
		t.copyToDest = o.Some(v)
	}
}

func ListFiles(v ...string) Arg {
	return func(t *Cacher) {
		t.listFiles = append(t.listFiles, v...)
		t.listFiles = slicer.Dedup(t.listFiles)
	}
}

const (
	DEFAULT_CHECKSUM_FILE_MODE = fs.FileMode(0644)
	DEFAULT_LIST_FILE          = "cacher.list"
	DEFAULT_RELATIVE_DIR       = ".config/cacher"
)

func (t *Cacher) init() error {
	if t.cacheDir == "" {
		homeDir := os.Getenv("HOME")

		if homeDir != "" {
			t.cacheDir = filepath.Join(os.Getenv("HOME"), DEFAULT_RELATIVE_DIR)
		} else {
			t.cacheDir = filepath.Join(os.TempDir(), "cacher")
		}
	}

	if err := pather.DirEnsure(t.cacheDir); err != nil {
		return err
	}

	if t.checkSumFileMode == 0 {
		t.checkSumFileMode = DEFAULT_CHECKSUM_FILE_MODE
	}

	if t.checkSumer == nil {
		v, err := check_sumer.New(
			check_sumer.DestDir(t.cacheDir),
			check_sumer.CheckSumFileMode(t.checkSumFileMode),
		)

		if err != nil {
			return err
		}

		t.checkSumer = v
	}

	if len(t.listFiles) == 0 {
		t.listFiles = append(t.listFiles, DEFAULT_LIST_FILE)
	}

	return nil
}
