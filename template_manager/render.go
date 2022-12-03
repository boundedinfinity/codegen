package template_manager

import (
	"boundedinfinity/codegen/model"
	"boundedinfinity/codegen/template_type"
	"boundedinfinity/codegen/util"
	"bytes"
	"go/format"

	"github.com/boundedinfinity/go-commoner/slicer"
	jmodel "github.com/boundedinfinity/go-jsonschema/model"
	"github.com/boundedinfinity/go-mimetyper/mime_type"
)

func (t *TemplateManager) RenderModel(schema jmodel.JsonSchema) ([]TemplateOutput, error) {
	tmpls := t.Find(template_type.Model)

	found := slicer.Filter(tmpls, func(tmpl TemplateContext) bool {
		return util.IsJsonSchemaTemplate(t.jsonSchemas.GetType(schema), tmpl.Path)
	})

	ctx := RenderContext{
		Info:   t.codeGenSchema.Info,
		Schema: schema,
	}

	return t.render(found, ctx)
}

func (t *TemplateManager) RenderOperation(schema model.CodeGenSchemaOperation) ([]TemplateOutput, error) {
	return t.render(t.Find(template_type.Operation), schema)
}

func (t *TemplateManager) RenderNamespace(schema model.CodeGenSchemaOperation) ([]TemplateOutput, error) {
	return t.render(t.Find(template_type.Namespace), schema)
}

func (t *TemplateManager) render(tmpls []TemplateContext, data any) ([]TemplateOutput, error) {
	output := make([]TemplateOutput, 0)

	for _, tmpl := range tmpls {
		var writer bytes.Buffer

		if err := tmpl.Template.Execute(&writer, data); err != nil {
			return output, err
		}

		rendered := writer.Bytes()

		if t.formatSource {
			switch tmpl.OutputMimeType {
			case mime_type.ApplicationXGo:
				formatted, err := format.Source([]byte(rendered))

				if err != nil {
					return output, err
				}

				rendered = formatted
			}
		}

		output = append(output, TemplateOutput{
			MimeType:     tmpl.TemplateMimeType,
			TemplateType: tmpl.TemplateType,
			Output:       writer.Bytes(),
			Path:         tmpl.Path,
		})
	}

	return output, nil
}
