package util

import (
	"path"

	"github.com/boundedinfinity/go-commoner/caser"
	"github.com/boundedinfinity/go-commoner/chain"
	"github.com/boundedinfinity/go-commoner/stringer"
	"github.com/boundedinfinity/jsonschema/model"
)

func GoTypeBase(s model.JsonSchemaCommon) string {
	result := chain.New[string]().
		Append(path.Base).
		Append(replaceNonWord).
		Append(caser.PhraseToPascal[string]).
		RunSingle(s.Id.Get())
	return result
}

func GoTypeDir(s model.JsonSchemaCommon) string {
	result := chain.New[string]().
		Append(removeScheme).
		Append(path.Dir).
		RunSingle(s.Id.Get())
	return result
}

func replaceNonWord(s string) string {
	return stringer.ReplaceNonWord(s, " ")
}

func removeScheme(s string) string {
	return chain.New[string]().
		Append(chain.StringRemover[string]("http://")).
		Append(chain.StringRemover[string]("https://")).
		Append(chain.StringRemover[string]("file://")).
		RunSingle(s)
}
