package renderer

import (
	"boundedinfinity/codegen/codegen_type"
	"boundedinfinity/codegen/render_context"
	"bytes"
	"go/format"

	"github.com/boundedinfinity/go-mimetyper/mime_type"
)

func (t *Renderer) RenderModel(schema render_context.RenderContext) ([]ModelOutput, error) {
	outputs := make([]ModelOutput, 0)
	tmpls := t.templateManager.Find(schema.Base().SchemaType)

	if tmpls.Empty() {
		return outputs, nil
	}

	for _, tmpl := range tmpls.Get() {
		if output, err := t.render(*tmpl, schema); err != nil {
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

func (t *Renderer) RenderOperation(schema render_context.RenderContextOperation) ([]TemplateOutput, error) {
	outputs := make([]TemplateOutput, 0)

	// for _, tmpl := range t.FindTemplateType(template_type.Operation) {
	// 	if output, err := t.render(tmpl, schema); err != nil {
	// 		return outputs, err
	// 	} else {
	// 		outputs = append(outputs, output)
	// 	}
	// }

	return outputs, nil
}

func (t *Renderer) RenderNamespace(schema codegen_type.CodeGenProjectOperation) ([]TemplateOutput, error) {
	outputs := make([]TemplateOutput, 0)

	// for _, tmpl := range t.FindTemplateType(template_type.Namespace) {
	// 	if output, err := t.render(tmpl, schema); err != nil {
	// 		return outputs, err
	// 	} else {
	// 		outputs = append(outputs, output)
	// 	}
	// }

	return outputs, nil
}

func (t *Renderer) render(tmpl codegen_type.TemplateMeta, data any) (TemplateOutput, error) {
	output := TemplateOutput{
		TemplateMeta: tmpl,
	}

	var writer bytes.Buffer

	if err := tmpl.Template.Execute(&writer, data); err != nil {
		return output, err
	}

	rendered := writer.Bytes()

	if t.projectManager.Merged.Info.FormatSource.Defined() && t.projectManager.Merged.Info.FormatSource.Get() {
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
