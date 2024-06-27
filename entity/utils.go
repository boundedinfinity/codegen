package entity

import (
	"fmt"
	"time"
)

func asString(elems ...EntityType) []string {
	var results []string
	for _, elem := range elems {
		results = append(results, fmt.Sprintf("[%v,%v]", elem, elem.String()))
	}
	return results
}

func aparam[T any](data map[string]any, name string, array []T) {
	if len(array) > 0 {
		data[name] = array
	}
}

func mparam(data map[string]any, name string, sub map[string]any) {
	if len(sub) > 0 {
		data[name] = sub
	}
}

func sparam(data map[string]any, name string, val string) {
	if val != "" {
		data[name] = val
	}
}

func iparam(data map[string]any, name string, val int) {
	if val != 0 {
		data[name] = val
	}
}

func bparam(data map[string]any, name string, val bool) {
	if val {
		data[name] = val
	}
}

func tparam(data map[string]any, name string, val time.Time) {
	var zero time.Time

	if val != zero {
		data[name] = val.Unix()
	}
}

func tparams(data map[string]any, name string, elems ...time.Time) {
	var results []int64

	for _, elem := range elems {
		results = append(results, elem.Unix())
	}

	if len(results) > 0 {
		data[name] = results
	}
}
