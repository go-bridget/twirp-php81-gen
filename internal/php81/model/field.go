package model

type Field struct {
	Type     string
	Name     string
	Repeated bool
}

func NewField(kind, name string, repeated bool) Field {
	return Field{
		Type:     kind,
		Name:     name,
		Repeated: repeated,
	}
}
