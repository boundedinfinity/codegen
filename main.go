package main

import (
	"boundedinfinity/codegen/system"
	"fmt"
	"os"
)

//go:generate enumer -standalone=true -package=uritype -name=UriType -items=file,http,https
//go:generate enumer -standalone=true -package=fileext -name=FileExt -items=yaml,yml,json
//go:generate enumer -standalone=true -package=template_type -name=TemplateType -items=model,operation,namespace

// go:generate enumeration -package=model -name=LanguageExt -suffix=Enum -items=unknown,go,mod,ts,js,html,css
// go:generate enumeration -package=model -name=TemplateExt -suffix=Enum -items=unknown,gotmpl,handlebars

func main() {
	schemaPaths := os.Args[1:]

	l := system.New()

	if err := l.Load(schemaPaths...); err != nil {
		handleError(err)
	}

	if err := l.Unmarshal(); err != nil {
		handleError(err)
	}

	if err := l.Process(); err != nil {
		handleError(err)
	}

	// if err := g.Generate(); err != nil {
	// 	os.Exit(handleError(err))
	// }
}

func handleError(err error) int {
	fmt.Printf("%v\n", err.Error())
	return 1
}
