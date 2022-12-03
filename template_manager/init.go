package template_manager

import "fmt"

func (t *TemplateManager) init() error {
	if t.cacher == nil {
		return fmt.Errorf("cacher is nil")
	}

	if t.pathMap == nil {
		return fmt.Errorf("pathMap is nil")
	}

	if t.funcs == nil {
		return fmt.Errorf("funcs is nil")
	}

	return nil
}
