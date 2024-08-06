package generator

import (
	"embed"

	"github.com/boundedinfinity/go-commoner/idiomatic/pather"
	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
)

//go:embed templates/*
var embeddedTemplates embed.FS

func (t *Generator) loadTemplates(params templateDescriptor) error {
	tnames, err := getTemplateDescriptors()
	if err != nil {
		return err
	}

	t.templ.Funcs(getHelpers(t.lang))

	for _, tname := range tnames {
		if t.templateAlreadyLoaded(tname) || !includeTemplate(params, tname) {
			continue
		}

		_, err := t.templ.ParseFS(embeddedTemplates, tname.path)
		if err != nil {
			return err
		}

		t.templateDescriptors = append(t.templateDescriptors, tname)
	}

	return nil
}

func (t *Generator) loadedTemplates() []string {
	var names []string

	for _, templ := range t.templ.Templates() {
		names = append(names, templ.Name())
	}

	return names
}

type templateDescriptor struct {
	name     string
	path     string
	baseType string
	function string
	lang     string
}

func getTemplateDescriptors() ([]templateDescriptor, error) {
	var results []templateDescriptor

	entries, err := embeddedTemplates.ReadDir("templates")
	if err != nil {
		return results, err
	}

	for _, entry := range entries {
		comps := stringer.Split(entry.Name(), ".")

		results = append(results,
			templateDescriptor{
				path:     pather.Paths.Join("templates", entry.Name()),
				name:     entry.Name(),
				baseType: comps[0],
				function: comps[1],
				lang:     comps[2],
			})
	}

	return results, nil
}

func (t *Generator) templateAlreadyLoaded(tname templateDescriptor) bool {
	var ok bool

	for _, name := range t.loadedTemplates() {
		if name == tname.name {
			ok = true
			continue
		}
	}

	return ok
}

func includeTemplate(params, name templateDescriptor) bool {
	ok := true

	if params.baseType != name.baseType {
		ok = false
	}

	if params.lang != "" && params.lang != name.lang {
		ok = false
	}

	if params.function != "" && params.function != name.function {
		ok = false
	}

	return ok
}
