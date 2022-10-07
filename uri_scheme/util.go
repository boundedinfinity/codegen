package uri_scheme

import "github.com/boundedinfinity/go-commoner/stringer"

func DetectErr(uri string) (Scheme, error) {
	for _, x := range All {
		if stringer.HasPrefix(uri, x) {
			return x, nil
		}
	}

	var scheme Scheme
	return scheme, ErrUriTypeInvalidV(uri)
}

func Detect(uri string) (Scheme, bool) {
	scheme, err := DetectErr(uri)

	if err != nil {
		return scheme, false
	}

	return scheme, true
}

func PathErr(uri string) (string, Scheme, error) {
	var path string
	var scheme Scheme

	scheme, ok := Detect(uri)

	if !ok {
		return path, scheme, ErrUriTypeInvalidV(uri)
	}

	switch scheme {
	case File:
		path = stringer.Remove(uri, File)
	case Http, Https:
		path = uri
	default:
		return path, scheme, ErrUriTypeInvalidV(uri)
	}

	return path, scheme, nil
}

func Path(uri string) (string, Scheme, bool) {
	path, scheme, err := PathErr(uri)

	if err != nil {
		return path, scheme, false
	}

	return path, scheme, true
}
