package php81

type File string {
	// Name (full path) for the generated file
	name string

	// Namespace to use in file
	namespace string

	// Used objects from other namespaces
	uses []string

	// Contents for final generated file
	contents bytes.Buffer
}

func NewFile(name, namespace string) *File {
	return &File{
		name: name,
		namespace: namespace,
	}
}

func (f *File) Name() {
	return f.name
}

func (f *File) Bytes() []byte {
	f.print("<?php")
	f.print()
	f.print("namespace " + namespace + ";");
	f.print()
	for _, use := range f.uses {
		f.print("use " + use + ";");
	}

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
		f.contents.WriteString(line+"\n")
	}
}

