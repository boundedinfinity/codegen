package check_sumer

import (
	"io/fs"
	"path/filepath"

	"github.com/boundedinfinity/go-commoner/optioner"
)

type Arg func(*CheckSumer)

func CheckSumFileMode(v fs.FileMode) Arg {
	return func(t *CheckSumer) {
		t.fileMode = v
	}
}

func Ext(v string) Arg {
	return func(t *CheckSumer) {
		t.ext = v
	}
}

func Overwrite(v bool) Arg {
	return func(t *CheckSumer) {
		t.overwrite = optioner.Some(v)
	}
}

func SourceDir(v string) Arg {
	return func(t *CheckSumer) {
		t.sourceDir = optioner.Some(v)
	}
}

func DestDir(v string) Arg {
	return func(t *CheckSumer) {
		t.destDir = optioner.Some(v)
	}
}

func Algo(v CheckSumAlgo) Arg {
	return func(t *CheckSumer) {
		t.algo = optioner.Some(v)
	}
}

const (
	DEFAULT_FILE_MODE = fs.FileMode(0644)
	DEFAULT_ALGO      = Md5
)

func (t *CheckSumer) init() error {
	if t.fileMode == 0 {
		t.fileMode = DEFAULT_FILE_MODE
	}

	if t.sourceDir.Defined() {
		if path, err := filepath.Abs(t.sourceDir.Get()); err != nil {
			return err
		} else {
			t.sourceDir = optioner.Some(path)
		}
	}

	if t.destDir.Defined() {
		if path, err := filepath.Abs(t.destDir.Get()); err != nil {
			return err
		} else {
			t.destDir = optioner.Some(path)
		}
	}

	if t.algo.Empty() {
		t.algo = optioner.Some(DEFAULT_ALGO)
	}

	if t.ext == "" {
		t.ext = "." + string(t.algo.Get())
	}

	return nil
}
