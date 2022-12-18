package loader_context

import (
	cp "boundedinfinity/codegen/codegen_project"
	"boundedinfinity/codegen/util"

	o "github.com/boundedinfinity/go-commoner/optioner"
	"github.com/boundedinfinity/go-commoner/optioner/mapper"
)

type CodeGenProjectManager struct {
	All         []*ProjectLoaderContext
	source2proj mapper.Mapper[string, *ProjectLoaderContext]
	root2Proj   mapper.Mapper[string, []*ProjectLoaderContext]
	source2Root mapper.Mapper[string, string]
	root2source mapper.Mapper[string, []string]
	id2proj     mapper.Mapper[string, *ProjectLoaderContext]
	Merged      *cp.CodeGenProjectProject
}

func ProjectManager() *CodeGenProjectManager {
	return &CodeGenProjectManager{
		All:         make([]*ProjectLoaderContext, 0),
		source2proj: make(mapper.Mapper[string, *ProjectLoaderContext], 0),
		root2Proj:   make(mapper.Mapper[string, []*ProjectLoaderContext], 0),
		source2Root: make(mapper.Mapper[string, string], 0),
		root2source: make(mapper.Mapper[string, []string], 0),
		id2proj:     make(mapper.Mapper[string, *ProjectLoaderContext], 0),
		Merged:      cp.NewProject(),
	}
}

func (t *CodeGenProjectManager) Register(lc *ProjectLoaderContext) {
	t.All = append(t.All, lc)
	t.source2proj[lc.FileInfo.Source] = lc
	t.source2Root[lc.FileInfo.Root] = lc.FileInfo.Source

	util.MapListAdd(t.root2Proj, lc.FileInfo.Root, lc)
	util.MapListAdd(t.root2source, lc.FileInfo.Root, lc.FileInfo.Source)

	if lc.Project.Info.Id.Defined() {
		t.id2proj[lc.Project.Info.Id.Get()] = lc
	}
}

func (t *CodeGenProjectManager) FindProject(id o.Option[string]) o.Option[*ProjectLoaderContext] {
	return o.FirstOf(t.source2proj.Get(id.Get()), t.id2proj.Get(id.Get()))
}

func (t *CodeGenProjectManager) FindRoot(id string) o.Option[string] {
	return o.FirstOf(t.source2Root.Get(id))
}

func (t *CodeGenProjectManager) TemplateFiles() []*cp.CodeGenProjectTemplateFile {
	files := make([]*cp.CodeGenProjectTemplateFile, 0)

	for _, lc := range t.All {
		for _, file := range lc.Project.Templates.Files {
			files = append(files, file)
		}
	}

	return files
}
