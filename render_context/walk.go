package render_context

import "errors"

var ErrExit = errors.New("")

func realErr(err error) error {
	if errors.Is(err, ErrExit) {
		return nil
	}

	return nil
}

func WalkBase(schema RenderContext, fn func(base *RenderContextBase) error) error {
	switch c := schema.(type) {
	case *RenderContextObject:
		if err := fn(schema.Base()); err != nil {
			return realErr(err)
		}

		for _, prop := range c.Properties {
			if err := WalkBase(prop, fn); err != nil {
				return realErr(err)
			}
		}
	case *RenderContextArray:
		if err := fn(schema.Base()); err != nil {
			return realErr(err)
		}

		if err := WalkBase(c.Items, fn); err != nil {
			return realErr(err)
		}
	case *RenderContextRef:
		if err := fn(schema.Base()); err != nil {
			return realErr(err)
		}

		if err := WalkBase(c.Ref, fn); err != nil {
			return realErr(err)
		}
	default:
		if err := fn(schema.Base()); err != nil {
			return realErr(err)
		}
	}

	return nil
}

func WalkConcrete(
	schema RenderContext,
	baseFn func(base *RenderContextBase) error,
	objectFn func(v *RenderContextObject) error,
	arrayFn func(v *RenderContextArray) error,
) error {
	switch c := schema.(type) {
	case *RenderContextObject:
		if err := baseFn(schema.Base()); err != nil {
			return realErr(err)
		}

		if err := objectFn(c); err != nil {
			return realErr(err)
		}

		for _, prop := range c.Properties {
			if err := WalkConcrete(prop, baseFn, objectFn, arrayFn); err != nil {
				return realErr(err)
			}
		}
	case *RenderContextArray:
		if err := baseFn(schema.Base()); err != nil {
			return realErr(err)
		}
		if err := arrayFn(c); err != nil {
			return realErr(err)
		}

		if err := WalkConcrete(c.Items, baseFn, objectFn, arrayFn); err != nil {
			return realErr(err)
		}
	default:
		if err := baseFn(schema.Base()); err != nil {
			return realErr(err)
		}
	}

	return nil
}
