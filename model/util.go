package model

import (
	"strings"
)

func splitDescription(s string) []string {
	var ss []string
	var splitChar string

	if s == "" {
		return ss
	}

	splitChar = DEFAULT_DESCRIPTION_SPLIT_CHARACTER
	s2 := strings.TrimSuffix(s, splitChar)
	ss = strings.Split(s2, splitChar)

	return ss
}
