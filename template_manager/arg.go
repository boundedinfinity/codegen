package template_manager

type Arg func(*TemplateManager)

// func OutputDir(v string) Arg {
// 	return func(t *TemplateManager) {
// 		t.outputDir = optioner.Some(v)
// 	}
// }

func (t *TemplateManager) init() error {
	return nil
}
