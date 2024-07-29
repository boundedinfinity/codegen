package validation

// var ErrIntegerLessThanMin = errors.New("length is less than min value")

// func Min[T ~string](name string, n int) func(v T) error {
// 	return func(v T) error {
// 		if len(v) < n {
// 			return fmt.Errorf("%s value %s %w of %d", name, v, ErrStringLessThanMin, n)
// 		}

// 		return nil
// 	}
// }

// var ErrStringGreaterThanMax = errors.New("length is greater than max value")

// func Max[T ~string](name string, n int) func(v T) error {
// 	return func(v T) error {
// 		if len(v) < n {
// 			return fmt.Errorf("%s value %s %w of %d", name, v, ErrStringGreaterThanMax, n)
// 		}

// 		return nil
// 	}
// }

// var ErrStringDoesNotMatchPattern = errors.New("does not match pattern")

// func Regex[T ~string](name string, pattern string) func(v T) error {
// 	regex := regexp.MustCompile(pattern)

// 	return func(v T) error {
// 		if !regex.MatchString(string(v)) {
// 			return fmt.Errorf("%s value %s %w of %s", name, v, ErrStringDoesNotMatchPattern, pattern)
// 		}

// 		return nil
// 	}
// }

// var ErrStringEmpty = errors.New("is empty")

// func NotEmpty[T ~string](name string, pattern string) func(v T) error {
// 	return func(v T) error {
// 		if v == "" {
// 			return fmt.Errorf("%s %w of %s", name, ErrStringEmpty, pattern)
// 		}

// 		return nil
// 	}
// }

// var ErrStringNotUpperCase = errors.New("is not upper cased")

// func NotUpperCase[T ~string](name string) func(v T) error {
// 	return func(v T) error {
// 		if stringer.Capitalize(v) != string(v) {
// 			return fmt.Errorf("%s value v %w", name, ErrStringNotUpperCase)
// 		}

// 		return nil
// 	}
// }

// var ErrStringNotLowerCase = errors.New("is not lower cased")

// func NotLowerCase[T ~string](name string) func(v T) error {
// 	return func(v T) error {
// 		if stringer.Capitalize(v) != string(v) {
// 			return fmt.Errorf("%s value v %w", name, ErrStringNotLowerCase)
// 		}

// 		return nil
// 	}
// }
