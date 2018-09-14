package model

type Struct struct {
	StructName string
	Attributes Attributes
}

type Attributes struct {
	Attribute []Attribute
}

type Attribute struct {
	Name     string
	DataType string
	Tag      string
}

// FOR COMMANDS
type Command struct {
	Name   string
	Status bool
}

// FOR ATTRIBUTE count
type AttributeCount struct {
	Name  string
	Count int
}
