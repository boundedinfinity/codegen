package processor

import "boundedinfinity/codegen/model"

func New() *Processor {
	return &Processor{
		inputs:         []*model.CodeGenProject{},
		sourceMap:      map[string]*model.CodeGenProject{},
		localSourceMap: map[string]*model.CodeGenProject{},
	}
}

type Processor struct {
	inputs         []*model.CodeGenProject
	sourceMap      map[string]*model.CodeGenProject
	localSourceMap map[string]*model.CodeGenProject
}

func (t *Processor) Process(projects ...model.CodeGenProject) error {
	for _, project := range projects {
		t.inputs = append(t.inputs, &project)

		if project.Source.Defined() {
			t.sourceMap[project.Source.Get()] = &project
		}

		if project.LocalSource.Defined() {
			t.sourceMap[project.LocalSource.Get()] = &project
		}
	}

	return nil
}
