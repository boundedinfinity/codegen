package loader

import (
	"boundedinfinity/codegen/model"
	"path"
	"strings"
)

func (t *Loader) processNamespace1(ctx *WalkContext) error {
	if ctx.Namespace.Input.Namespaces != nil {
		for _, child := range ctx.Namespace.Input.Namespaces {
			ctx.Namespace.Output.Children = append(ctx.Namespace.Output.Children, child.Name)
		}
	}

	return nil
}

func (t *Loader) processNamespace2(ctx *WalkContext) error {
	tmpls := make([]model.InputTemplate, 0)

	if ctx.Namespace.Input.Templates != nil {
		for _, tmpl := range ctx.Namespace.Input.Templates {
			if err := t.processTemplate1(ctx, &tmpls, tmpl); err != nil {
				return err
			}
		}
	}

	ns := ctx.Namespace.Output.Namespace

	for {
		if ns == t.rootName() || ns == "." {
			break
		}

		ns = path.Dir(ns)

		if vs, ok := t.templateMap[ns]; ok {
			for _, v := range vs {
				if err := t.processTemplate1(ctx, &tmpls, v); err != nil {
					return err
				}
			}
		}
	}

	t.templateMap[ctx.Namespace.Output.Namespace] = tmpls

	return nil
}

func (t *Loader) processNamespace3(ctx *WalkContext) error {
	output := ctx.Namespace.Output
	namespace := output.Namespace

	if strings.HasSuffix(namespace, model.NAMESPACE_BUILTIN) {
		return nil
	}

	vs := t.getTemplates(namespace, model.TemplateType_NAMESPACE)

	for _, v := range vs {
		outputTemplate := model.NewOutputTemplate()

		if err := t.processTemplate2(ctx, "", v, outputTemplate); err != nil {
			return err
		}

		output.Templates = append(output.Templates, outputTemplate)
	}

	return nil
}
