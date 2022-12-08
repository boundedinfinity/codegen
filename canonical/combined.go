package canonical

import (
	"fmt"

	o "github.com/boundedinfinity/go-commoner/optioner"
	"github.com/boundedinfinity/go-commoner/optioner/mapper"
)

type CanonicalCombined struct {
	typeMap mapper.Mapper[string, Canonical]
}

func NewCombinded() *CanonicalCombined {
	return &CanonicalCombined{
		typeMap: make(mapper.Mapper[string, Canonical], 0),
	}
}

func (t CanonicalCombined) All() []Canonical {
	return t.typeMap.Values().Get()
}

func (t CanonicalCombined) Has(id string) bool {
	return t.typeMap.Has(id)
}

func (t CanonicalCombined) Find(id string) o.Option[Canonical] {
	return o.FirstOf(t.FindById(id))
}

func (t CanonicalCombined) FindById(id string) o.Option[Canonical] {
	return t.typeMap.Get(id)
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

func (t *CanonicalCombined) Register(schema Canonical) error {
	id := t.Id(schema)

	if id.Empty() {
		return nil
	}

	if t.typeMap.Has(id.Get()) {
		fmt.Printf("already contains %v", id.Get())
	}

	t.typeMap[id.Get()] = schema

	return nil
}
