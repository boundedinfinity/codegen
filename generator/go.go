package generator

// https://www.digitalocean.com/community/tutorials/how-to-use-templates-in-go
// https://gobyexample.com/embed-directive

//lint:file-ignore ST1006
// https://staticcheck.dev/docs/checks#ST1006

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

func (this *Generator) CasserConvertion(v string) *Generator {
	return model.SetV(this, &this.caserConversion, v)
}

func (this *Generator) GenerateProject(project *model.CodeGenProject) (map[string]string, error) {
	results := map[string]string{}

	if project == nil {
		return results, nil
	}

	for _, typ := range project.Types {
		typeResults, err := this.GenerateType(typ)
		if err != nil {
			return results, err
		}

		mapper.MergeInto(results, typeResults)
	}

	return results, nil
}

func (this *Generator) GenerateType(typ model.CodeGenSchema) (map[string]string, error) {
	results := map[string]string{}
	params := templateDescriptor{lang: "go", baseType: typ.Schema()}

	if err := this.loadTemplates(params); err != nil {
		return results, err
	}

	var templateDescriptors []templateDescriptor

	switch i := typ.(type) {
	case *model.CodeGenString, *model.CodeGenInteger, *model.CodeGenFloat, *model.CodeGenObject:
		templateDescriptors = slicer.Filter(
			func(_ int, td templateDescriptor) bool { return includeTemplate(params, td) },
			this.templateDescriptors...,
		)
	default:
		fmt.Printf("unsupported type %v", i.Schema())
	}

	for _, td := range templateDescriptors {
		var buffer bytes.Buffer

		err := this.templ.ExecuteTemplate(&buffer, td.name, typ)
		if err != nil {
			return results, err
		}

		if this.formatSource {
			formatted, err := format.Source(buffer.Bytes())
			if err != nil {
				content := buffer.String()
				buffer.Reset()
				buffer.WriteString("// FORMAT ERROR: " + err.Error() + "\n\n")
				buffer.WriteString(content)
			} else {
				buffer.Reset()
				buffer.Write(formatted)
			}
		}

		var dir string
		dir = typ.Common().Id.Get()
		dir = pather.Paths.Dir(dir)

		var filename string
		filename = td.path
		filename = pather.Paths.Base(filename)
		filename = stringer.Replace(filename, typ.Common().Id.Get(), typ.Schema())
		filename = extentioner.Strip(filename)

		path := pather.Paths.Join(dir, filename)
		results[path] = buffer.String()
	}

	return results, nil
}

func (this *Generator) WriteType(typ model.CodeGenSchema) (map[string]string, error) {
	results, err := this.GenerateType(typ)
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
