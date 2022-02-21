package model

import (
	"bytes"
	"fmt"
	"path"
	"strings"
)

type Message struct {
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

func NewMessage(name, namespace string) *Message {
	return &Message{
		name:      name,
		namespace: namespace,
	}
}

func (f *Message) AddField(n Field) {
	f.fields = append(f.fields, n)
}

func (f *Message) Name() string {
	return f.name
}

func (f *Message) nativeDefault(kind string) string {
	switch kind {
	case "int":
		return "0"
	case "array":
		return "[]"
	case "double":
		return "0.0"
	case "float":
		return "0.0"
	case "string":
		return "\"\""
	case "bool":
		return "false"
	}
	return "null"
}

func (f *Message) Bytes() []byte {
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
		fieldType := v.Type
		fieldFormat := v.Type
		if v.Repeated {
			fieldType = "array"
		} else {
			if p, ok := typeAliases[v.Type]; ok {
				fieldType = p.Type
				fieldFormat = p.Format
			}
		}
		if fieldType == fieldFormat {
			fieldFormat = ""
		}
		format := "\t\tpublic ?%s $%s = %s,"
		if fieldFormat != "" {
			// add comment with field format
			format += " // %s"
			f.print(fmt.Sprintf(format, fieldType, v.Name, f.nativeDefault(fieldType), fieldFormat))
			continue
		}
		// field without separate format comment
		f.print(fmt.Sprintf(format, fieldType, v.Name, f.nativeDefault(fieldType)))
	}
	f.print("\t) {}")
	f.print("}")

	return f.contents.Bytes()
}

func (f *Message) print(lines ...string) {
	if len(lines) == 0 {
		f.contents.WriteString("\n")
		return
	}
	for _, line := range lines {
		f.contents.WriteString(line + "\n")
	}
}
