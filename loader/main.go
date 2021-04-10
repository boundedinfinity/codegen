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

	if ok, err := util.PathExists(modelPath); err != nil {
		return t.rc, err
	} else if !ok {
		return t.rc, fmt.Errorf("%v not found", modelPath)
	}

	if abs, err := filepath.Abs(modelPath); err != nil {
		return t.rc, err
	} else {
		t.rc.Input = optional.NewStringValue(abs)
	}

	if err := util.UnmarshalFromFile(t.rc.Input.Get(), &t.rc.Model); err != nil {
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

	// for k, v := range t.rc.Model.Components.Schemas {
	// 	t.errStack.Push(fmt.Sprintf("schemas[%v]", k))

	// 	// if err := t.validate_Model_Components_Schema_XbiGo(v.X_Bi_Go); err != nil {
	// 	// 	return err
	// 	// }

	// 	t.errStack.Pop()
	// }

	return nil
}
