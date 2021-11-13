package system

import (
	"boundedinfinity/codegen/uritype"
	"net/url"
)

func (t *System) Load(uris ...string) error {
	for _, uri := range uris {
		var typ uritype.UriType

		if err := t.detectUriType(uri, &typ); err != nil {
			return err
		}

		t.uri2uriType[uri] = typ
	}

	for uri, typ := range t.uri2uriType {
		switch typ {
		case uritype.File:
			t.local2uriType[uri] = typ
		case uritype.Http, uritype.Https:

		}
	}

	return nil
}

func (t *System) detectUriType(uri string, typ *uritype.UriType) error {
	parsed, err := url.Parse(uri)

	if err != nil {
		return err
	}

	x, err := uritype.Parse(parsed.Scheme)

	if err != nil {
		return err
	}

	*typ = x

	return nil
}

// func (t *System) loadUri(uri string, ut uritype.UriType, bs *[]byte) error {
// 	switch ut {
// 	case uritype.File:
// 		if err := t.readFile(uri, bs); err != nil {
// 			return err
// 		}
// 	case uritype.Http:
// 		if err := t.readHttp(uri, bs); err != nil {
// 			return err
// 		}
// 	default:
// 		return ErrUriTypeUnsupportedV(uri)
// 	}

// 	return nil
// }
