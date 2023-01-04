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

func (t *CodeGenProjectManager) RegisterProject(lc *ct.CodeGenProject) error {
	t.Projects = append(t.Projects, lc)
	t.source2proj[lc.SourcePath.Get()] = lc
	t.source2Root[lc.RootPath.Get()] = lc.SourcePath.Get()

	util.MapListAdd(t.root2Proj, lc.RootPath.Get(), lc)
	util.MapListAdd(t.root2source, lc.RootPath.Get(), lc.SourcePath.Get())

	if lc.Info.Id.Defined() {
		t.id2proj[lc.Info.Id.Get()] = lc
	}

	return nil
}

func (t *CodeGenProjectManager) RegisterOperation(lc *ct.CodeGenProjectOperation) error {
	t.Operations = append(t.Operations, lc)
	return nil
}

func (t *CodeGenProjectManager) FindProject(id o.Option[string]) o.Option[*ct.CodeGenProject] {
	return o.FirstOf(t.source2proj.Get(id.Get()), t.id2proj.Get(id.Get()))
}

func (t *CodeGenProjectManager) FindRoot(id string) o.Option[string] {
	return o.FirstOf(t.source2Root.Get(id))
}

func (t *CodeGenProjectManager) TemplateFiles() []*ct.CodeGenProjectTemplateFile {
	files := make([]*ct.CodeGenProjectTemplateFile, 0)

	for _, lc := range t.Projects {
		files = append(files, lc.Templates.Files...)
	}

	return files
}
