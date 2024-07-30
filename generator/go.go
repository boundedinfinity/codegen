package generator

// https://www.digitalocean.com/community/tutorials/how-to-use-templates-in-go
// https://gobyexample.com/embed-directive

import (
	"boundedinfinity/codegen/model"
	"bytes"
	"fmt"
	"go/format"
	"os"
	"text/template"

	"github.com/boundedinfinity/go-commoner/idiomatic/extentioner"
	"github.com/boundedinfinity/go-commoner/idiomatic/mapper"
	"github.com/boundedinfinity/go-commoner/idiomatic/pather"
	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
)

func New(lang string) (*Generator, error) {
	gen := &Generator{
		lang:                lang,
		caserConversion:     "kebab-to-pascal",
		templ:               template.New(""),
		templateDescriptors: []templateDescriptor{},
		formatSource:        true,
	}

	return gen, nil
}

type Generator struct {
	templ               *template.Template
	lang                string
	caserConversion     string
	formatSource        bool
	templateDescriptors []templateDescriptor
}

func (t *Generator) CasserConvertion(v string) *Generator {
	return model.SetV(t, &t.caserConversion, v)
}

func (t *Generator) GenerateProject(project *model.CodeGenProject) (map[string]string, error) {
	results := map[string]string{}

	if project == nil {
		return results, nil
	}

	for _, typ := range project.Types {
		typeResults, err := t.GenerateType(typ)
		if err != nil {
			return results, err
		}

		mapper.MergeInto(results, typeResults)
	}

	return results, nil
}

func (t *Generator) GenerateType(typ model.CodeGenType) (map[string]string, error) {
	results := map[string]string{}
	params := templateDescriptor{lang: "go", baseType: typ.GetType()}

	if err := t.loadTemplates(params); err != nil {
		return results, err
	}

	var templateDescriptors []templateDescriptor

	switch i := typ.(type) {
	case *model.CodeGenString, *model.CodeGenInteger, *model.CodeGenFloat:
		templateDescriptors = slicer.Filter(
			func(_ int, td templateDescriptor) bool { return includeTemplate(params, td) },
			t.templateDescriptors...,
		)
	default:
		fmt.Printf("unsupported %v", i.GetType())
	}

	for _, td := range templateDescriptors {
		var buffer bytes.Buffer

		err := t.templ.ExecuteTemplate(&buffer, td.name, typ)
		if err != nil {
			return results, err
		}

		if t.formatSource {
			formatted, err := format.Source(buffer.Bytes())
			if err != nil {
				return results, err
			}

			buffer.Reset()
			_, err = buffer.Write(formatted)
			if err != nil {
				return results, err
			}
		}

		var dir string
		dir = typ.Common().Id.Get()
		dir = pather.Paths.Dir(dir)

		var filename string
		filename = td.path
		filename = pather.Paths.Base(filename)
		filename = stringer.Replace(filename, typ.Common().Id.Get(), typ.GetType())
		filename = extentioner.Strip(filename)

		path := pather.Paths.Join(dir, filename)
		results[path] = buffer.String()
	}

	return results, nil
}

func (t *Generator) WriteType(typ model.CodeGenType) (map[string]string, error) {
	results, err := t.GenerateType(typ)
	if err != nil {
		return results, err
	}

	for path, content := range results {
		dir := pather.Paths.Dir(path)

		if _, err = pather.Dirs.EnsureErr(dir); err != nil {
			return results, err
		}

		if err = os.WriteFile(path, []byte(content), os.FileMode(0755)); err != nil {
			return results, err
		}
	}

	return results, nil
}
