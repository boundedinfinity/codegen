package render_context

import (
	"errors"
	"fmt"
)

var ErrExit = errors.New("")

func realErr(err error) error {
	if errors.Is(err, ErrExit) {
		return nil
	}

	return nil
}

func WalkBase(schema RenderContext, fn func(base *RenderContextBase) error) error {
	if schema == nil || schema.Base() == nil {
		return nil
	}

	if err := fn(schema.Base()); err != nil {
		return err
	}

	switch c := schema.(type) {
	case *RenderContextObject:
		for _, prop := range c.Properties {
			if err := WalkBase(prop, fn); err != nil {
				return realErr(err)
			}
		}
	case *RenderContextArray:
		if err := WalkBase(c.Items, fn); err != nil {
			return realErr(err)
		}
	case *RenderContextRef:
		if err := WalkBase(c.Ref, fn); err != nil {
			return realErr(err)
		}
	default:
		fmt.Printf("WalkBase missing: %v", schema.Base().Id)
	}

	return nil
}

func WalkConcrete(
	schema RenderContext,
	baseFn func(base *RenderContextBase) error,
	objectFn func(v *RenderContextObject) error,
	arrayFn func(v *RenderContextArray) error,
	stringFn func(v *RenderContextString) error,
	urlFn func(v *RenderContextUrl) error,
	integerFn func(v *RenderContextInteger) error,
	floatFn func(v *RenderContextFloat) error,
) error {
	if schema == nil {
		return nil
	}

	if schema.Base() != nil {
		if err := baseFn(schema.Base()); err != nil {
			return realErr(err)
		}
	}

	switch c := schema.(type) {
	case *RenderContextObject:
		if objectFn != nil {
			if err := objectFn(c); err != nil {
				return realErr(err)
			}
		}

		for _, prop := range c.Properties {
			if prop != nil {
				if err := WalkConcrete(prop, baseFn, objectFn, arrayFn, stringFn, urlFn, integerFn, floatFn); err != nil {
					return realErr(err)
				}
			}
		}
	case *RenderContextArray:
		if arrayFn != nil {
			if err := arrayFn(c); err != nil {
				return realErr(err)
			}
		}

		if err := WalkConcrete(c.Items, baseFn, objectFn, arrayFn, stringFn, urlFn, integerFn, floatFn); err != nil {
			return realErr(err)
		}
	case *RenderContextString:
		if stringFn != nil {
			if err := stringFn(c); err != nil {
				return realErr(err)
			}
		}
	case *RenderContextUrl:
		if urlFn != nil {
			if err := urlFn(c); err != nil {
				return realErr(err)
			}
		}
	case *RenderContextFloat:
		if floatFn != nil {
			if err := floatFn(c); err != nil {
				return realErr(err)
			}
		}
	case *RenderContextInteger:
		if integerFn != nil {
			if err := integerFn(c); err != nil {
				return realErr(err)
			}
		}
	default:
		if err := baseFn(schema.Base()); err != nil {
			return realErr(err)
		}
	}

	return nil
}
