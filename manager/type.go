package manager

import (
	ct "boundedinfinity/codegen/codegen_type"
	"boundedinfinity/codegen/util"

	o "github.com/boundedinfinity/go-commoner/optioner"
	"github.com/boundedinfinity/go-commoner/optioner/mapper"
)

type CodeGenTypeManager struct {
	order     []ct.CodeGenType
	id2Type   mapper.Mapper[string, ct.CodeGenType]
	path2Type mapper.Mapper[string, ct.CodeGenType]
	id2path   mapper.Mapper[string, string]
	path2id   mapper.Mapper[string, string]
	root2path mapper.Mapper[string, []string]
	path2root mapper.Mapper[string, string]
	root2Type mapper.Mapper[string, []ct.CodeGenType]
}

func TypeManager() *CodeGenTypeManager {
	return &CodeGenTypeManager{
		order:     make([]ct.CodeGenType, 0),
		id2Type:   make(mapper.Mapper[string, ct.CodeGenType], 0),
		path2Type: make(mapper.Mapper[string, ct.CodeGenType], 0),
		root2Type: make(mapper.Mapper[string, []ct.CodeGenType], 0),
		id2path:   make(mapper.Mapper[string, string], 0),
		path2id:   make(mapper.Mapper[string, string], 0),
		root2path: make(mapper.Mapper[string, []string], 0),
		path2root: make(mapper.Mapper[string, string], 0),
	}
}

func (t *CodeGenTypeManager) Register(lc ct.CodeGenType) error {
	if lc.Base().Id.Defined() {
		t.id2Type[lc.Base().Id.Get()] = lc
		t.id2path[lc.Base().Id.Get()] = lc.Source().SourcePath.Get()
		t.path2id[lc.Source().SourcePath.Get()] = lc.Base().Id.Get()
	}

	t.path2Type[lc.Source().SourcePath.Get()] = lc
	t.path2root[lc.Source().SourcePath.Get()] = lc.Source().RootPath.Get()
	util.MapListAdd(t.root2Type, lc.Source().RootPath.Get(), lc)
	util.MapListAdd(t.root2path, lc.Source().RootPath.Get(), lc.Source().SourcePath.Get())

	t.order = append(t.order, lc)

	return nil
}

func (t CodeGenTypeManager) All() []ct.CodeGenType {
	return t.order
}

func (t CodeGenTypeManager) Has(id string) bool {
	return t.id2Type.Has(id)
}

func (t CodeGenTypeManager) Find(id o.Option[string]) o.Option[ct.CodeGenType] {
	a := t.id2Type.Get(id.Get())
	b := t.path2Type.Get(id.Get())
	return o.FirstOf(a, b)
}

func (t CodeGenTypeManager) Resolve(schema o.Option[ct.CodeGenType]) o.Option[ct.CodeGenType] {
	if schema.Empty() {
		return schema
	}

	switch c := schema.Get().(type) {
	case *ct.CodeGenTypeRef:
		return t.Find(c.Ref)
	case *ct.CodeGenTypeArray:
		return t.Resolve(o.Some(c.Items))
	default:
		return t.Find(schema.Get().Base().Id)
	}
}

func (t CodeGenTypeManager) ResolveRef(typ ct.CodeGenType) error {
	if typ == nil {
		return ct.ErrCodeGenRefNotFound
	}

	switch c := typ.(type) {
	case *ct.CodeGenTypeRef:
		found := t.Find(c.Ref)

		if found.Defined() {
			c.Resolved = found.Get()
		} else {
			return ct.ErrCodeGenRefNotFoundv(typ)
		}
	case *ct.CodeGenTypeArray:
		return t.ResolveRef(c.Items)
	case *ct.CodeGenTypeObject:
		for _, property := range c.Properties {
			if err := t.ResolveRef(property); err != nil {
				return err
			}
		}
	}

	return nil
}

func (t CodeGenTypeManager) FindSource(id o.Option[string]) o.Option[string] {
	if id.Empty() {
		return o.None[string]()
	}

	c := t.Find(id)

	if c.Empty() || c.Get().Base().Id.Empty() {
		return o.None[string]()
	}

	return t.id2path.Get(c.Get().Base().Id.Get())
}
