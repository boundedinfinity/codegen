package loader

import (
	"boundedinfinity/codegen/model"
	"boundedinfinity/codegen/util"
	"fmt"
	"path/filepath"
	"strings"
)

type Loader struct {
	errPath *model.StrStack
}

func New() *Loader {
	return &Loader{
		errPath: model.NewStrStack(),
	}
}

func (t *Loader) Load(modelPath, configPath string) (model.RunContext, error) {
	rc := model.RunContext{
		ModelPath: modelPath,
	}

	t.errPath.Push("loader")

	if err := t.loadModel(&rc); err != nil {
		ep := strings.Join(t.errPath.S(), ".")
		return rc, fmt.Errorf("%v: %w", ep, err)
	}

	return rc, nil
}

func (t *Loader) loadModel(rc *model.RunContext) error {
	t.errPath.Push("model")

	if rc.ModelPath == "" {
		return model.CannotBeEmptyErr
	}

	if abs, ok := util.FileSearch(&rc.ModelPath); ok {
		rc.ModelPath = abs
		t.errPath.Pop()
		t.errPath.Push(fmt.Sprintf("model[%v]", abs))
	} else {
		return fmt.Errorf("%v not found", rc.ModelPath)
	}

	var m model.OpenApiV310

	if err := util.UnmarshalFromFile(rc.ModelPath, &m); err != nil {
		return fmt.Errorf("load failed %w", err)
	}

	rc.Model = m

	if err := t.loadGlobal(rc); err != nil {
		return err
	}

	return nil
}

func (t *Loader) loadGlobal(rc *model.RunContext) error {
	t.errPath.Push("openapi")

	if util.StrIsEmpty(rc.Model.Openapi) {
		return model.CannotBeEmptyErr
	}

	t.errPath.Pop()

	if err := t.loadInfo(rc); err != nil {
		return err
	}

	if err := t.loadXbiGo(rc); err != nil {
		return err
	}

	return nil
}

func (t *Loader) loadInfo(rc *model.RunContext) error {
	t.errPath.Push("info")

	if rc.Model.Info == nil {
		return model.CannotBeEmptyErr
	}

	t.errPath.Push("version")
	if util.StrIsEmpty(rc.Model.Info.Version) {
		return model.CannotBeEmptyErr
	}
	t.errPath.Pop()

	t.errPath.Push("title")
	if util.StrIsEmpty(rc.Model.Info.Title) {
		return model.CannotBeEmptyErr
	}
	t.errPath.Pop()

	t.errPath.Push("description")
	if util.StrIsEmpty(rc.Model.Info.Description) {
		return model.CannotBeEmptyErr
	}
	t.errPath.Pop()

	t.errPath.Pop()
	return nil
}

func (t *Loader) loadXbiGo(rc *model.RunContext) error {
	t.errPath.Push("x-bi-go")

	if rc.Model.X_Bi_Go == nil {
		return model.CannotBeEmptyErr
	}

	{
		t.errPath.Push("module")
		if util.StrIsEmpty(rc.Model.X_Bi_Go.Module) {
			return model.CannotBeEmptyErr
		}
		t.errPath.Pop()
	}

	{
		t.errPath.Push("version")
		if util.StrIsEmpty(rc.Model.X_Bi_Go.Module) {
			return model.CannotBeEmptyErr
		}
		t.errPath.Pop()
	}

	{
		t.errPath.Push("input")

		if rc.Model.X_Bi_Go.Input == nil {
			return model.CannotBeEmptyErr
		}

		if abs, ok := util.FileSearch(rc.Model.X_Bi_Go.Input, util.StrPrt(filepath.Dir(rc.ModelPath))); ok {
			*rc.Model.X_Bi_Go.Input = abs
			t.errPath.Pop()
			t.errPath.Push(fmt.Sprintf("input[%v]", abs))
		} else {
			return fmt.Errorf("%v not found", *rc.Model.X_Bi_Go.Input)
		}

		t.errPath.Pop()
	}

	t.errPath.Pop()
	return nil
}

// func (t *Loader) loadExtentions(rctx *model.RunContext) error {
// 	if err := t.loadExtentionsGlobal(rctx); err != nil {
// 		return model.NewXerr("x-bi-go", err)
// 	}

// 	if err := t.loadExtentionsComponents(rctx); err != nil {
// 		return model.NewXerr("x-bi-go", err)
// 	}

// 	return nil
// }

// func (t *Loader) loadExtentionsGlobal(rctx *model.RunContext) error {
// 	if rctx.Model.X_Bi_Go == nil {
// 		return errors.New("cannot be empty")
// 	}

// 	if util.StrIsEmpty(rctx.Model.X_Bi_Go.Version) {
// 		return model.XerrF("version", "cannot be emtpy")
// 	}

// 	if util.StrIsEmpty(rctx.Model.X_Bi_Go.Module) {
// 		return model.XerrF("module", "cannot be emtpy")
// 	}

// 	if abs, ok := util.FileSearch(rctx.Model.X_Bi_Go.Input, util.StrPrt(filepath.Dir(rctx.ModelPath))); ok {
// 		*rctx.Model.X_Bi_Go.Input = abs
// 	} else {
// 		return model.XerrF("model", "not found")
// 	}

// 	if rctx.Model.X_Bi_Go.Requires != nil {
// 		for i, r := range rctx.Model.X_Bi_Go.Requires {

// 			if r.Package == "" {
// 				errPath := fmt.Sprintf("requires[%v].package", i)
// 				return model.XerrF(errPath, "cannot be empty")
// 			}

// 			if r.Version == "" {
// 				errPath := fmt.Sprintf("requires[%v].version", i)
// 				return model.XerrF(errPath, "cannot be empty")
// 			}
// 		}
// 	}

// 	return nil
// }

// func (t *Loader) loadExtentionsComponents(rctx *model.RunContext) error {
// 	if rctx.Model.Components == nil {
// 		return nil
// 	}

// 	return nil
// }
