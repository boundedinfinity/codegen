package render_context

import (
	"errors"
)

var ErrExit = errors.New("")

func realErr(err error) error {
	if errors.Is(err, ErrExit) {
		return nil
	}

	return nil
}

type walker struct {
	baseFn    func(RenderContext, *RenderContextBase) error
	objectFn  func(*RenderContextObject) error
	arrayFn   func(*RenderContextArray) error
	stringFn  func(*RenderContextString) error
	urlFn     func(*RenderContextUrl) error
	integerFn func(*RenderContextInteger) error
	floatFn   func(*RenderContextFloat) error
}

func (t *walker) Walk(schema ...RenderContext) error {
	for _, s := range schema {
		if err := t.walk(s); err != nil {
			return realErr(err)
		}
	}

	return nil
}

func (t *walker) walk(schema RenderContext) error {
	if schema == nil {
		return ErrExit
	}

	if t.baseFn != nil && schema.Base() != nil {
		if err := t.baseFn(schema, schema.Base()); err != nil {
			return realErr(err)
		}
	}

	switch c := schema.(type) {
	case *RenderContextObject:
		if t.objectFn != nil {
			if err := t.objectFn(c); err != nil {
				return realErr(err)
			}
		}

		for _, prop := range c.Properties {
			if prop != nil {
				if err := t.Walk(prop); err != nil {
					return realErr(err)
				}
			}
		}
	case *RenderContextArray:
		if t.arrayFn != nil {
			if err := t.arrayFn(c); err != nil {
				return realErr(err)
			}
		}

		if err := t.Walk(c.Items); err != nil {
			return realErr(err)
		}
	case *RenderContextString:
		if t.stringFn != nil {
			if err := t.stringFn(c); err != nil {
				return realErr(err)
			}
		}
	case *RenderContextUrl:
		if t.urlFn != nil {
			if err := t.urlFn(c); err != nil {
				return realErr(err)
			}
		}
	case *RenderContextFloat:
		if t.floatFn != nil {
			if err := t.floatFn(c); err != nil {
				return realErr(err)
			}
		}
	case *RenderContextInteger:
		if t.integerFn != nil {
			if err := t.integerFn(c); err != nil {
				return realErr(err)
			}
		}
	default:
	}

	return nil
}

func NewWalker() *walker {
	return &walker{}
}

func (t *walker) Base(v func(RenderContext, *RenderContextBase) error) *walker {
	t.baseFn = v
	return t
}

func (t *walker) Object(v func(*RenderContextObject) error) *walker {
	t.objectFn = v
	return t
}

func (t *walker) Array(v func(*RenderContextArray) error) *walker {
	t.arrayFn = v
	return t
}

func (t *walker) String(v func(*RenderContextString) error) *walker {
	t.stringFn = v
	return t
}

func (t *walker) Url(v func(*RenderContextUrl) error) *walker {
	t.urlFn = v
	return t
}

func (t *walker) Integer(v func(*RenderContextInteger) error) *walker {
	t.integerFn = v
	return t
}

func (t *walker) Float(v func(*RenderContextFloat) error) *walker {
	t.floatFn = v
	return t
}
