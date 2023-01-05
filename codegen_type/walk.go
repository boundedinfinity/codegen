package codegen_type

import "errors"

var ErrExit = errors.New("walker exit")

type projectWalker struct {
	infoFn           func(CodeGenProject, CodeGenInfo) error
	operationFn      func(CodeGenProject, CodeGenProjectOperation) error
	templateFn       func(CodeGenProject, CodeGenProjectTemplates, CodeGenProjectTemplateFile) error
	typeFn           func(CodeGenProject, CodeGenType) error
	typeStringFn     func(CodeGenProject, *CodeGenTypeString) error
	typeArrayFn      func(CodeGenProject, *CodeGenTypeArray) error
	typeArrayItemsFn func(CodeGenProject, *CodeGenTypeArray, CodeGenType) error
	typeObjectFn     func(CodeGenProject, *CodeGenTypeObject) error
	typeObjectPropFn func(CodeGenProject, *CodeGenTypeObject, CodeGenType) error
}

func Walker() *projectWalker {
	return &projectWalker{}
}

func (w *projectWalker) Walk(projects ...CodeGenProject) error {
	for _, project := range projects {
		if err := w.walk(project); err != nil {
			if errors.Is(err, ErrExit) {
				return nil
			}

			return err
		}
	}

	return nil
}

func (w *projectWalker) walk(project CodeGenProject) error {
	if w.infoFn != nil {
		if err := w.infoFn(project, project.Info); err != nil {
			return err
		}
	}

	if w.operationFn != nil {
		for _, operation := range project.Operations {
			if err := w.operationFn(project, operation); err != nil {
				return err
			}
		}
	}

	if w.typeFn != nil {
		for _, typ := range project.Types {
			if err := w.typeFn(project, typ); err != nil {
				return err
			}

			switch c := typ.(type) {
			case *CodeGenTypeObject:
				if w.typeObjectFn != nil {
					if err := w.typeObjectFn(project, c); err != nil {
						return err
					}
				}
				if w.typeObjectPropFn != nil {
					for _, property := range c.Properties {
						if err := w.typeObjectPropFn(project, c, property); err != nil {
							return err
						}
					}
				}
			case *CodeGenTypeArray:
				if w.typeArrayFn != nil {
					if err := w.typeArrayFn(project, c); err != nil {
						return err
					}
				}

				if w.typeArrayItemsFn != nil {
					if err := w.typeArrayItemsFn(project, c, c.Items); err != nil {
						return err
					}
				}
			}
		}
	}

	if w.templateFn != nil && project.Templates.Files != nil {
		for _, file := range project.Templates.Files {
			if file == nil {
				continue
			}

			if err := w.templateFn(project, project.Templates, *file); err != nil {
				return err
			}
		}
	}

	return nil
}

func (w *projectWalker) Info(fn func(CodeGenProject, CodeGenInfo) error) *projectWalker {
	w.infoFn = fn
	return w
}

func (w *projectWalker) Operation(fn func(CodeGenProject, CodeGenProjectOperation) error) *projectWalker {
	w.operationFn = fn
	return w
}

func (w *projectWalker) Type(fn func(CodeGenProject, CodeGenType) error) *projectWalker {
	w.typeFn = fn
	return w
}

func (t *projectWalker) TypeString(v func(CodeGenProject, *CodeGenTypeString) error) *projectWalker {
	t.typeStringFn = v
	return t
}

func (t *projectWalker) TypeArray(v func(CodeGenProject, *CodeGenTypeArray) error) *projectWalker {
	t.typeArrayFn = v
	return t
}

func (t *projectWalker) TypeArrayItems(v func(CodeGenProject, *CodeGenTypeArray, CodeGenType) error) *projectWalker {
	t.typeArrayItemsFn = v
	return t
}

func (t *projectWalker) TypeObject(v func(CodeGenProject, *CodeGenTypeObject) error) *projectWalker {
	t.typeObjectFn = v
	return t
}

func (t *projectWalker) TypeObjectProperty(v func(CodeGenProject, *CodeGenTypeObject, CodeGenType) error) *projectWalker {
	t.typeObjectPropFn = v
	return t
}

func (w *projectWalker) Template(fn func(CodeGenProject, CodeGenProjectTemplates, CodeGenProjectTemplateFile) error) *projectWalker {
	w.templateFn = fn
	return w
}
