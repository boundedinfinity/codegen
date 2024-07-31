package {{ typePackage . }}{{ $tname := typeName . }}

import (
    "database/sql/driver"
    "errors"
)

func (t {{ $tname }}) Value() (driver.Value, error) {
    return int32(t), nil
}

func (t *{{ $tname }}) Scan(value interface{}) error {
	if value == nil {
		*t = {{ $tname }}(0)
		return nil
	}

	if bv, err := driver.Int32.ConvertValue(value); err == nil {
		if v, ok := bv.({{ langType . }}); ok {
			*t = {{ $tname }}(v)
			return nil
		}
	}
	
	return errors.New("failed to scan {{ $tname }}")
}
