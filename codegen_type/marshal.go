package codegen_type

import (
	"boundedinfinity/codegen/codegen_type/codegen_type_id"
	"boundedinfinity/codegen/model"
	"encoding/json"
	"fmt"

	"github.com/boundedinfinity/go-commoner/optioner"
	"github.com/ghodss/yaml"
)

type codeGenDiscriminator struct {
	model.Source
	Type optioner.Option[string] `json:"type,omitempty"`
	Ref  optioner.Option[string] `json:"ref,omitempty"`
}

func UnmarshalYaml(data []byte, v *CodeGenType) error {
	if bs, err := yaml.YAMLToJSON(data); err != nil {
		return err
	} else {
		return UnmarshalJson(bs, v)
	}
}

func UnmarshalJson(data []byte, v *CodeGenType) error {
	var c model.Type
	var err error
	var typ codegen_type_id.CodgenTypeId

	d, err := unmarshalConcrete[codeGenDiscriminator](data)

	if err != nil {
		return err
	}

	switch {
	case d.Type.Defined():
		typ, err = codegen_type_id.Parse(d.Type.Get())

		if err != nil {
			return err
		}

		switch typ {
		case codegen_type_id.Array:
			c, err = unmarshalConcrete[model.Array](data)
		// case codegen_type_id.Coordinate:
		// 	c, err = unmarshalConcrete[CodeGenTypeCoordinate](data)
		// case codegen_type_id.CreditCardNumber:
		// 	c, err = unmarshalConcrete[CodeGenTypeCreditCardNumber](data)
		// case codegen_type_id.Date:
		// 	c, err = unmarshalConcrete[CodeGenTypeDate](data)
		// case codegen_type_id.Datetime:
		// 	c, err = unmarshalConcrete[CodeGenTypeDateTime](data)
		// case codegen_type_id.Duration:
		// 	c, err = unmarshalConcrete[CodeGenTypeDuration](data)
		// case codegen_type_id.Email:
		// 	c, err = unmarshalConcrete[CodeGenTypeEmail](data)
		// case codegen_type_id.Enum:
		// 	c, err = unmarshalConcrete[CodeGenTypeEnum](data)
		case codegen_type_id.Float:
			c, err = unmarshalConcrete[model.Float](data)
		case codegen_type_id.Integer:
			c, err = unmarshalConcrete[model.Int](data)
		// case codegen_type_id.Ipv4:
		// 	c, err = unmarshalConcrete[CodeGenTypeIpv4](data)
		// case codegen_type_id.Ipv6:
		// 	c, err = unmarshalConcrete[CodeGenTypeIpv6](data)
		// case codegen_type_id.Mac:
		// 	c, err = unmarshalConcrete[CodeGenTypeMac](data)
		case codegen_type_id.Object:
			c, err = unmarshalConcrete[model.Object](data)
		// case codegen_type_id.Phone:
		// 	c, err = unmarshalConcrete[CodeGenTypePhone](data)
		case codegen_type_id.SemanticVersion:
			// c, err = unmarshalConcrete[CanonicalSemanticVersion](data)
		case codegen_type_id.String:
			c, err = unmarshalConcrete[model.String](data)
			// case codegen_type_id.Time:
			// 	c, err = unmarshalConcrete[CodeGenTypeTime](data)
			// case codegen_type_id.Url:
			// 	c, err = unmarshalConcrete[CodeGenTypeUrl](data)
			// case codegen_type_id.Uuid:
			// 	c, err = unmarshalConcrete[CodeGenTypeUuid](data)
		}
	case d.Ref.Defined():
		c, err = unmarshalConcrete[model.Ref](data)
	case d.SourcePath.Defined() || d.RootPath.Defined():
		c = &CodeGenTypePath{
			SourceMeta: d.SourceMeta,
		}
	default:
		err = fmt.Errorf("%v not implemented", typ)
	}

	*v = c
	return err
}

func unmarshalConcrete[T any](data []byte) (*T, error) {
	var t T

	if err := json.Unmarshal(data, &t); err != nil {
		return &t, err
	}

	return &t, nil
}
