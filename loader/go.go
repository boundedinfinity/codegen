package loader

import (
	"boundedinfinity/codegen/model"
	"boundedinfinity/codegen/util"
	"fmt"
	"path/filepath"

	"github.com/boundedinfinity/optional"
)

func (t *Loader) searchFromInput(p optional.StringOptional) (optional.StringOptional, error) {
	if p.IsEmpty() {
		return p, model.CannotBeEmptyErr
	}

	d := filepath.Dir(t.rc.Input.Get())
	if abs, ok := util.FileSearch(p, optional.NewStringValue(d)); ok {
		return optional.NewStringValue(abs), nil
	} else {
		return p, fmt.Errorf("invalid path: %v", p.Get())
	}
}

func (t *Loader) searchFromTemplateRoot(p optional.StringOptional) (optional.StringOptional, error) {
	if p.IsEmpty() {
		return p, model.CannotBeEmptyErr
	}

	if abs, ok := util.FileSearch(p, t.rc.Model.X_Bi_Go.TemplateRoot); ok {
		return optional.NewStringValue(abs), nil
	} else {
		return p, fmt.Errorf("invalid path: %v", p.Get())
	}
}

func (t *Loader) searchFromGenRoot(p optional.StringOptional) (optional.StringOptional, error) {
	if p.IsEmpty() {
		return p, model.CannotBeEmptyErr
	}

	if abs, ok := util.FileSearch(p, t.rc.Model.X_Bi_Go.GenRoot); ok {
		return optional.NewStringValue(abs), nil
	} else {
		return p, fmt.Errorf("invalid path: %v", p.Get())
	}
}

func (t *Loader) pathFromPackage(tmpl *model.X_Bi_GoTemplate) (optional.StringOptional, error) {
	// mn := t.rc.Model.X_Bi_Go.Module.Name.Get()
	// gr := t.rc.Model.X_Bi_Go.GenRoot.Get()
	// pkg := tmpl.Package.Get()
	// i := tmpl.Input.Get()

	// if !strings.Contains(pkg, mn) {
	// 	return optional.NewStringEmpty(), fmt.Errorf("package %v not in module name %v", pkg, mn)
	// }

	// var o string
	// o = strings.TrimPrefix(pkg, mn)
	// o = strings.TrimPrefix(o, "/")
	// o = filepath.Join(gr, o)

	// e := filepath.Ext(i)
	// i = filepath.Base(i)
	// i = strings.TrimSuffix(i, e)
	// o = filepath.Join(o, i)

	// return optional.NewStringValue(o), nil
	return optional.NewStringEmpty(), nil
}

func (t *Loader) validate_Model_XBiGo() error {
	t.errStack.Push("x-bi-go")

	if t.rc.Model.X_Bi_Go == nil {
		return nil
	}

	xbi := t.rc.Model.X_Bi_Go

	{
		t.errStack.Push("templateRoot")

		if xbi.TemplateRoot.IsEmpty() {
			return model.CannotBeEmptyErr
		}

		if abs, err := t.searchFromInput(xbi.TemplateRoot); err != nil {
			return fmt.Errorf("invalid path: %v", xbi.TemplateRoot.Get())
		} else {
			xbi.TemplateRoot = optional.NewStringValue(abs.Get())
		}

		t.errStack.Pop()
	}

	{
		t.errStack.Push("genRoot")

		if xbi.GenRoot.IsEmpty() {
			if abs, err := filepath.Abs(".codegen"); err != nil {
				return err
			} else {
				xbi.GenRoot = optional.NewStringValue(abs)
			}
		} else {
			if ok := filepath.IsAbs(xbi.GenRoot.Get()); !ok {
				if abs, err := filepath.Abs(xbi.GenRoot.Get()); err != nil {
					return err
				} else {
					xbi.GenRoot = optional.NewStringValue(abs)
				}
			}
		}

		t.errStack.Pop()
	}

	if err := t.validate_Model_XBiGo_Module(); err != nil {
		return err
	}

	if err := t.validate_Model_XBiGo_Templates(); err != nil {
		return err
	}

	t.errStack.Pop()
	return nil
}

