package generator

type StructDescriptor struct {
	Comment    string
	Package    string
	Name       string
	Properties []StructPropertyDescriptor
	Functions  []StructFunctionDescriptor
	Imports    []string
}

type StructFunctionDescriptor struct {
	Comment string
	Name    string
	Public  bool
	Inputs  []string
	Outputs []string
	Body    string
}

type StructPropertyDescriptor struct {
	Comment string
	Package string
	Name    string
	Tags    []StructPropertyTagDescriptor
}

type StructPropertyTagDescriptor struct {
	Name   string
	Values []string
}

type ValidatorDescriptor struct {
	Comment string
}

type StringDescriptorValidator struct {
}
