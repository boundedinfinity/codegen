package template_manager

import (
	"boundedinfinity/codegen/canonical"
	"boundedinfinity/codegen/model"
	"boundedinfinity/codegen/template_type"
	"bytes"
	"go/format"

	"github.com/boundedinfinity/go-mimetyper/mime_type"
)

func (t *TemplateManager) RenderModel(schema canonical.Canonical) ([]ModelOutput, error) {
	tmpls := t.FindSchemaTemplate(schema.SchemaType())
	outputs := make([]ModelOutput, 0)

	for _, tmpl := range tmpls {
		if output, err := t.render(tmpl, schema); err != nil {
			return outputs, err
		} else {
			outputs = append(outputs, ModelOutput{
				TemplateOutput: output,
				Schema:         schema,
			})
		}
	}

	return outputs, nil
}

func (t *TemplateManager) RenderOperation(schema model.CodeGenSchemaOperation) ([]TemplateOutput, error) {
	outputs := make([]TemplateOutput, 0)

	for _, tmpl := range t.FindTemplateType(template_type.Operation) {
		if output, err := t.render(tmpl, schema); err != nil {
			return outputs, err
		} else {
			outputs = append(outputs, output)
		}
	}

	return outputs, nil
}

func (t *TemplateManager) RenderNamespace(schema model.CodeGenSchemaOperation) ([]TemplateOutput, error) {
	outputs := make([]TemplateOutput, 0)

	for _, tmpl := range t.FindTemplateType(template_type.Namespace) {
		if output, err := t.render(tmpl, schema); err != nil {
			return outputs, err
		} else {
			outputs = append(outputs, output)
		}
	}

	return outputs, nil
}

func (t *TemplateManager) render(tmpl TemplateContext, data any) (TemplateOutput, error) {
	output := TemplateOutput{
		TemplateContext: tmpl,
	}

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

	output.Output = rendered
	return output, nil
}
