package main

type ErrorPrimitive struct {
	Name        string
	Code        string
	Desctiption string
}

type MethodPrimitive struct {
	Name        string
	Description string
	Parameters  []*ObjectPrimitive
	Responses   []*ResponsePrimitive
	Errors      []*ErrorPrimitive
	Open        bool
}

type ResponsePrimitive struct {
	Name   string
	Object *ObjectPrimitive
}

type ObjectPrimitive struct {
	Name       string
	Properties []*ObjectPropertyPrimitive
}

type ObjectPropertyPrimitive struct {
	Name        string
	Description string
	Required    bool
	Type        BaseTypePrimitive
}

type Type string

const (
	IntegerType     Type = "int"
	StringType      Type = "string"
	NumberType      Type = "float64"
	IntegerEnumType Type = "undefined"
	StringEnumType  Type = "undefined"
	ArrayType       Type = "undefined"
	ObjectType      Type = "undefined"
)

type BaseTypePrimitive interface {
	Type() Type
}

type IntegerTypePrimitive struct {
	MinValue int
	MaxValue int
}

func (t *IntegerTypePrimitive) Type() Type {
	return IntegerType
}

type StringTypePrimitive struct {
}

func (t *StringTypePrimitive) Type() Type {
	return StringType
}

type NumberTypePrimitive struct {
}

func (t *NumberTypePrimitive) Type() Type {
	return NumberType
}

type IntegerEnumRecord struct {
	Value int
	Name  string
}

type IntegerEnumTypePrimitive struct {
	Records []*IntegerEnumRecord
}

func (t *IntegerEnumTypePrimitive) Type() Type {
	return IntegerEnumType
}

type StringEnumRecord struct {
	Value string
	Name  string
}

type StringEnumTypePrimitive struct {
	Records []*StringEnumRecord
}

type ArrayTypePrimitive struct {
	ItemsType BaseTypePrimitive
	MinItems  int
	MaxItems  int
}

func (t *ArrayTypePrimitive) Type() Type {
	return ArrayType
}

type ObjectTypePrimitive struct {
	Object *ObjectPrimitive
}

func (t *ObjectTypePrimitive) Type() Type {
	return ObjectType
}
