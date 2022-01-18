package model

import (
	"path/filepath"

	"github.com/boundedinfinity/mimetyper/mime_type"
)

const (
	TYPE_UNKNOWN                        = "<UNKNOWN_TYPE>"
	TYPE_BUILTIN                        = "<builtin>"
	NAMESPACE_SEP                       = "/"
	COLLECTION_SUFFIX                   = "[]"
	SUMMERY_SIZE                        = 35
	DEFAULT_DESCRIPTION_SPLIT_CHARACTER = "\n"
	DEFAULT_FILENAME_MARKER             = "gen"
	DEFAULT_FILENAME_DISABLE            = "none"
	JSON_DEFAULT_STRING                 = "a string"
	JSON_DEFAULT_NUMBER                 = 1
	WORK_DIR                            = ".codegen"
)

var (
	CACHE_DIR           = filepath.Join(WORK_DIR, "cache")
	BUILD_DIR           = filepath.Join(WORK_DIR, "build")
	SUPPORTED_MIMETYPES = mime_type.Slice(mime_type.ApplicationXYaml, mime_type.ApplicationJson)
)
