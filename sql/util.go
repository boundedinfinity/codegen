package sql

func setAndReturn[B any, V comparable](builder B, opt *V, value V) B {
	*opt = value
	return builder
}
