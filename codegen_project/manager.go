package codegen_project

import (
	"boundedinfinity/codegen/util"

	"github.com/boundedinfinity/go-commoner/optioner"
	"github.com/boundedinfinity/go-commoner/optioner/mapper"
)

type CodeGenProjectManager struct {
	All         []*CodeGenProjectProject
	source2proj mapper.Mapper[string, *CodeGenProjectProject]
	root2Proj   mapper.Mapper[string, []*CodeGenProjectProject]
	source2Root mapper.Mapper[string, string]
	root2source mapper.Mapper[string, []string]
	id2proj     mapper.Mapper[string, *CodeGenProjectProject]
	Merged      *CodeGenProjectProject
}

func Manager() *CodeGenProjectManager {
	return &CodeGenProjectManager{
		All:         make([]*CodeGenProjectProject, 0),
		source2proj: make(mapper.Mapper[string, *CodeGenProjectProject], 0),
		root2Proj:   make(mapper.Mapper[string, []*CodeGenProjectProject], 0),
		source2Root: make(mapper.Mapper[string, string], 0),
		root2source: make(mapper.Mapper[string, []string], 0),
		id2proj:     make(mapper.Mapper[string, *CodeGenProjectProject], 0),
		Merged:      NewProject(),
	}
}

func (t *CodeGenProjectManager) Register(root, source string, project *CodeGenProjectProject) {
	if project == nil {
		return
	}

	t.All = append(t.All, project)
	project.Root = root
	project.Source = source
	t.source2proj[source] = project
	t.source2Root[root] = source

	util.MapListAdd(t.root2Proj, root, project)
	util.MapListAdd(t.root2source, root, source)

	if project.Info.Id.Defined() {
		t.id2proj[project.Info.Id.Get()] = project
	}
}

func (t *CodeGenProjectManager) FindProject(id optioner.Option[string]) optioner.Option[*CodeGenProjectProject] {
	return optioner.FirstOf(t.source2proj.Get(id.Get()), t.id2proj.Get(id.Get()))
}

func (t *CodeGenProjectManager) FindRoot(id string) optioner.Option[string] {
	return optioner.FirstOf(t.source2Root.Get(id))
}

func (t *CodeGenProjectManager) TemplateFiles() []CodeGenProjectTemplateFile {
	files := make([]CodeGenProjectTemplateFile, 0)

	for _, project := range t.All {
		for _, file := range project.Templates.Files {
			files = append(files, file)
		}
	}

	return files
}
