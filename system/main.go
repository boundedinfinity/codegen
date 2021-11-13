package system

import "boundedinfinity/codegen/uritype"

type System struct {
	uri2uriType   map[string]uritype.UriType
	local2uriType map[string]uritype.UriType
	mimeType      map[string]string
}

func New() *System {
	return &System{
		uri2uriType:   make(map[string]uritype.UriType),
		local2uriType: make(map[string]uritype.UriType),
		mimeType:      make(map[string]string),
	}
}
