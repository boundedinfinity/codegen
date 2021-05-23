package loader

import (
	"boundedinfinity/codegen/model"
	"boundedinfinity/codegen/util"
	"fmt"
	"path"
	"path/filepath"
	"strings"
)

func (t *Loader) processTemplate1() error {
	for namespace, inputTemplates := range t.inputTemplates {
		newTemplates := make([]model.InputTemplate, 0)

		for _, inputTemplate := range inputTemplates {
			newTemplate := model.InputTemplate{
				Header:    inputTemplate.Header,
				Type:      inputTemplate.Type,
				Namespace: inputTemplate.Namespace,
				Recurse:   inputTemplate.Recurse,
			}

			if filepath.IsAbs(inputTemplate.Path) {
				if ok, err := util.PathExists(inputTemplate.Path); err != nil {
					return err
				} else if !ok {
					return t.ErrTemplatePathNotFound(inputTemplate.Path)
				} else {
					newTemplate.Path = inputTemplate.Path
				}
			} else {
				var found bool

				for _, inputPath := range t.inputPaths {
					abs, err := filepath.Abs(inputPath)

					if err != nil {
						return err
					}

					abs = filepath.Dir(inputPath)
					abs = filepath.Join(abs, inputTemplate.Path)

					if ok, err := util.PathExists(abs); err != nil {
						return err
					} else if !ok {
						return t.ErrTemplatePathNotFound(inputTemplate.Path)
					} else {
						newTemplate.Path = abs
						found = true
						break
					}
				}

				if !found && t.OutputSpec.Info.InputDir != "" {
					if t.OutputSpec.Info.InputDir != "" {
						relPath := filepath.Join(t.OutputSpec.Info.InputDir, inputTemplate.Path)
						abs, err := filepath.Abs(relPath)

						if err != nil {
							return err
						}

						if ok, err := util.PathExists(abs); err != nil {
							return err
						} else if !ok {
							return t.ErrTemplatePathNotFound(inputTemplate.Path)
						} else {
							newTemplate.Path = abs
						}
					} else {
						return t.ErrTemplatePathNotFound(inputTemplate.Path)
					}
				}
			}

			newTemplates = append(newTemplates, newTemplate)
		}

		t.inputTemplates[namespace] = newTemplates
	}

	return nil
}

func (t *Loader) processTemplate2() error {
	shouldAddTemplate := func(it model.InputTemplate, name string) bool {
		if it.Namespace == "." {
			return true
		}

		if it.Recurse && strings.HasPrefix(name, it.Namespace) {
			return true
		}

		return false
	}

	for _, templates := range t.inputTemplates {
		for _, template := range templates {
			switch template.Type {
			case model.TemplateType_Model:
				for _, omodel := range t.outputModels {
					if shouldAddTemplate(template, omodel.FullName) {
						oTemplate := model.NewOutputTemplateWithOutput(template)
						name := path.Base(omodel.FullName)
						namespace := path.Dir(omodel.FullName)

						if err := t.processOutputTemplate(name, namespace, oTemplate); err != nil {
							return err
						}

						omodel.Templates = append(omodel.Templates, oTemplate)
					}
				}
			case model.TemplateType_Operation:
				for _, operation := range t.outputOperations {
					if shouldAddTemplate(template, operation.Name) {
						oTemplate := model.NewOutputTemplateWithOutput(template)
						name := path.Base(operation.Name)
						namespace := path.Dir(operation.Name)

						if err := t.processOutputTemplate(name, namespace, oTemplate); err != nil {
							return err
						}

						operation.Templates = append(operation.Templates, oTemplate)
					}
				}
			case model.TemplateType_Namespace:
				for name, namespace := range t.outputNamespace {
					if shouldAddTemplate(template, namespace.Name) {
						oTemplate := model.NewOutputTemplateWithOutput(template)
						var name2 string

						if name == "." {
							name2 = "root"
						} else {
							name2 = path.Base(name)
						}

						name2 = fmt.Sprintf("%v_ns", name2)
						namespace2 := namespace.Name

						if err := t.processOutputTemplate(name2, namespace2, oTemplate); err != nil {
							return err
						}

						namespace.Templates = append(namespace.Templates, oTemplate)
					}
				}
			}
		}
	}

	return nil
}

func (t *Loader) processOutputTemplate(name, namespace string, template *model.OutputTemplate) error {
	var templateTypeExt string
	var templateExt string
	var fullPath string

	templateTypeExt = template.Input
	templateTypeExt = filepath.Ext(templateTypeExt)
	templateTypeExt = strings.TrimPrefix(templateTypeExt, ".")

	templateExt = template.Input
	templateExt = strings.TrimSuffix(templateExt, filepath.Ext(template.Input))
	templateExt = filepath.Ext(templateExt)
	templateExt = strings.TrimPrefix(templateExt, ".")

	fullPath = name

	if t.OutputSpec.Info.FilenameMarker != model.DEFAULT_FILENAME_DISABLE {
		if t.OutputSpec.Info.FilenameMarker != "" {
			fullPath = fmt.Sprintf("%v.%v", fullPath, t.OutputSpec.Info.FilenameMarker)
		} else {
			fullPath = fmt.Sprintf("%v.%v", fullPath, model.DEFAULT_FILENAME_MARKER)
		}
	}

	fullPath = fmt.Sprintf("%v.%v", fullPath, templateExt)
	fullPath = path.Join(t.OutputSpec.Info.OutputDir, namespace, fullPath)

	template.Output = fullPath
	return nil
}
