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
		t.operationNamespace(operation)

		if operation.Input.Defined() {
			t.operationIoNamespace(*operation, operation.Input.Get())
		}

		if operation.Output.Defined() {
			t.operationIoNamespace(*operation, operation.Output.Get())
		}

		return nil
	}, t.projectManager.Merged)

	err = ct.WalkTemplateType(func(_ *ct.CodeGenProject, _ *ct.CodeGenProjectTemplates, file *ct.CodeGenProjectTemplateFile) error {
		ns := t.namespace(file.SourceMeta)
		*file.Namespace() = ns
		return nil
	}, t.projectManager.Merged)

	return nil
}

func (t *Loader) namespace(meta ct.SourceMeta) ct.RenderNamespace {
	var ns ct.RenderNamespace

	ns.RootNs = t.projectManager.Merged.Info.Namespace.Get()
	ns.SchemaNs = path.Base(meta.SourcePath.Get())

	for strings.Contains(ns.SchemaNs, ".") {
		ns.SchemaNs = extentioner.Strip(ns.SchemaNs)
	}

	ns.SchemaNs = path.Join(path.Dir(meta.SourcePath.Get()), ns.SchemaNs)
	ns.SchemaNs = strings.Replace(ns.SchemaNs, meta.RootPath.Get(), "", 1)

	if strings.HasPrefix(ns.SchemaNs, "/") {
		ns.SchemaNs = strings.Replace(ns.SchemaNs, "/", "", 1)
	}

	ns.RelNs = ns.SchemaNs
	ns.SchemaNs = path.Join(ns.RootNs, ns.SchemaNs)

	return ns
}

func (t *Loader) typeNamespace(typ ct.CodeGenType) ct.RenderNamespace {
	ns := t.namespace(*typ.Source())
	ns.CurrNs = ns.SchemaNs
	return ns
}

func (t *Loader) operationNamespace(operation *ct.CodeGenProjectOperation) {
	operation.RenderNamespace = t.namespace(operation.SourceMeta)
	operation.CurrNs = operation.SchemaNs
}

func (t *Loader) operationIoNamespace(operation ct.CodeGenProjectOperation, typ ct.CodeGenType) {
	switch c := typ.(type) {
	case *ct.CodeGenTypeArray:
		*c.Namespace() = *operation.Namespace()
		t.operationIoNamespace(operation, c.Items)
	case *ct.CodeGenTypeRef:
		*c.Namespace() = *operation.Namespace()
		t.operationIoNamespace(operation, c.Resolved)
	default:
		typ.Namespace().CurrNs = operation.SchemaNs
	}
}
