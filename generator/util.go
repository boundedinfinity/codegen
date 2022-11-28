package generator

import (
	"io/fs"
	"text/template"

	o "github.com/boundedinfinity/go-commoner/optioner"
	"github.com/boundedinfinity/go-jsonschema/model"
)

func makeWalkFn(tmpl *template.Template) fs.WalkDirFunc {
	return func(path string, d fs.DirEntry, err error) error {
		// if err != nil {
		// 	return nil
		// }

		// if d.IsDir() {
		// 	return nil
		// }

		// fmt.Println(path)

		// content, err := fs.ReadFile(templates.TemplateFs, path)

		// if err != nil {
		// 	return err
		// }

		// if _, err := tmpl.New(path).Parse(string(content)); err != nil {
		// 	return nil
		// }

		return nil
	}
}

func (t *Generator) GenerateSchema(name o.Option[string], schema model.JsonSchema) error {
	// tn, err := util.GetTypeName(name, schema, t.replacements)

	// if err != nil {
	// 	return err
	// }

	// pn, err := util.GetPackageName(o.None[string](), schema, t.replacements)

	// if err != nil {
	// 	return err
	// }

	// ctx := context{
	// 	Package: pn,
	// 	Name:    tn,
	// 	Schema:  schema,
	// }

	// sp, err := util.GetTemplateSourcePath(lang, schema)

	// if err != nil {
	// 	return err
	// }

	// if err := t.tmpl.ExecuteTemplate(os.Stdout, sp, ctx); err != nil {
	// 	return err
	// }

	return nil
}
