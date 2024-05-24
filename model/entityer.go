package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

type entityerDescriptor struct {
	name  string
	value any
	fn    func() any
}

type Entityer struct {
	items     []entityerDescriptor
	OmitEmpty bool
	OmitZero  bool
	OmitNil   bool
}

func (t *Entityer) Value(name string, value any) error {
	if name == "" {
		return ErrJsonerMissingName
	}

	kind := reflect.ValueOf(value).Kind()

	switch kind {
	case reflect.Pointer:
		fmt.Println(kind)
	default:
		return ErrNJsonerotPointer
	}

	t.items = append(t.items, entityerDescriptor{
		name:  name,
		value: value,
	})

	return nil
}

func (t *Entityer) Fn(name string, fn func() any) error {
	if name == "" {
		return ErrJsonerMissingName
	}

	if fn != nil {
		t.items = append(t.items, entityerDescriptor{
			name: name,
			fn:   fn,
		})
	}

	return nil
}

var ErrJsonerMultipleDefinitions = errors.New("multiple definitions")
var ErrJsonerMissingName = errors.New("missing name")
var ErrNJsonerotPointer = errors.New("not pointer")

func (t Entityer) MarshalJSON() ([]byte, error) {
	m := map[string]any{}

	for _, item := range t.items {
		if item.value != nil {
			m[item.name] = item.value
		}

		if item.fn != nil {
			v := item.fn()

			if v != nil {
				m[item.name] = v
			}
		}
	}

	return json.Marshal(m)
}
