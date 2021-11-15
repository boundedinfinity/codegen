package model

import "github.com/boundedinfinity/jsonschema/mimetype"

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
	CACHE_DIR                           = ".codegen"
)

var (
	SUPPORTED_MIMETYPES = mimetype.Slice(mimetype.ApplicationXYaml, mimetype.ApplicationJson)
)
