package canonical

import (
	"fmt"

	o "github.com/boundedinfinity/go-commoner/optioner"
	"github.com/boundedinfinity/go-commoner/optioner/mapper"
)

type CanonicalCombined struct {
	idMap   mapper.Mapper[string, Canonical]
	pathMap mapper.Mapper[string, Canonical]
	id2path mapper.Mapper[string, string]
	path2id mapper.Mapper[string, string]
}

func NewCombinded() *CanonicalCombined {
	return &CanonicalCombined{
		idMap:   make(mapper.Mapper[string, Canonical], 0),
		pathMap: make(mapper.Mapper[string, Canonical], 0),
		id2path: make(mapper.Mapper[string, string], 0),
		path2id: make(mapper.Mapper[string, string], 0),
	}
}

func (t CanonicalCombined) All() []Canonical {
	return t.idMap.Values().Get()
}

func (t CanonicalCombined) Has(id string) bool {
	return t.idMap.Has(id)
}

func (t CanonicalCombined) Find(id o.Option[string]) o.Option[Canonical] {
	if id.Empty() {
		return o.None[Canonical]()
	}

	return o.FirstOf(t.idMap.Get(id.Get()), t.pathMap.Get(id.Get()))
}

func (t CanonicalCombined) FindSource(id o.Option[string]) o.Option[string] {
	if id.Empty() {
		return o.None[string]()
	}

	c := t.Find(id)

	if c.Empty() || c.Get().SchemaId().Empty() {
		return o.None[string]()
	}

	return t.id2path.Get(c.Get().SchemaId().Get())
}

func (t CanonicalCombined) Id(schema Canonical) o.Option[string] {
	var id o.Option[string]

	switch v := schema.(type) {
	case CanonicalArray:
		id = v.Id
	case *CanonicalArray:
		id = v.Id
	case CanonicalCoordinate:
		id = v.Id
	case *CanonicalCoordinate:
		id = v.Id
	case CanonicalCreditCardNumber:
		id = v.Id
	case *CanonicalCreditCardNumber:
		id = v.Id
	case CanonicalDate:
		id = v.Id
	case *CanonicalDate:
		id = v.Id
	case CanonicalDateTime:
		id = v.Id
	case *CanonicalDateTime:
		id = v.Id
	case CanonicalDuration:
		id = v.Id
	case *CanonicalDuration:
		id = v.Id
	case CanonicalEmail:
		id = v.Id
	case CanonicalEnum:
		id = v.Id
	case *CanonicalEnum:
		id = v.Id
	case CanonicalFloat:
		id = v.Id
	case *CanonicalFloat:
		id = v.Id
	case CanonicalInteger:
		id = v.Id
	case *CanonicalInteger:
		id = v.Id
	case CanonicalIpv4:
		id = v.Id
	case *CanonicalIpv4:
		id = v.Id
	case CanonicalIpv6:
		id = v.Id
	case *CanonicalIpv6:
		id = v.Id
	case CanonicalMac:
		id = v.Id
	case *CanonicalMac:
		id = v.Id
	case CanonicalObject:
		id = v.Id
	case *CanonicalObject:
		id = v.Id
	case CanonicalPhone:
		id = v.Id
	case *CanonicalPhone:
		id = v.Id
	case CanonicalString:
		id = v.Id
	case *CanonicalString:
		id = v.Id
	case CanonicalTime:
		id = v.Id
	case *CanonicalTime:
		id = v.Id
	case CanonicalUrl:
		id = v.Id
	case *CanonicalUrl:
		id = v.Id
	case CanonicalUuid:
		id = v.Id
	case *CanonicalUuid:
		id = v.Id
	}

	return id
}

func (t *CanonicalCombined) Register(schema Canonical, path o.Option[string]) error {
	id := t.Id(schema)

	if id.Defined() {
		if t.idMap.Has(id.Get()) {
			fmt.Printf("already contains %v", id.Get())
		}

		t.idMap[id.Get()] = schema
	}

	if path.Defined() {
		t.pathMap[path.Get()] = schema
	}

	if id.Defined() && path.Defined() {
		t.id2path[id.Get()] = path.Get()
		t.path2id[path.Get()] = id.Get()
	}

	return nil
}
