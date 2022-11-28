package check_sumer

import (
	"path/filepath"
	"strings"

	"github.com/boundedinfinity/go-commoner/pather"
)

func (t *CheckSumer) calcDestPath(path string) string {
	dest := path
	dest = filepath.Clean(dest)

	switch {
	case t.sourceDir.Defined() && t.destDir.Defined():
		dest = strings.ReplaceAll(dest, t.sourceDir.Get(), t.destDir.Get())
	case t.destDir.Defined():
		dest = filepath.Join(t.destDir.Get(), pather.Base(dest))
	default:
		// Don't modify the dest
	}

	if !strings.HasSuffix(dest, t.ext) {
		dest += t.ext
	}

	return dest
}
