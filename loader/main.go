package loader

import (
	"boundedinfinity/codegen/model"
	"boundedinfinity/codegen/util"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/boundedinfinity/optional"
)

type Loader struct {
	errStack     *model.StrStack
	pointerStack *model.StrStack
	rc           model.RunContext
}

func New() *Loader {
	return &Loader{
		errStack:     model.NewStrStack(),
		pointerStack: model.NewStrStack(),
	}
}

func (t *Loader) Load(modelPath string) (model.RunContext, error) {
	t.rc = model.RunContext{}

	if abs, err := filepath.Abs(modelPath); err != nil {
		return t.rc, err
	} else {
		t.rc.Input = abs
	}

	if err := util.UnmarshalFromFile(t.rc.Input, &t.rc.Model); err != nil {
		return t.rc, err
	}

	if err := t.validate_Model(); err != nil {
		ep := strings.Join(t.errStack.S(), ".")
		return t.rc, fmt.Errorf("%v: %w", ep, err)
	}

	return t.rc, nil
}

func (t *Loader) validate_Model() error {
	t.errStack.Push("model")

	{
		t.errStack.Push("openapi")

		if t.rc.Model.Openapi.IsEmpty() {
			return model.CannotBeEmptyErr
		}

		t.errStack.Pop()
	}

	if err := t.validate_Model_Info(); err != nil {
		return err
	}

	if err := t.validate_Model_Servers(); err != nil {
		return err
	}

	if err := t.validate_Model_XBiGo(); err != nil {
		return err
	}

	if err := t.validate_Model_Components(); err != nil {
		return err
	}

	t.errStack.Pop()
	return nil
}

func (t *Loader) validate_Model_XBiGo() error {
	t.errStack.Push("x-bi-go")

	if t.rc.Model.X_Bi_Go == nil {
		return model.CannotBeEmptyErr
	}

	{
		t.errStack.Push("input")
		if t.rc.Model.X_Bi_Go.Input.IsDefined() {
			return model.CannotBeEmptyErr
		}

		d := filepath.Dir(t.rc.Input)
		if abs, ok := util.FileSearch(t.rc.Model.X_Bi_Go.Input, optional.NewStringValue(d)); ok {
			t.rc.Model.X_Bi_Go.Input = optional.NewStringValue(abs)
		} else {
			return fmt.Errorf("invalid path: %v", t.rc.Model.X_Bi_Go.Input.Get())
		}

		t.errStack.Pop()
	}

	{
		t.errStack.Push("version")
		if t.rc.Model.X_Bi_Go.Version.IsEmpty() {
			return model.CannotBeEmptyErr
		}
		t.errStack.Pop()
	}

	{
		t.errStack.Push("module")
		if t.rc.Model.X_Bi_Go.Module.IsEmpty() {
			return model.CannotBeEmptyErr
		}
		t.errStack.Pop()
	}

	{
		t.errStack.Push("output")
		if t.rc.Model.X_Bi_Go.Output.IsEmpty() {
			return model.CannotBeEmptyErr
		}
		t.errStack.Pop()
	}

	if err := t.validate_Model_XBiGo_Templates(t.rc.Model.X_Bi_Go.Templates); err != nil {
		return err
	}

	t.errStack.Pop()
	return nil
}

func (t *Loader) validate_Model_XBiGo_Templates(tmpls []model.XBiGoTemplate) error {
	if tmpls == nil {
		return nil
	}

	for i, tmpl := range tmpls {
		t.errStack.Push(fmt.Sprintf("templates[%v]", i))

		{
			t.errStack.Push("input")

			if tmpl.Input.IsEmpty() {
				return model.CannotBeEmptyErr
			}

			if abs, ok := util.FileSearch(tmpl.Input, t.rc.Model.X_Bi_Go.Input); ok {
				tmpl.Input = optional.NewStringValue(abs)
			} else {
				return fmt.Errorf("invalid path: %v", tmpl.Input.Get())
			}

			t.errStack.Pop()
		}

		{
			t.errStack.Push("package")

			if tmpl.Package.IsEmpty() {
				tmpl.Package = optional.NewStringValue("main")
			} else {
				p := tmpl.Package.Get()
				p = strings.ReplaceAll(p, t.rc.Model.X_Bi_Go.Module.Get(), "")
				p = strings.Replace(p, "/", "", 1)

				if p == "" {
					p = "main"
				}

				tmpl.Package = optional.NewStringValue(p)
			}

			t.errStack.Pop()
		}

		{
			t.errStack.Push("output")

			if tmpl.Output.IsEmpty() {
				if tmpl.Package.IsEmpty() {
					return model.CannotBeEmptyErr
				} else {
					if tmpl.Package.Get() == "main" {
						tmpl.Package = optional.NewStringValue("")
					} else {
						tmpl.Output = optional.NewStringValue(tmpl.Package.Get())
					}
				}
			}

			p := filepath.Join(t.rc.Model.X_Bi_Go.Output.Get(), tmpl.Output.Get())

			if abs, err := filepath.Abs(p); err != nil {
				return err
			} else {
				tmpl.Output = optional.NewStringValue(abs)
			}

			t.errStack.Pop()
		}

		t.errStack.Pop()
	}

	return nil
}

func (t *Loader) validate_Model_Info() error {
	if t.rc.Model.Info == nil {
		return nil
	}

	t.errStack.Push("info")

	{
		t.errStack.Push("version")
		if t.rc.Model.Info.Version.IsEmpty() {
			return model.CannotBeEmptyErr
		}
		t.errStack.Pop()
	}

	{
		t.errStack.Push("title")
		if t.rc.Model.Info.Title.IsEmpty() {
			return model.CannotBeEmptyErr
		}
		t.errStack.Pop()
	}

	{
		t.errStack.Push("description")
		if t.rc.Model.Info.Description.IsEmpty() {
			return model.CannotBeEmptyErr
		}
		t.errStack.Pop()
	}

	t.errStack.Pop()
	return nil
}

func (t *Loader) validate_Model_Servers() error {
	if t.rc.Model.Servers == nil {
		return nil
	}

	for i, s := range t.rc.Model.Servers {
		t.errStack.Push(fmt.Sprintf("servers[%v]", i))

		{
			t.errStack.Push("url")
			if s.Url.IsEmpty() {
				return model.CannotBeEmptyErr
			}
			t.errStack.Pop()
		}

		t.errStack.Pop()
	}

	return nil

}

func (t *Loader) validate_Model_Components() error {
	if t.rc.Model.Components == nil {
		return nil
	}

	t.errStack.Push("components")

	if err := t.validate_Model_Components_Schemas(); err != nil {
		return err
	}

	t.errStack.Pop()
	return nil
}

func (t *Loader) validate_Model_Components_Schemas() error {
	if t.rc.Model.Components.Schemas == nil {
		return nil
	}

	for k, v := range t.rc.Model.Components.Schemas {
		t.errStack.Push(fmt.Sprintf("schemas[%v]", k))

		if err := t.validate_Model_Components_Schema_XbiGo(v.X_Bi_Go); err != nil {
			return err
		}

		t.errStack.Pop()
	}

	t.errStack.Pop()
	return nil
}

func (t *Loader) validate_Model_Components_Schema_XbiGo(s *model.OpenApiV310ExtentionSchema) error {
	if s == nil {
		return nil
	}

	if err := t.validate_Model_XBiGo_Templates(s.Templates); err != nil {
		return err
	}

	t.errStack.Pop()
	return nil
}
