package manager

import (
	ct "boundedinfinity/codegen/codegen_type"
	"boundedinfinity/codegen/codegen_type/codegen_type_id"
	"boundedinfinity/codegen/codegen_type/template_type"
	"boundedinfinity/codegen/util"
	"fmt"

	"github.com/boundedinfinity/go-commoner/optioner"
	"github.com/boundedinfinity/go-commoner/optioner/mapper"
)

type CodeGenTemplateManager struct {
	All             []*ct.TemplateMeta
	source2template mapper.Mapper[string, *ct.TemplateMeta]
	root2template   mapper.Mapper[string, []*ct.TemplateMeta]
	source2root     mapper.Mapper[string, string]
	root2source     mapper.Mapper[string, []string]
	tt2template     mapper.Mapper[template_type.TemplateType, []*ct.TemplateMeta]
	tId2template    mapper.Mapper[codegen_type_id.CodgenTypeId, []*ct.TemplateMeta]
}

func TemplateManager() *CodeGenTemplateManager {
	return &CodeGenTemplateManager{
		All:             make([]*ct.TemplateMeta, 0),
		source2template: make(mapper.Mapper[string, *ct.TemplateMeta]),
		source2root:     make(mapper.Mapper[string, string]),
		root2template:   make(mapper.Mapper[string, []*ct.TemplateMeta]),
		root2source:     make(mapper.Mapper[string, []string]),
		tt2template:     make(mapper.Mapper[template_type.TemplateType, []*ct.TemplateMeta]),
		tId2template:    make(mapper.Mapper[codegen_type_id.CodgenTypeId, []*ct.TemplateMeta]),
	}
}

func (t *CodeGenTemplateManager) Register(lc *ct.TemplateMeta) {
	t.All = append(t.All, lc)
	t.source2template[lc.SourcePath.Get()] = lc
	t.source2root[lc.SourcePath.Get()] = lc.RootPath.Get()
	util.MapListAdd(t.root2template, lc.RootPath.Get(), lc)
	util.MapListAdd(t.root2source, lc.RootPath.Get(), lc.SourcePath.Get())
	util.MapListAdd(t.tt2template, lc.TemplateType, lc)
	util.MapListAdd(t.tId2template, lc.TypeId, lc)
}

func (t *CodeGenTemplateManager) Find(id any) optioner.Option[[]*ct.TemplateMeta] {
	s := fmt.Sprintf("%v", id)

	return optioner.FirstOf(
		t.tId2template.Get(codegen_type_id.CodgenTypeId(s)),
		t.tt2template.Get(template_type.TemplateType(s)),
	)
}
