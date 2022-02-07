package php81

import (
	"bytes"
	"path"
	"strings"
)

type Field struct {
	Type     string
	Name     string
	Repeated bool
}

type File struct {
	// Name (full path) for the generated file
	name string

	// Namespace to use in file
	namespace string

	// Used objects from other namespaces
	uses []string

	// Fields
	fields []Field

	// Contents for final generated file
	contents bytes.Buffer
}

func NewFile(name, namespace string) *File {
	return &File{
		name:      name,
		namespace: namespace,
	}
}

func (f *File) AddField(n Field) {
	f.fields = append(f.fields, n)
}

func (f *File) Name() string {
	return f.name
}

func (f *File) typeAlias(t string, repeated bool) string {
	if repeated {
		return "array"
	}
	if t == "string" || t == "float" {
		return t
	}
	if strings.Contains(t, "int64") {
		return "string"
	}
	if strings.Contains(t, "int") {
		return "int"
	}
	if strings.Contains(t, "double") {
		return "float"
	}
	if strings.Contains(t, "bool") {
		return "bool"
	}
	return "mixed"
}

func (f *File) typeLiteral(t string, repeated bool) string {
	if repeated {
		return " // []" + t
	}
	alias := f.typeAlias(t, repeated)
	if alias != t {
		return " // " + t
	}
	return ""
}

func (f *File) Bytes() []byte {
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
	f.print("class " + className)
	f.print("{")
	f.print("\tpublic function __construct(")
	for _, v := range f.fields {
		f.print("\t\tpublic " + f.typeAlias(v.Type, v.Repeated) + " $" + v.Name + "," + f.typeLiteral(v.Type, v.Repeated))
	}
	f.print("\t) {}")
	f.print("}")

	return f.contents.Bytes()
}

func (f *File) use(name string) {
	f.uses = append(f.uses, name)
}

func (f *File) print(lines ...string) {
	if len(lines) == 0 {
		f.contents.WriteString("\n")
		return
	}
	for _, line := range lines {
		f.contents.WriteString(line + "\n")
	}
}
