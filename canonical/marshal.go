package canonical

import (
	"boundedinfinity/codegen/canonical/canonical_type"
	"encoding/json"
	"fmt"

	"github.com/ghodss/yaml"
)

type canonicalDiscriminator struct {
	Type string `json:"type,omitempty"`
	Ref  string `json:"ref,omitempty"`
}

func UnmarshalCanonicalSchemaYaml(data []byte) (Canonical, error) {
	if bs, err := yaml.YAMLToJSON(data); err != nil {
		return nil, err
	} else {
		return UnmarshalCanonicalSchemaJson(bs)
	}
}

func UnmarshalCanonicalSchemaJson(data []byte) (Canonical, error) {
	var c Canonical
	var err error

	d, err := unmarshalConcrete[canonicalDiscriminator](data)

	if err != nil {
		return nil, err
	}

	if d.Ref == "" {
		typ, err := canonical_type.Parse(d.Type)

		if err != nil {
			return nil, err
		}

		switch typ {
		case canonical_type.Array:
			c, err = unmarshalConcrete[CanonicalArray](data)
		case canonical_type.Date:
			c, err = unmarshalConcrete[CanonicalDate](data)
		case canonical_type.Datetime:
			c, err = unmarshalConcrete[CanonicalDateTime](data)
		case canonical_type.Duration:
			c, err = unmarshalConcrete[CanonicalDuration](data)
		case canonical_type.Enum:
			c, err = unmarshalConcrete[CanonicalEnum](data)
		case canonical_type.Float:
			c, err = unmarshalConcrete[CanonicalFloat](data)
		case canonical_type.Integer:
			c, err = unmarshalConcrete[CanonicalInteger](data)
		case canonical_type.Object:
			c, err = unmarshalConcrete[CanonicalObject](data)
		case canonical_type.String:
			c, err = unmarshalConcrete[CanonicalString](data)
		case canonical_type.Uuid:
			c, err = unmarshalConcrete[CanonicalUuid](data)
		default:
			err = fmt.Errorf("%v not implemented", typ)
		}
	} else {
		c, err = unmarshalConcrete[CanonicalRef](data)
	}

	return c, err
}

func unmarshalConcrete[T any](data []byte) (*T, error) {
	var t T

	if err := json.Unmarshal(data, &t); err != nil {
		return &t, err
	}

	return &t, nil
}
