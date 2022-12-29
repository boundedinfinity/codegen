package manager

import (
	ct "boundedinfinity/codegen/codegen_type"
	"boundedinfinity/codegen/util"

	o "github.com/boundedinfinity/go-commoner/optioner"
	"github.com/boundedinfinity/go-commoner/optioner/mapper"
)

type CodeGenProjectManager struct {
	Projects    []*ct.ProjectLoaderContext
	Operations  []*ct.OperationLoaderContext
	source2proj mapper.Mapper[string, *ct.ProjectLoaderContext]
	root2Proj   mapper.Mapper[string, []*ct.ProjectLoaderContext]
	source2Root mapper.Mapper[string, string]
	root2source mapper.Mapper[string, []string]
	id2proj     mapper.Mapper[string, *ct.ProjectLoaderContext]
	op2proj     mapper.Mapper[string, *ct.ProjectLoaderContext]
	Merged      *ct.CodeGenProjectProject
}

func ProjectManager() *CodeGenProjectManager {
	return &CodeGenProjectManager{
		Projects:    make([]*ct.ProjectLoaderContext, 0),
		Operations:  make([]*ct.OperationLoaderContext, 0),
		source2proj: make(mapper.Mapper[string, *ct.ProjectLoaderContext], 0),
		root2Proj:   make(mapper.Mapper[string, []*ct.ProjectLoaderContext], 0),
		source2Root: make(mapper.Mapper[string, string], 0),
		root2source: make(mapper.Mapper[string, []string], 0),
		id2proj:     make(mapper.Mapper[string, *ct.ProjectLoaderContext], 0),
		op2proj:     make(mapper.Mapper[string, *ct.ProjectLoaderContext], 0),
		Merged:      ct.NewProject(),
	}
}

func (t *CodeGenProjectManager) RegisterProject(lc *ct.ProjectLoaderContext) error {
	t.Projects = append(t.Projects, lc)
	t.source2proj[lc.FileInfo.Source] = lc
	t.source2Root[lc.FileInfo.Root] = lc.FileInfo.Source

	util.MapListAdd(t.root2Proj, lc.FileInfo.Root, lc)
	util.MapListAdd(t.root2source, lc.FileInfo.Root, lc.FileInfo.Source)

	if lc.Project.Info.Id.Defined() {
		t.id2proj[lc.Project.Info.Id.Get()] = lc
	}

	return nil
}

func (t *CodeGenProjectManager) RegisterOperation(lc *ct.OperationLoaderContext) error {
	t.Operations = append(t.Operations, lc)
	return nil
}

func (t *CodeGenProjectManager) FindProject(id o.Option[string]) o.Option[*ct.ProjectLoaderContext] {
	return o.FirstOf(t.source2proj.Get(id.Get()), t.id2proj.Get(id.Get()))
}

func (t *CodeGenProjectManager) FindRoot(id string) o.Option[string] {
	return o.FirstOf(t.source2Root.Get(id))
}

func (t *CodeGenProjectManager) TemplateFiles() []*ct.CodeGenProjectTemplateFile {
	files := make([]*ct.CodeGenProjectTemplateFile, 0)

	for _, lc := range t.Projects {
		files = append(files, lc.Project.Templates.Files...)
	}

	return files
}
