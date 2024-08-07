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
	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
)

func New(lang string) (*Generator, error) {
	gen := &Generator{
		lang:                lang,
		caserConversion:     "kebab-to-pascal",
		templateDescriptors: []templateDescriptor{},
		formatSource:        true,
		dumpTemplates:       true,
	}

	glob := fmt.Sprintf("templates/*.%s.tpl", lang)
	templ, err := template.New("").Funcs(getHelpers(lang)).
		ParseFS(embeddedTemplates, glob)

	if err != nil {
		return nil, err
	}

	gen.templ = templ
	// gen._dumpTemplates()

	return gen, nil
}

type Generator struct {
	templ               *template.Template
	lang                string
	caserConversion     string
	formatSource        bool
	dumpTemplates       bool
	templateDescriptors []templateDescriptor
}

func (this *Generator) _dumpTemplates() {
	message := this.templ.DefinedTemplates()
	message = stringer.Replace(message, "", "; defined templates are: ")
	message = stringer.Replace(message, "\n", ",")

	fmt.Println("====================================================================")
	fmt.Println(message)
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

	var buffer bytes.Buffer
	var err error

	switch i := typ.(type) {
	case *model.CodeGenString:
		err = this.templ.ExecuteTemplate(&buffer, "string_type", typ)
	case *model.CodeGenObject:
		err = this.templ.ExecuteTemplate(&buffer, "object_type", typ)
	case *model.CodeGenInteger:
		err = this.templ.ExecuteTemplate(&buffer, "integer_type", typ)
	case *model.CodeGenFloat:
		err = this.templ.ExecuteTemplate(&buffer, "float_type", typ)
	// case *model.CodeGenArray:
	// 	err = this.templ.ExecuteTemplate(&buffer, "array_type", typ)
	default:
		fmt.Printf("unsupported type %v", i.Schema())
	}

	if err != nil {
		return nil, err
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

		var dir string
		dir = typ.Common().Id.Get()
		dir = pather.Paths.Dir(dir)

		var filename string
		filename = typ.Common().Id.Get()
		filename = pather.Paths.Base(filename)
		filename = extentioner.Join(filename, ".go")

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
