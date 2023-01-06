package manager

import (
	ct "boundedinfinity/codegen/codegen_type"
	"boundedinfinity/codegen/util"

	o "github.com/boundedinfinity/go-commoner/optioner"
	"github.com/boundedinfinity/go-commoner/optioner/mapper"
)

type CodeGenProjectManager struct {
	Projects    []*ct.CodeGenProject
	Operations  []*ct.CodeGenProjectOperation
	source2proj mapper.Mapper[string, *ct.CodeGenProject]
	root2Proj   mapper.Mapper[string, []*ct.CodeGenProject]
	source2Root mapper.Mapper[string, string]
	root2source mapper.Mapper[string, []string]
	id2proj     mapper.Mapper[string, *ct.CodeGenProject]
	op2proj     mapper.Mapper[string, *ct.CodeGenProject]
	Merged      *ct.CodeGenProject
}

func ProjectManager() *CodeGenProjectManager {
	return &CodeGenProjectManager{
		Projects:    make([]*ct.CodeGenProject, 0),
		Operations:  make([]*ct.CodeGenProjectOperation, 0),
		source2proj: make(mapper.Mapper[string, *ct.CodeGenProject], 0),
		root2Proj:   make(mapper.Mapper[string, []*ct.CodeGenProject], 0),
		source2Root: make(mapper.Mapper[string, string], 0),
		root2source: make(mapper.Mapper[string, []string], 0),
		id2proj:     make(mapper.Mapper[string, *ct.CodeGenProject], 0),
		op2proj:     make(mapper.Mapper[string, *ct.CodeGenProject], 0),
		Merged:      ct.NewProject(),
	}
}

func (t *CodeGenProjectManager) RegisterProject(projects ...*ct.CodeGenProject) error {
	for _, project := range projects {
		if err := t.registerProject(project); err != nil {
			return err
		}
	}

	return nil
}

func (t *CodeGenProjectManager) registerProject(project *ct.CodeGenProject) error {
	t.Projects = append(t.Projects, project)
	t.source2proj[project.SourcePath.Get()] = project
	t.source2Root[project.RootPath.Get()] = project.SourcePath.Get()

	util.MapListAdd(t.root2Proj, project.RootPath.Get(), project)
	util.MapListAdd(t.root2source, project.RootPath.Get(), project.SourcePath.Get())

	if project.Info.Id.Defined() {
		t.id2proj[project.Info.Id.Get()] = project
	}

	return nil
}

func (t *CodeGenProjectManager) RegisterOperation(operation ct.CodeGenProjectOperation) error {
	t.Operations = append(t.Operations, &operation)
	return nil
}

func (t *CodeGenProjectManager) FindProject(id o.Option[string]) o.Option[*ct.CodeGenProject] {
	return o.FirstOf(t.source2proj.Get(id.Get()), t.id2proj.Get(id.Get()))
}

func (t *CodeGenProjectManager) FindRoot(id string) o.Option[string] {
	return o.FirstOf(t.source2Root.Get(id))
}

func (t *CodeGenProjectManager) TemplateFiles() []*ct.CodeGenProjectTypeTemplateFile {
	files := make([]*ct.CodeGenProjectTypeTemplateFile, 0)
	files = append(files, t.TemplateTypeFiles()...)
	files = append(files, t.TemplateOperationFiles()...)

	return files
}

func (t *CodeGenProjectManager) TemplateTypeFiles() []*ct.CodeGenProjectTypeTemplateFile {
	files := make([]*ct.CodeGenProjectTypeTemplateFile, 0)

	for _, project := range t.Projects {
		files = append(files, project.Templates.Types...)
	}

	return files
}

func (t *CodeGenProjectManager) TemplateOperationFiles() []*ct.CodeGenProjectTypeTemplateFile {
	files := make([]*ct.CodeGenProjectTypeTemplateFile, 0)

	for _, project := range t.Projects {
		files = append(files, project.Templates.Operations...)
	}

	return files
}
