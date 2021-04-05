package util

func StrPrt(v string) *string {
	return &v
}

func StrIsEmpty(v *string) bool {
	if v == nil || *v == "" {
		return true
	}

	return false
}

func StrSliceIsEmpty(v []string) bool {
	if v == nil {
		return true
	}

	for _, s := range v {
		if s != "" {
			return false
		}
	}

	return true
}

func BoolPrt(v bool) *bool {
	return &v
}

func BoolIsEmpty(v *bool) bool {
	if v == nil {
		return true
	}

	return false
}
