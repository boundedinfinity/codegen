package loader

import (
	"boundedinfinity/codegen/model"
	"boundedinfinity/codegen/util"
	"path/filepath"
)

func (t *Loader) getTemplates(namespace string, typ model.TemplateTypeEnum) (tmpls []model.InputTemplate) {
	if namespace == "" || typ == "" {
		return tmpls
	}

	if vs, ok := t.templateMap[namespace]; !ok {
		return tmpls
	} else {
		for _, v := range vs {
			if v.Type == string(typ) {
				tmpls = append(tmpls, v)
			}
		}
	}

	return tmpls
}

func (t *Loader) processTemplate1(ctx *model.WalkContext, tmpls *[]model.InputTemplate, input model.InputTemplate) error {
	if input.Path == "" {
		return t.ErrCannotBeEmpty()
	}

	if input.Type == "" {
		return t.ErrCannotBeEmpty()
	}

	if _, err := model.TemplateTypeEnumParse(input.Type); err != nil {
		return err
	}

	for _, v := range *tmpls {
		if v.Path == input.Path && v.Type == input.Type {
			return nil
		}
	}

	var output model.InputTemplate

	if filepath.IsAbs(input.Path) {
		if ok, err := util.PathExists(input.Path); err != nil {
			return err
		} else if !ok {
			return t.ErrNotFound()
		} else {
			output.Path = input.Path
		}
	} else {
		relPath := filepath.Join(t.OutputSpec.Info.InputDir, input.Path)
		abs, err := filepath.Abs(relPath)

		if err != nil {
			return err
		}

		if ok, err := util.PathExists(abs); err != nil {
			return err
		} else if !ok {
			return t.ErrNotFound()
		} else {
			output.Path = abs
		}
	}

	output.Header = input.Header
	output.Type = input.Type

	for _, v := range *tmpls {
		if v.Path == output.Path && v.Type == output.Type {
			return nil
		}
	}

	*tmpls = append(*tmpls, output)

	return nil
}

func (t *Loader) processTemplate2(ctx *model.WalkContext, name string, input model.InputTemplate, output *model.OutputTemplate) error {
	// var tmplExt string
	// var tmplName string
	// var relPath string
	// var abs string
	// var fn string

	// info := t.OutputSpec.Info

	// relPath = ctx.Namespace.Output.Namespace
	// relPath = strings.TrimPrefix(relPath, t.rootName())
	// relPath = strings.TrimPrefix(relPath, "/")

	// tmplName = input.Path
	// tmplName = filepath.Base(tmplName)
	// tmplName = util.TrimTemplateExt(tmplName)
	// tmplExt = filepath.Ext(tmplName)
	// tmplName = strings.TrimSuffix(tmplName, tmplExt)
	// tmplExt = strings.TrimPrefix(tmplExt, ".")

	// if name == "" {
	// 	fn = tmplName
	// } else {
	// 	fn = name
	// }

	// if info.FilenameMarker != "" && info.FilenameMarker != model.DEFAULT_FILENAME_DISABLE {
	// 	fn = fmt.Sprintf("%v.%v", fn, info.FilenameMarker)
	// }

	// fn = fmt.Sprintf("%v.%v", fn, tmplExt)
	// abs = path.Join(info.OutputDir, relPath, fn)

	// output.Input = input.Path
	// output.Output = abs
	// output.Header = t.splitDescription(input.Header)
	return nil
}
