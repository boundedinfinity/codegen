package generator

import (
	ct "boundedinfinity/codegen/codegen_type"
	rc "boundedinfinity/codegen/render_context"
	"path"
	"strings"

	"github.com/boundedinfinity/go-commoner/extentioner"
	o "github.com/boundedinfinity/go-commoner/optioner"
)

func (t *Generator) processOperation(po ct.CodeGenProjectOperation, rctx *rc.RenderContextOperation) error {
	var err error
	fileInfo := po.SourceMeta
	rootNs := t.projectManager.Merged.Info.Namespace.Get()
	var schemaNs string
	var relNs string

	schemaNs = fileInfo.SourcePath.Get()
	schemaNs = strings.Replace(schemaNs, fileInfo.RootPath.Get(), "", 1)
	schemaNs = path.Join(rootNs, schemaNs)
	schemaNs = extentioner.Strip(schemaNs)
	schemaNs = extentioner.Strip(schemaNs)

	relNs = schemaNs
	relNs = strings.ReplaceAll(schemaNs, rootNs, "")
	relNs = strings.Replace(relNs, "/", "", 1)

	input := t.typeManager.Find(po.Input.Get().Base().Id)

	if input.Empty() {
		// TODO
	}

	var inputRc rc.RenderContext

	if err := t.processType(o.None[string](), input.Get(), input.Get(), &inputRc); err != nil {
		return err
	}

	output := t.typeManager.Find(po.Output.Get().Base().Id)

	if output.Empty() {
		// TODO
	}

	var outputRc rc.RenderContext

	if err := t.processType(o.None[string](), output.Get(), output.Get(), &outputRc); err != nil {
		return err
	}

	*rctx = rc.RenderContextOperation{
		RenderContextBase: rc.RenderContextBase{
			RenderNamespace: ct.RenderNamespace{
				RootNs:   rootNs,
				SchemaNs: schemaNs,
				RelNs:    relNs,
			},
			SourceMeta:  fileInfo,
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
