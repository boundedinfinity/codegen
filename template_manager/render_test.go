package template_manager_test

import (
	"boundedinfinity/codegen/conical"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"text/template"

	"github.com/boundedinfinity/go-commoner/optioner"
	"github.com/stretchr/testify/assert"
)

var (
	tmpl *template.Template
)

func loadTemplates() {
	xtmpl, err := template.ParseFiles(
		"../../codegen-templates/templates/go/types/object.model.gotmpl",
		"../../codegen-templates/templates/go/types/string.model.gotmpl",
		"../../codegen-templates/templates/go/types/integer.model.gotmpl",
	)

	if err != nil {
		log.Fatal(err)
	}

	tmpl = xtmpl

}

func renderTemplate(data conical.Conical) error {
	var name string

	switch data.(type) {
	case conical.ConicalInteger:
		name = "integer.model.gotmpl"
	case conical.ConicalString:
		name = "string.model.gotmpl"
	case conical.ConicalObject:
		name = "object.model.gotmpl"
	}

	path := filepath.Join("/tmp/codegen", strings.ReplaceAll(name, ".gotmpl", ".go"))

	// if err := os.Remove(path); err != nil {
	// 	return err
	// }

	file, err := os.Create(path)

	if err != nil {
		return err
	}

	defer file.Close()

	if err := tmpl.ExecuteTemplate(file, name, data); err != nil {
		return err
	}

	return nil
}

func Test_String_Marshal(t *testing.T) {
	loadTemplates()

	v1 := conical.ConicalInteger{
		ConicalBase: conical.ConicalBase{
			Source:      "../../codegen-templates/templates/go/types/base/integer.model.go.gotmpl",
			Import:      "boundedinfinity/example/conical/SomeInteger",
			Type:        "SomeInteger",
			Package:     "conical",
			Name:        "Si",
			Description: optioner.Some("A custom integer type"),
		},
		Maximum:    optioner.Some(10),
		Minimum:    optioner.Some(1),
		MultipleOf: optioner.Some(2),
	}

	v3 := conical.ConicalString{
		ConicalBase: conical.ConicalBase{
			Source:      "../../codegen-templates/templates/go/types/base/string.model.go.gotmpl",
			Import:      "boundedinfinity/example/conical/SomeString",
			Type:        "SomeString",
			Package:     "conical",
			Name:        "Ss",
			Description: optioner.Some("A custom string type"),
		},
		Maximum: optioner.Some(10),
		Minimum: optioner.Some(1),
	}

	v2 := conical.ConicalObject{
		ConicalBase: conical.ConicalBase{
			Source:      "../../codegen-templates/templates/go/types/base/object.model.go.gotmpl",
			Import:      "boundedinfinity/example/conical/SomeObject",
			Type:        "SomeObject",
			Package:     "conical",
			Description: optioner.Some("A custom object type"),
		},
		Properties: []conical.Conical{
			v1, v3,
		},
	}

	assert.Nil(t, renderTemplate(v1))
	assert.Nil(t, renderTemplate(v3))
	assert.Nil(t, renderTemplate(v2))

	expected := ""
	actual := ""
	assert.Equal(t, expected, string(actual))
}
