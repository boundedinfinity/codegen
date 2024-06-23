package generator

// https://www.digitalocean.com/community/tutorials/how-to-use-templates-in-go
// https://gobyexample.com/embed-directive

import (
	"boundedinfinity/codegen/model"
	"embed"
	"fmt"
	"text/template"
)

//go:embed templates/*.go.tpl
var embeddedTemplates embed.FS

func New(lang string) (*Generator, error) {
	gen := &Generator{lang: lang}

	tmpl, err := template.ParseFS(embeddedTemplates)
	if err != nil {
		return nil, err
	}

	tmpl.Funcs(gen.getHelpers())
	gen.templ = tmpl

	return gen, nil
}

type Generator struct {
	templ *template.Template
	lang  string
}

func (t *Generator) Generate(project *model.CodeGenProject) (map[string]string, error) {
	results := map[string]string{}

	if project == nil {
		return results, nil
	}

	for _, typ := range project.Types {
		result, err := t.generateType(typ)

		if err != nil {
			return results, err
		}

		results[typ.CodeGenId()] = result
	}

	return results, nil
}

func (t *Generator) generateType(typ model.CodeGenType) (string, error) {
	var result string

	switch i := typ.(type) {
	case *model.CodeGenString:

	default:
		fmt.Println(i)
	}

	return result, nil
}
