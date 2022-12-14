package template_manager_test

import (
	"boundedinfinity/codegen/cacher"
	"boundedinfinity/codegen/model"
	"boundedinfinity/codegen/template_manager"
	"log"
	"testing"

	o "github.com/boundedinfinity/go-commoner/optioner"
)

var (
	tm *template_manager.TemplateManager
)

func loadTemplates() {
	p1 := "file://../../codegen-templates/templates/go/types/base/date.model.go.gotmpl"

	var err error

	c, err := cacher.New()

	if err != nil {
		log.Fatal(err)
	}

	if err := c.Cache("test", p1); err != nil {
		log.Fatal(err)
	}

	tm, err = template_manager.New(template_manager.Cacher(c))

	if err != nil {
		log.Fatal(err)
	}

	tm.Register(model.CodeGenSchemaTemplates{
		Files: []model.CodeGenSchemaTemplateFile{
			{Path: o.Some(p1)},
		},
	})
}

// func renderTemplate(data canonical.Canonical) error {
// 	var name string

// 	switch data.(type) {
// 	case canonical.CanonicalInteger:
// 		name = "integer.model.gotmpl"
// 	case canonical.CanonicalString:
// 		name = "string.model.gotmpl"
// 	case canonical.CanonicalObject:
// 		name = "object.model.gotmpl"
// 	}

// 	path := filepath.Join("/tmp/codegen", strings.ReplaceAll(name, ".gotmpl", ".go"))

// 	// if err := os.Remove(path); err != nil {
// 	// 	return err
// 	// }

// 	file, err := os.Create(path)

// 	if err != nil {
// 		return err
// 	}

// 	defer file.Close()

// 	if err := tmpl.ExecuteTemplate(file, name, data); err != nil {
// 		return err
// 	}

// 	return nil
// }

func Test_String_Marshal(t *testing.T) {
	loadTemplates()

	// v1 := canonical.CanonicalInteger{
	// 	CanonicalBase: canonical.CanonicalBase{
	// 		Source: "../../codegen-templates/templates/go/types/base/integer.model.go.gotmpl",
	// 		// Import:      "boundedinfinity/example/canonical/SomeInteger",
	// 		Id: o.Some("boundedinfinity/example/canonical/SomeInteger"),
	// 		// Package:     "canonical",
	// 		Name:        o.Some("Si"),
	// 		Description: o.Some("A custom integer type"),
	// 	},
	// 	Maximum:    o.Some(10),
	// 	Minimum:    o.Some(1),
	// 	MultipleOf: o.Some(2),
	// }

	// v3 := canonical.CanonicalString{
	// 	CanonicalBase: canonical.CanonicalBase{
	// 		Source:      "../../codegen-templates/templates/go/types/base/string.model.go.gotmpl",
	// 		Import:      "boundedinfinity/example/canonical/SomeString",
	// 		Type:        "SomeString",
	// 		Package:     "canonical",
	// 		Name:        "Ss",
	// 		Description: optioner.Some("A custom string type"),
	// 	},
	// 	Maximum: optioner.Some(10),
	// 	Minimum: optioner.Some(1),
	// }

	// v2 := canonical.CanonicalObject{
	// 	CanonicalBase: canonical.CanonicalBase{
	// 		Source:      "../../codegen-templates/templates/go/types/base/object.model.go.gotmpl",
	// 		Import:      "boundedinfinity/example/canonical/SomeObject",
	// 		Type:        "SomeObject",
	// 		Package:     "canonical",
	// 		Description: optioner.Some("A custom object type"),
	// 	},
	// 	Properties: []canonical.Canonical{
	// 		v1, v3,
	// 	},
	// }

	// _, err := tm.RenderModel(v1)
	// assert.Nil(t, err)
	// // assert.Nil(t, renderTemplate(v3))
	// // assert.Nil(t, renderTemplate(v2))

	// expected := ""
	// actual := ""
	// assert.Equal(t, expected, string(actual))
}
