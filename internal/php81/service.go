package php81

import (
	"bytes"
	"path"
	"strings"
)

type Service struct {
	// Name (full path) for the generated file
	name string

	// Namespace to use in file
	namespace string

	// Comment from proto file
	comment string

	// Used objects from other namespaces
	uses []string

	// Fields
	funcs []string

	// Contents for final generated file
	contents bytes.Buffer
}

func NewService(name, namespace string) *Service {
	return &Service{
		name:      name,
		namespace: namespace,
	}
}

func (f *Service) Name() string {
	return f.name
}

func (f *Service) addFunction(def []string) {
	f.funcs = append(f.funcs, "\t" + strings.TrimSpace(strings.Join(def, "\n\t")))
}

func (f *Service) Bytes() []byte {
	f.print("<?php")
	f.print()
	f.print("namespace " + f.namespace + ";")
	f.print()
	for _, use := range f.uses {
		f.print("use " + use + ";")
	}
	if len(f.uses) > 0 {
		f.print()
	}
	className := strings.TrimSuffix(path.Base(f.name), ".php")
	f.print("/** " + f.comment + " */")
	f.print("interface " + className)
	f.print("{")
	for k, v := range f.funcs {
		if (k > 0) {
			f.print()
		}
		f.print(v)
	}
	f.print("}")

	return f.contents.Bytes()
}

func (f *Service) use(name string) {
	f.uses = append(f.uses, name)
}

func (f *Service) print(lines ...string) {
	if len(lines) == 0 {
		f.contents.WriteString("\n")
		return
	}
	for _, line := range lines {
		f.contents.WriteString(line + "\n")
	}
}
