package system

func (t *System) Process(uris ...string) error {
	for _, uri := range uris {
		if err := t.processUri(uri); err != nil {
			return err
		}
	}

	// for _, info := range t.sourceInfo {
	// 	bs, err := ioutil.ReadFile(info.LocalPath)

	// 	if err != nil {
	// 		return err
	// 	}

	// 	if util.IsCodeGenFile(info.LocalPath) {
	// 		if err := t.unmarshalCodeGen(info, bs); err != nil {
	// 			return err
	// 		}
	// 	}

	// 	if util.IsJsonSchemaFile(info.LocalPath) {
	// 		if err := t.unmarshalJsonSchema(info, bs); err != nil {
	// 			return err
	// 		}
	// 	}
	// }

	// for _, schema := range t.codeGen {
	// 	if err := t.process1(schema); err != nil {
	// 		return err
	// 	}
	// }

	// for _, schema := range t.codeGen {
	// 	if err := t.process3(schema); err != nil {
	// 		return err
	// 	}
	// }

	// if err := t.process4(); err != nil {
	// 	return err
	// }

	return nil
}

// func (t *System) process1(schema *model.Schema) error {
// 	// for _, v := range schema.Models {
// 	// 	if v.Ref.IsDefined() {
// 	// 		if err := t.jsonSchema.Assert(v.Ref.Get()); err != nil {
// 	// 			return err
// 	// 		}
// 	// 	} else {
// 	// 		if err := t.jsonSchema.Add(v); err != nil {
// 	// 			return err
// 	// 		}
// 	// 	}
// 	// }

// 	// for _, v := range schema.Operations {
// 	// 	if err := t.jsonSchema.Resolve(v.Input); err != nil {
// 	// 		return err
// 	// 	}

// 	// 	if err := t.jsonSchema.Resolve(v.Output); err != nil {
// 	// 		return err
// 	// 	}
// 	// }

// 	return nil
// }

// func (t *System) process3(schema *model.Schema) error {
// 	fn := func(path string, d fs.DirEntry, err error) error {
// 		if err != nil {
// 			return err
// 		}

// 		if d.IsDir() {
// 			return nil
// 		}

// 		localUri := util.Path2Uri(path)

// 		if _, ok := t.template[localUri]; ok {
// 			return model.ErrTemplateDuplicateV(localUri)
// 		}

// 		return nil
// 	}

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
