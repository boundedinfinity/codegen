// Package name contains the enumeration of kind names
package name

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
)

type KindName string

const (
	unknown  KindName = "unknown"
	Array    KindName = "array"
	Boolean  KindName = "boolean"
	Date     KindName = "date"
	DateTime KindName = "date-time"
	Duration KindName = "duration"
	Enum     KindName = "enumeration"
	Float    KindName = "float"
	Integer  KindName = "integer"
	Object   KindName = "object"
	Ref      KindName = "ref"
	String   KindName = "string"
	Time     KindName = "time"
	Union    KindName = "union"
)

func (this KindName) String() string {
	return string(this)
}

func (this *KindName) MarshalJSON() ([]byte, error) {
	if *this == "" {
		return nil, fmt.Errorf(`%w : "<empty>"`, ErrKindInvalid)
	}

	return json.Marshal(*this)
}

func (this *KindName) UnmarshalJSON(data []byte) error {
	var raw string
	var found KindName

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	for _, name := range names {
		if string(name) == raw {
			found = name
			break
		}
	}

	if found == unknown {
		return fmt.Errorf(`%w : must be one of %s`,
			ErrKindInvalid,
			stringer.Join(", ", names...),
		)
	}

	return nil
}

var (
	ErrKindInvalid = errors.New("invalid kind")

	names = []KindName{
		Array,
		Boolean,
		Date,
		DateTime,
		Duration,
		Enum,
		Float,
		Integer,
		Object,
		Ref,
		String,
		Time,
		Union,
	}
)

func Names() []KindName {
	return names
}
