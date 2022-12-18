package loader_context

import (
	"boundedinfinity/codegen/codegen_type/codegen_type_id"
	"boundedinfinity/codegen/template_type"

	"github.com/boundedinfinity/go-commoner/optioner/mapper"
)

type CodeGenTemplateManager struct {
	path2project         mapper.Mapper[string, TemplateLoaderContext]
	root2project         mapper.Mapper[string, []TemplateLoaderContext]
	templateType2Context mapper.Mapper[template_type.TemplateType, []TemplateLoaderContext]
	typeId2Context       mapper.Mapper[codegen_type_id.CodgenTypeId, []TemplateLoaderContext]
}

func TemplateManager() *CodeGenTemplateManager {
	return &CodeGenTemplateManager{
		path2project:         make(mapper.Mapper[string, TemplateLoaderContext]),
		root2project:         make(mapper.Mapper[string, []TemplateLoaderContext]),
		templateType2Context: make(mapper.Mapper[template_type.TemplateType, []TemplateLoaderContext]),
		typeId2Context:       make(mapper.Mapper[codegen_type_id.CodgenTypeId, []TemplateLoaderContext]),
	}
}

func (t *CodeGenTemplateManager) Register(lc *ProjectLoaderContext) {
	// t.All = append(t.All, lc)
	// t.source2proj[lc.Source] = lc
	// t.source2Root[lc.Root] = lc.Source

	// util.MapListAdd(t.root2Proj, lc.Root, lc)
	// util.MapListAdd(t.root2source, lc.Root, lc.Source)

	// if lc.CodeGenProjectProject.Info.Id.Defined() {
	// 	t.id2proj[lc.CodeGenProjectProject.Info.Id.Get()] = lc
	// }
}
