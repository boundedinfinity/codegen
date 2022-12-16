package codegen_type

import (
	"boundedinfinity/codegen/codegen_type/codegen_type_id"
	"encoding/json"
	"fmt"

	"github.com/ghodss/yaml"
)

type codeGenDiscriminator struct {
	Type string `json:"type,omitempty"`
	Ref  string `json:"ref,omitempty"`
}

func UnmarshalYaml(data []byte) (CodeGenType, error) {
	if bs, err := yaml.YAMLToJSON(data); err != nil {
		return nil, err
	} else {
		return UnmarshalJson(bs)
	}
}

func UnmarshalJson(data []byte) (CodeGenType, error) {
	var c CodeGenType
	var err error

	d, err := unmarshalConcrete[codeGenDiscriminator](data)

	if err != nil {
		return nil, err
	}

	if d.Ref == "" {
		typ, err := codegen_type_id.Parse(d.Type)

		if err != nil {
			return nil, err
		}

		switch typ {
		case codegen_type_id.Array:
			c, err = unmarshalConcrete[CodeGenTypeArray](data)
		case codegen_type_id.Coordinate:
			c, err = unmarshalConcrete[CodeGenTypeCoordinate](data)
		case codegen_type_id.CreditCardNumber:
			c, err = unmarshalConcrete[CodeGenTypeCreditCardNumber](data)
		case codegen_type_id.Date:
			c, err = unmarshalConcrete[CodeGenTypeDate](data)
		case codegen_type_id.Datetime:
			c, err = unmarshalConcrete[CodeGenTypeDateTime](data)
		case codegen_type_id.Duration:
			c, err = unmarshalConcrete[CodeGenTypeDuration](data)
		case codegen_type_id.Email:
			c, err = unmarshalConcrete[CodeGenTypeEmail](data)
		case codegen_type_id.Enum:
			c, err = unmarshalConcrete[CodeGenTypeEnum](data)
		case codegen_type_id.Float:
			c, err = unmarshalConcrete[CodeGenTypeFloat](data)
		case codegen_type_id.Integer:
			c, err = unmarshalConcrete[CodeGenTypeInteger](data)
		case codegen_type_id.Ipv4:
			c, err = unmarshalConcrete[CodeGenTypeIpv4](data)
		case codegen_type_id.Ipv6:
			c, err = unmarshalConcrete[CodeGenTypeIpv6](data)
		case codegen_type_id.Mac:
			c, err = unmarshalConcrete[CodeGenTypeMac](data)
		case codegen_type_id.Object:
			c, err = unmarshalConcrete[CodeGenTypeObject](data)
		case codegen_type_id.Phone:
			c, err = unmarshalConcrete[CodeGenTypePhone](data)
		case codegen_type_id.SemanticVersion:
			// c, err = unmarshalConcrete[CanonicalSemanticVersion](data)
		case codegen_type_id.String:
			c, err = unmarshalConcrete[CodeGenTypeString](data)
		case codegen_type_id.Time:
			c, err = unmarshalConcrete[CodeGenTypeTime](data)
		case codegen_type_id.Url:
			c, err = unmarshalConcrete[CodeGenTypeUrl](data)
		case codegen_type_id.Uuid:
			c, err = unmarshalConcrete[CodeGenTypeUuid](data)
		default:
			err = fmt.Errorf("%v not implemented", typ)
		}
	} else {
		c, err = unmarshalConcrete[CodeGenTypeRef](data)
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
