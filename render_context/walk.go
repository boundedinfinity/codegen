package render_context

func WalkBase(schema RenderContext, fn func(base *RenderContextBase) error) error {
	switch c := schema.(type) {
	case *RenderContextObject:
		if err := fn(schema.Base()); err != nil {
			return err
		}

		for _, prop := range c.Properties {
			if err := WalkBase(prop, fn); err != nil {
				return err
			}
		}
	case *RenderContextArray:
		if err := fn(schema.Base()); err != nil {
			return err
		}

		if err := WalkBase(c.Items, fn); err != nil {
			return err
		}
	case *RenderContextRef:
		if err := fn(schema.Base()); err != nil {
			return err
		}

		if err := WalkBase(c.Ref, fn); err != nil {
			return err
		}
	default:
		if err := fn(schema.Base()); err != nil {
			return err
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
			return err
		}

		if err := objectFn(c); err != nil {
			return nil
		}

		for _, prop := range c.Properties {
			if err := WalkConcrete(prop, baseFn, objectFn, arrayFn); err != nil {
				return nil
			}
		}
	case *RenderContextArray:
		if err := baseFn(schema.Base()); err != nil {
			return err
		}
		if err := arrayFn(c); err != nil {
			return err
		}

		if err := WalkConcrete(c.Items, baseFn, objectFn, arrayFn); err != nil {
			return err
		}
	default:
		if err := baseFn(schema.Base()); err != nil {
			return err
		}
	}

	return nil
}
