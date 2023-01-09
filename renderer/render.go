package renderer

import (
	ct "boundedinfinity/codegen/codegen_type"
	"boundedinfinity/codegen/codegen_type/template_type"
	"bytes"
	"fmt"
	"go/format"

	o "github.com/boundedinfinity/go-commoner/optioner"
	"github.com/boundedinfinity/go-mimetyper/mime_type"
)

func (t *Renderer) getTemplates(tt template_type.TemplateType, s o.Option[ct.CodeGenType]) []*ct.CodeGenProjectTemplateFile {
	var found []*ct.CodeGenProjectTemplateFile
	var group []*ct.CodeGenProjectTemplateFile

	switch tt {
	case template_type.Model:
		group = t.projectManager.Merged.Templates.Types
	case template_type.Operation:
		group = t.projectManager.Merged.Templates.Operations
	default:
		fmt.Printf("template type %v not implemented\n", tt)
	}

	for _, file := range group {
		if s.Defined() {
			if file.Type.Defined() && s.Get().SchemaType() == file.Type.Get() {
				found = append(found, file)
			}
		} else {
			found = append(found, file)
		}
	}

	return found
}

func (t *Renderer) RenderModel(schema ct.CodeGenType) ([]ModelOutput, error) {
	outputs := make([]ModelOutput, 0)
	tmpls := t.getTemplates(template_type.Model, o.Some(schema))

	for _, tmpl := range tmpls {
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

func (t *Renderer) RenderOperation(schema ct.CodeGenProjectTemplateFile) ([]TemplateOutput, error) {
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

func (t *Renderer) RenderNamespace(schema ct.CodeGenProjectOperation) ([]TemplateOutput, error) {
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

func (t *Renderer) render(tmpl ct.CodeGenProjectTemplateFile, data any) (TemplateOutput, error) {
	output := TemplateOutput{
		CodeGenProjectTemplateFile: tmpl,
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
