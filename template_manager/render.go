package template_manager

import (
	"boundedinfinity/codegen/canonical/canonical_type"
	"boundedinfinity/codegen/model"
	"boundedinfinity/codegen/render_context"
	"boundedinfinity/codegen/template_type"
	"boundedinfinity/codegen/util"
	"bytes"
	"go/format"

	"github.com/boundedinfinity/go-mimetyper/mime_type"
)

func (t *TemplateManager) RenderModel(schema render_context.RenderContext) ([]ModelOutput, error) {
	tmpls := t.FindSchemaTemplate(canonical_type.CanonicalType(schema.Base().SchemaType))
	outputs := make([]ModelOutput, 0)

	for _, tmpl := range tmpls {
		outputPath := util.DestPath(t.codeGenSchema.Info, schema, tmpl.Path)

		err := render_context.WalkBase(schema, func(base *render_context.RenderContextBase) error {
			base.OutputPath = outputPath
			base.CurrNs = util.CurrentNs(t.codeGenSchema.Info, outputPath)
			base.SourceUri = schema.Base().SourceUri
			return nil
		})

		if err != nil {
			return outputs, err
		}

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

	if err := tmpl.Template.Execute(&writer, data); err != nil {
		return output, err
	}

	// if err := t.combinedTemplates.ExecuteTemplate(&writer, tmpl.Path, data); err != nil {
	// 	return output, err
	// }

	rendered := writer.Bytes()

	if t.codeGenSchema.Info.FormatSource.Defined() && t.codeGenSchema.Info.FormatSource.Get() {
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
