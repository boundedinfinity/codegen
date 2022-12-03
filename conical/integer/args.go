package integer

import "boundedinfinity/codegen/conical"

func New(args ...Arg) *conical.ConicalInteger {
	t := &conical.ConicalInteger{}

	return t
}

type Arg func(*conical.ConicalInteger)

// func Namespace(v string) func(*conical.ConicalInteger) {
// 	return func(t *conical.ConicalInteger) {
// 		t.Namespace = v
// 	}
// }

// func Name(v string) func(*conical.ConicalInteger) {
// 	return func(t *conical.ConicalInteger) {
// 		t.Name = v
// 	}
// }

// func Version(v string) func(*conical.ConicalInteger) {
// 	return func(t *conical.ConicalInteger) {
// 		t.Version = v
// 	}
// }

// func Min(v int) func(*conical.ConicalInteger) {
// 	return func(t *conical.ConicalInteger) {
// 		t.Minimum = v
// 	}
// }

// func Max(v int) func(*conical.ConicalInteger) {
// 	return func(t *conical.ConicalInteger) {
// 		t.Maximum = v
// 	}
// }
