package template_manager

import (
	"boundedinfinity/codegen/canonical"
	"boundedinfinity/codegen/model"
	"boundedinfinity/codegen/template_type"
	"bytes"
	"go/format"

	"github.com/boundedinfinity/go-mimetyper/mime_type"
)

func (t *TemplateManager) RenderModel(schema canonical.Canonical) ([]TemplateOutput, error) {
	tmpls := t.Find(template_type.Model)

	ctx := RenderContext{
		Info:   t.codeGenSchema.Info,
		Schema: schema,
	}

	return t.render(tmpls, ctx)
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

		if err := t.combinedTemplates.ExecuteTemplate(&writer, tmpl.Path, data); err != nil {
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