func (t *Loader) validate_Model_XBiGo_Module() error {
	t.errStack.Push("module")

	if t.rc.Model.X_Bi_Go.Module == nil {
		return model.CannotBeEmptyErr
	}

	module := t.rc.Model.X_Bi_Go.Module

	{
		t.errStack.Push("name")

		if module.Name.IsEmpty() {
			return model.CannotBeEmptyErr
		}

		t.errStack.Pop()
	}

	{
		t.errStack.Push("version")

		if module.Version.IsEmpty() {
			return model.CannotBeEmptyErr
		}

		t.errStack.Pop()
	}

	// if module.Templates != nil {
	// 	for i, tmpl := range module.Templates {
	// 		t.errStack.Push(fmt.Sprintf("templates[%v]", i))

	// 		if tmpl == nil {
	// 			return model.CannotBeEmptyErr
	// 		}

	// 		{
	// 			t.errStack.Push("input")
	// 			if tmpl.Input.IsEmpty() {
	// 				return model.CannotBeEmptyErr
	// 			}

	// 			if abs, err := t.searchFromTemplateRoot(tmpl.Input); err != nil {
	// 				return model.CannotBeEmptyErr
	// 			} else {
	// 				tmpl.Input = abs
	// 			}

	// 			t.errStack.Pop()
	// 		}

	// 		{
	// 			t.errStack.Push("output")

	// 			if tmpl.Output.IsEmpty() {
	// 				return model.CannotBeEmptyErr
	// 			}

	// 			abs := filepath.Join(module.GenRoot.Get(), tmpl.Output.Get())
	// 			tmpl.Output = optional.NewStringValue(abs)

	// 			t.errStack.Pop()
	// 		}

	// 		t.errStack.Pop()
	// 	}
	// }

	t.errStack.Pop()
	return nil
}

func (t *Loader) validate_Model_XBiGo_Templates() error {
	if t.rc.Model.X_Bi_Go.Templates == nil {
		return nil
	}

	for i, tmpl := range t.rc.Model.X_Bi_Go.Templates {
		t.errStack.Push(fmt.Sprintf("templates[%v]", i))

		{
			t.errStack.Push("input")
			if tmpl.Input.IsEmpty() {
				return model.CannotBeEmptyErr
			}

			if abs, err := t.searchFromTemplateRoot(tmpl.Input); err != nil {
				return model.CannotBeEmptyErr
			} else {
				tmpl.Input = abs
			}
			t.errStack.Pop()
		}

		// {
		// 	t.errStack.Push("package")

		// 	if tmpl.Package.IsEmpty() {
		// 		tmpl.Package = optional.NewStringValue("main")
		// 	}

		// 	t.errStack.Pop()
		// }

		// {
		// 	t.errStack.Push("output")

		// 	if tmpl.Output.IsDefined() {
		// 		if ok := filepath.IsAbs(tmpl.Output.Get()); !ok {
		// 			abs := filepath.Join(t.rc.Model.X_Bi_Go.Module.GenRoot.Get(), tmpl.Output.Get())
		// 			tmpl.Output = optional.NewStringValue(abs)
		// 		}
		// 	} else {
		// 		if abs, err := t.pathFromPackage(tmpl); err != nil {
		// 			return err
		// 		} else {
		// 			tmpl.Output = abs
		// 		}
		// 	}

		// 	t.errStack.Pop()
		// }

		t.errStack.Pop()
	}

	return nil
}

// func (t *Loader) validate_Model_Components_Schema_XbiGo(s *model.OpenApiV310ExtentionSchema) error {
// 	if s == nil {
// 		return nil
// 	}

// 	t.errStack.Push("x-bi-go")

// 	// if err := t.validate_Model_XBiGo_Templates(s.Templates); err != nil {
// 	// 	return err
// 	// }

// 	t.errStack.Pop()
// 	return nil
// }
