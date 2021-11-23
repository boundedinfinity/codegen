package main

import (
	"boundedinfinity/codegen/system"
	"fmt"
	"os"
)

//go:generate enumer -standalone=true -package=uritype -name=UriType -items=file,http,https
//go:generate enumer -standalone=true -package=schema_ext -name=SchemaExt -items=yaml,yml,json
//go:generate enumer -standalone=true -package=template_type -name=TemplateType -items=model,operation,namespace
//-go:generate enumer -standalone=true -package=lang_ext -name=LanguageExt -items=go,mod,ts,js,html,css
//-go:generate enumer -standalone=true -package=template_ext -name=TemplateExt -items=gotmpl,handlebars

func main() {
	schemaPaths := os.Args[1:]

	s := system.New()

	if err := s.Process(schemaPaths...); err != nil {
		os.Exit(handleError(err))
	}

	if err := s.Generate(); err != nil {
		os.Exit(handleError(err))
	}
}

func handleError(err error) int {
	fmt.Printf("%v\n", err.Error())
	return 1
}
