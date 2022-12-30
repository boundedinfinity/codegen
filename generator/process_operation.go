package generator

import (
	ct "boundedinfinity/codegen/codegen_type"
	rc "boundedinfinity/codegen/render_context"
	"path"
	"strings"

	"github.com/boundedinfinity/go-commoner/extentioner"
	o "github.com/boundedinfinity/go-commoner/optioner"
)

func (t *Generator) processOperation(po ct.OperationContext, rctx *rc.RenderContextOperation) error {
	var err error
	fileInfo := po.FileInfo
	rootNs := t.projectManager.Merged.Info.Namespace.Get()
	var schemaNs string
	var relNs string

	schemaNs = fileInfo.Source
	schemaNs = strings.Replace(schemaNs, fileInfo.Root, "", 1)
	schemaNs = path.Join(rootNs, schemaNs)
	schemaNs = extentioner.Strip(schemaNs)
	schemaNs = extentioner.Strip(schemaNs)

	relNs = schemaNs
	relNs = strings.ReplaceAll(schemaNs, rootNs, "")
	relNs = strings.Replace(relNs, "/", "", 1)

	input := t.typeManager.Find(po.Input)

	if input.Empty() {
		// TODO
	}

	var inputRc rc.RenderContext

	if err := t.processType(o.None[string](), input.Get(), input.Get().Schema, &inputRc); err != nil {
		return err
	}

	output := t.typeManager.Find(po.Output)

	if output.Empty() {
		// TODO
	}

	var outputRc rc.RenderContext

	if err := t.processType(o.None[string](), output.Get(), output.Get().Schema, &outputRc); err != nil {
		return err
	}

	*rctx = rc.RenderContextOperation{
		RenderContextBase: rc.RenderContextBase{
			Namespace: ct.Namespace{
				RootNs:   rootNs,
				SchemaNs: schemaNs,
				RelNs:    relNs,
			},
			FileInfo:    fileInfo,
			Name:        po.Name.Get(),
			Description: po.Description.Get(),
			// IsPublic:    b.Public.OrElse(true),
			// IsInterface: false,
			// IsRequired:  b.Required.OrElse(false),
		},
		Input:  inputRc,
		Output: outputRc,
	}

	return err
}
