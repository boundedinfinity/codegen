package generator

// https://www.digitalocean.com/community/tutorials/how-to-use-templates-in-go
// https://gobyexample.com/embed-directive

import (
	"boundedinfinity/codegen/model"
	"embed"
	"fmt"
	"text/template"

	"github.com/boundedinfinity/go-commoner/idiomatic/pather"
)

//go:embed templates/*.go.tpl
var embeddedTemplates embed.FS

func New(lang string) (*Generator, error) {
	gen := &Generator{
		lang:  lang,
		templ: template.New(""),
	}

	if err := gen.loadTemplates(); err != nil {
		return gen, err
	}

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

func (t *Generator) templateNames() []string {
	var names []string

	for _, templ := range t.templ.Templates() {
		names = append(names, templ.Name())
	}

	return names
}

func (t *Generator) loadTemplates() error {
	t.templ.Funcs(t.getHelpers())

	entries, err := embeddedTemplates.ReadDir("templates")
	if err != nil {
		return err
	}

	for _, entry := range entries {
		path := pather.Paths.Join("templates", entry.Name())

		_, err := t.templ.ParseFS(embeddedTemplates, path)
		if err != nil {
			return err
		}
	}
	return nil
}
