package loader

// func (t *Loader) uris2files(uris ...string) ([]string, error) {
// 	var files []string

// 	for _, uri := range uris {
// 		var uriType uritype.UriType

// 		if err := t.detectUriType(uri, &uriType); err != nil {
// 			return files, err
// 		}

// 		switch uriType {
// 		case uritype.File:
// 			p := strings.TrimPrefix(uri, "file://")
// 			stat, err := os.Stat(p)

// 			if err != nil {
// 				return files, err
// 			}

// 			if stat.IsDir() {
// 				infos, err := ioutil.ReadDir(p)

// 				if err != nil {
// 					return files, err
// 				}

// 				for _, info := range infos {
// 					file := filepath.Join(p, info.Name())
// 					file = "file://" + file
// 					files = append(files, file)
// 				}
// 			} else {
// 				files = append(files, uri)
// 			}
// 		case uritype.Http:
// 			files = append(files, uri)
// 		}

// 	}

// 	return files, nil
// }
