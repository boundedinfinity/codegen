package template_manager

import (
	"boundedinfinity/codegen/codegen_project"
	"boundedinfinity/codegen/codegen_type/codegen_type_id"
	"boundedinfinity/codegen/render_context"
	"boundedinfinity/codegen/template_type"
	"bytes"
	"go/format"

	"github.com/boundedinfinity/go-mimetyper/mime_type"
)

func (t *TemplateManager) RenderModel(schema render_context.RenderContext) ([]ModelOutput, error) {
	tmpls := t.FindSchemaTemplate(codegen_type_id.CodgenTypeId(schema.Base().SchemaType))
	outputs := make([]ModelOutput, 0)

	for _, tmpl := range tmpls {
		outputPath := codegen_project.DestPath(t.projectManager.Merged.Info, schema, tmpl.Source)

		err := render_context.WalkBase(schema, func(base *render_context.RenderContextBase) error {
			base.OutputPath = outputPath
			base.CurrNs = codegen_project.CurrentNs(t.projectManager.Merged.Info, outputPath)
			base.Source = schema.Base().Source
			base.Root = schema.Base().Root
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

func (t *TemplateManager) RenderOperation(schema render_context.RenderContextOperation) ([]TemplateOutput, error) {
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

func (t *TemplateManager) RenderNamespace(schema codegen_project.CodeGenProjectOperation) ([]TemplateOutput, error) {
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
