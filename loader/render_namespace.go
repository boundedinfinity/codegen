package loader

import (
	ct "boundedinfinity/codegen/codegen_type"
	"path"
	"strings"

	"github.com/boundedinfinity/go-commoner/extentioner"
)

func (t *Loader) ProcessNamespace() error {
	err := ct.WalkType(func(_ *ct.CodeGenProject, typ ct.CodeGenType) error {
		*typ.Namespace() = t.typeNamespace(typ)
		return nil
	}, t.projectManager.Merged)

	if err != nil {
		return err
	}

	err = ct.WalkOperation(func(_ *ct.CodeGenProject, operation *ct.CodeGenProjectOperation) error {
		if operation.Input.Defined() {
			*operation.Input.Get().Namespace() = t.typeNamespace(operation.Input.Get())
		}

		if operation.Output.Defined() {
			*operation.Output.Get().Namespace() = t.typeNamespace(operation.Output.Get())
		}

		return nil
	}, t.projectManager.Merged)

	return nil
}

func (t *Loader) namespace(meta ct.SourceMeta) ct.RenderNamespace {
	var ns ct.RenderNamespace

	ns.RootNs = t.projectManager.Merged.Info.Namespace.Get()
	ns.SchemaNs = meta.SourcePath.Get()
	ns.SchemaNs = strings.Replace(ns.SchemaNs, meta.RootPath.Get(), "", 1)
	ns.SchemaNs = extentioner.Strip(ns.SchemaNs)
	ns.SchemaNs = extentioner.Strip(ns.SchemaNs)

	if strings.HasPrefix(ns.CurrNs, "/") {
		ns.SchemaNs = strings.Replace(ns.SchemaNs, "/", "", 1)
	}

	ns.RelNs = ns.SchemaNs
	ns.SchemaNs = path.Join(ns.RootNs, ns.SchemaNs)

	return ns
}

func (t *Loader) typeNamespace(typ ct.CodeGenType) ct.RenderNamespace {
	return t.namespace(*typ.Source())
}
