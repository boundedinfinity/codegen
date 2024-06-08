package processor

import (
	"boundedinfinity/codegen/model"
	"encoding/json"
	"os"

	"github.com/boundedinfinity/go-commoner/idiomatic/environmenter"
	"github.com/boundedinfinity/go-commoner/idiomatic/extentioner"
	"github.com/boundedinfinity/go-commoner/idiomatic/pather"
	"github.com/invopop/yaml"
)

func New() *Processor {
	return &Processor{
		projects:   []*model.CodeGenProject{},
		sourcesMap: map[string][]string{},
		typeIdMap:  map[string]model.CodeGenType{},
	}
}

type Processor struct {
	combined   model.CodeGenProject
	projects   []*model.CodeGenProject
	sourcesMap map[string][]string
	typeIdMap  map[string]model.CodeGenType
}

func (t *Processor) ProcessFiles(paths ...string) error {
	var projects []model.CodeGenProject
	wd, err := os.Getwd()

	if err != nil {
		return err
	}

	for _, path := range paths {
		path = environmenter.Sub(path)

		if !pather.Paths.IsAbs(path) {
			path = pather.Join(wd, path)
		}

		var project model.CodeGenProject
		bs, err := os.ReadFile(path)

		if err != nil {
			return ErrCodeGenCantReadFileFn(path, err)
		}

		switch v := extentioner.Ext(path); v {
		case ".yaml", ".yml":
			if jbs, err := yaml.YAMLToJSON(bs); err != nil {
				return ErrCodeGenCantReadFileFn(path, err)
			} else {
				if err := json.Unmarshal(jbs, &project); err != nil {
					return ErrCodeGenCantReadFileFn(path, err)
				}
			}
		case ".json":
			if err := json.Unmarshal(bs, &project); err != nil {
				return ErrCodeGenCantReadFileFn(path, err)
			}
		default:
			return ErrCodeGenUnsupportedFileTypeFn(v)
		}

		project.Sources = append(project.Sources, path)
		projects = append(projects, project)
	}

	return t.ProcessProjects(projects...)
}

func (t *Processor) ProcessProjects(projects ...model.CodeGenProject) error {
	for _, project := range projects {
		t.projects = append(t.projects, &project)
	}

	if err := t.pass01(); err != nil {
		return err
	}

	if err := t.pass02(); err != nil {
		return err
	}

	if err := t.pass03(); err != nil {
		return err
	}

	return nil
}

func (t *Processor) pass01() error {
	for _, project := range t.projects {
		for _, source := range project.Sources {
			if _, ok := t.sourcesMap[source]; !ok {
				t.sourcesMap[source] = []string{source}
			}
		}
	}

	return nil
}

func (t *Processor) pass02() error {
	for _, project := range t.projects {
		for _, typ := range project.Types {
			if _, ok := t.typeIdMap[typ.TypeId().Get()]; ok {
				return ErrCodeGenTypeSchemaIdDuplicateFn(typ)
			}

			t.typeIdMap[typ.TypeId().Get()] = typ
			t.combined.Types = append(t.combined.Types, typ)
		}
	}

	return nil
}

func (t *Processor) pass03() error {
	for _, typ := range t.combined.Types {
		if err := t.checkType(typ); err != nil {
			return err
		}
	}

	return nil
}
