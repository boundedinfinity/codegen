package uripath

import (
	"boundedinfinity/codegen/model"
	"net/url"
	"os"
	"path"
)

func Join(u *url.URL, elem ...string) (*url.URL, error) {
	v := *u
	v.Path = path.Join(elem...)

	return &v, nil
}

func Exists(u *url.URL) (bool, error) {
	switch u.Scheme {
	case "file":
	default:
		return false, model.ErrUnsupportedSchemeV(u.Scheme)
	}

	_, err := os.Stat(u.Path)

	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		} else {
			return false, err
		}
	}

	return true, nil
}
