package generator

import (
	"boundedinfinity/codegen/lang_ext"
	"boundedinfinity/codegen/util"
	"embed"
	"fmt"
	"io/fs"
	"os"
	"text/template"

	"github.com/boundedinfinity/jsonschema"
	"github.com/boundedinfinity/optioner"
)

//go:embed templates/*
var templateFs embed.FS

type context struct {
	Package string
	Name    string
	Schema  jsonschema.JsonSchmea
}

type Generator struct {
	tmpl         *template.Template
	replacements map[string]string
}

func New(replacements map[string]string) (*Generator, error) {
	g := &Generator{
		tmpl:         template.New(""),
		replacements: replacements,
	}

	if err := fs.WalkDir(templateFs, "templates", makeWalkFn(g.tmpl)); err != nil {
		return nil, err
	}

	return g, nil
}

func makeWalkFn(tmpl *template.Template) fs.WalkDirFunc {
	return func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}

		if d.IsDir() {
			return nil
		}

		fmt.Println(path)

		content, err := fs.ReadFile(templateFs, path)

		if err != nil {
			return err
		}

		if _, err := tmpl.New(path).Parse(string(content)); err != nil {
			return nil
		}

		return nil
	}
}

func (t *Generator) GenerateSchema(lang lang_ext.LanguageExt, name optioner.StringOption, schema jsonschema.JsonSchmea) error {
	tn, err := util.GetTypeName(name, schema, t.replacements)

	if err != nil {
		return err
	}

	pn, err := util.GetPackageName(optioner.NewStringEmpty(), schema, t.replacements)

	if err != nil {
		return err
	}

	ctx := context{
		Package: pn,
		Name:    tn,
		Schema:  schema,
	}

	sp, err := util.GetTemplateSourcePath(lang, schema)

	if err != nil {
		return err
	}

	if err := t.tmpl.ExecuteTemplate(os.Stdout, sp, ctx); err != nil {
		return err
	}

	return nil
}
