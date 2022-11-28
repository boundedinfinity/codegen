package system

// 	for _, v := range schema.Templates.Files {
// 		if v.Name == "" {
// 			return model.ErrTemplateEmpty
// 		}

// 		path := util.Uri2Path(v.Name)
// 		path, err := filepath.Abs(path)

// 		if err != nil {
// 			return err
// 		}

// 		file, err := pather.IsFileErr(path)

// 		if err != nil {
// 			return err
// 		}

// 		if file {
// 			if _, ok := t.template[v.Name]; ok {
// 				return model.ErrTemplateDuplicateV(v.Name)
// 			}
// 		} else {
// 			if err := filepath.WalkDir(path, fn); err != nil {
// 				return err
// 			}
// 		}
// 	}

// 	return nil
// }

// func (t *System) process4() error {
// 	for _, v := range t.template {
// 		var typ template_type.TemplateType

// 		if err := t.detectTemplateType(v.Name, &typ); err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }
