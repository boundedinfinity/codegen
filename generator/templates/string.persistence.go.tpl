package {{ typePackage . }}{{ $tname := typeName . }}

import (
    "database/sql/driver"
    "errors"
)

func (t {{ $tname }}) Value() (driver.Value, error) {
    return string(t), nil
}

func (t *{{ $tname }}) Scan(value interface{}) error {
	if value == nil {
		*t = {{ $tname }}("")
		return nil
	}

	if bv, err := driver.String.ConvertValue(value); err == nil {
		if v, ok := bv.(string); ok {
			*t = {{ $tname }}(v)
			return nil
		}
	}
	
	return errors.New("failed to scan {{ $tname }}")
}
