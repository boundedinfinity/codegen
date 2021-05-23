package loader

import (
	"boundedinfinity/codegen/model"
	"boundedinfinity/codegen/util"
	"path"
	"strings"
)

func (t *Loader) processNamespace1() error {
	createNs := func(ns string) {
		working := path.Dir(ns)

		for {
			if _, ok := t.outputNamespace[working]; !ok {
				ns := model.NewOutputNamespace(working)
				t.outputNamespace[working] = ns
			}

			working = path.Dir(working)

			if working == "." {
				break
			}
		}
	}

	t.outputNamespace["."] = model.NewOutputNamespace(".")

	for n := range t.outputModels {
		createNs(n)
	}

	for n := range t.outputOperations {
		createNs(n)
	}

	return nil
}

func (t *Loader) processNamespace2() error {
	for n1, ns1 := range t.outputNamespace {
		for n2 := range t.outputNamespace {
			if n1 == n2 {
				continue
			}

			if n1 == "." || strings.HasPrefix(n2, n1) {
				ns1.Children = append(ns1.Children, n2)
			}
		}
	}

	for _, ns := range t.outputNamespace {
		var children []string

		for _, child := range ns.Children {
			var temp string
			temp = child
			temp = strings.TrimPrefix(temp, ns.Name)
			temp = strings.TrimPrefix(temp, "/")
			comps := strings.Split(temp, "/")

			if len(comps) >= 1 {
				children = append(children, comps[0])
			}
		}

		children = util.StrSliceDedup(children)
		ns.Children = children
	}

	return nil
}

func (t *Loader) processNamespace3() error {
	for _, ns := range t.outputNamespace {
		t.OutputSpec.Namespaces = append(t.OutputSpec.Namespaces, ns)
	}

	return nil
}
