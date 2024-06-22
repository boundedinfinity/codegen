package processor

import (
	"boundedinfinity/codegen/model"
	"encoding/json"
	"os"

	"github.com/boundedinfinity/go-commoner/functional/optioner"
	"github.com/boundedinfinity/go-commoner/idiomatic/environmenter"
	"github.com/boundedinfinity/go-commoner/idiomatic/extentioner"
	"github.com/boundedinfinity/go-commoner/idiomatic/pather"
	"github.com/boundedinfinity/go-mimetyper/file_extention"
	"github.com/invopop/yaml"
)

func New() *Processor {
	return &Processor{
		projects:  []*model.CodeGenProject{},
		typeIdMap: map[string]model.CodeGenType{},
	}
}

type Processor struct {
	combined  model.CodeGenProject
	projects  []*model.CodeGenProject
	typeIdMap map[string]model.CodeGenType
}

func (t *Processor) ProcessFiles(paths ...string) error {
	var projects []model.CodeGenProject
	wd, err := os.Getwd()

	if err != nil {
		return err
	}

	for _, path := range paths {
		path = environmenter.Substitue(path)
		if !pather.Paths.IsAbs(path) {
			path = pather.Paths.Join(wd, path)
		}

		var project model.CodeGenProject

		bs, err := os.ReadFile(path)
		if err != nil {
			return ErrCodeGenCantReadFileFn(path, err)
		}

		switch v := extentioner.Ext(path); v {
		case file_extention.FileExtentions.Yaml.String(), file_extention.FileExtentions.Yml.String():
			if jbs, err := yaml.YAMLToJSON(bs); err != nil {
				return ErrCodeGenCantReadFileFn(path, err)
			} else {
				if err := json.Unmarshal(jbs, &project); err != nil {
					return ErrCodeGenCantReadFileFn(path, err)
				}
			}
		case file_extention.FileExtentions.Json.String():
			if err := json.Unmarshal(bs, &project); err != nil {
				return ErrCodeGenCantReadFileFn(path, err)
			}
		default:
			return ErrCodeGenUnsupportedFileTypeFn(v)
		}

		projects = append(projects, project)
	}

	return t.ProcessProjects(projects...)
}

func (t *Processor) ProcessProjects(projects ...model.CodeGenProject) error {
	for _, project := range projects {
		t.projects = append(t.projects, &project)
	}

	if err := t.processOutputRoot(t.projects...); err != nil {
		return err
	}

	if err := t.processPackage(t.projects...); err != nil {
		return err
	}

	if err := t.processTypes(); err != nil {
		return err
	}

	if err := t.checkCombinedTypes(); err != nil {
		return err
	}

	if err := t.calculatePackageNames(); err != nil {
		return err
	}

	if err := t.validate(); err != nil {
		return err
	}

	return nil
}

func (t *Processor) validate() error {
	if t.combined.OutputRoot.Empty() {
		t.combined.OutputRoot = optioner.Some(".")
	}

	if abs, err := pather.Paths.AbsErr("."); err != nil {
		return err
	} else {
		t.combined.OutputRoot = optioner.Some(abs)
	}

	return nil
}
