package util

import (
	"path/filepath"

	o "github.com/boundedinfinity/go-commoner/optioner"
	"github.com/boundedinfinity/go-commoner/pather"
)

func FileSearch(v o.Option[string], rs ...o.Option[string]) (string, bool) {
	if v.Empty() {
		return "", false
	}

	abs, err := filepath.Abs(v.Get())

	if err == nil {
		ok, err := pather.PathExistsErr(abs)

		if err == nil && ok {
			return abs, true
		}
	}

	if rs != nil {
		for _, r := range rs {
			if r.Empty() {
				continue
			}

			p := filepath.Join(r.Get(), v.Get())
			abs, err := filepath.Abs(p)

			if err != nil {
				continue
			}

			ok, err := pather.PathExistsErr(abs)

			if err == nil && ok {
				return abs, true
			}
		}
	}

	return "", false
}
