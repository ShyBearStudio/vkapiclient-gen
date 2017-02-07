package main

import (
	"math"
)

type PrimitivesSystem struct {
	ErrorPrimitives  []*ErrorPrimitive
	MethodPrimitives []*MethodPrimitive
}

type ErrorPrimitive struct {
	Name        string
	Code        int
	Desctiption string
}

type MethodPrimitive struct {
	Name        string
	Description string
	Parameters  []*ParamPrimitive
	Responses   []*ResponsePrimitive
	Errors      []*ErrorPrimitive
	Open        bool
}

type ParamPrimitive struct {
	Name        string
	Desctiption string
	Type        BaseTypePrimitive
	Required    bool
}

func NewParamPrimitive(
	Name, desctiption string, paramType BaseTypePrimitive, required bool) *ParamPrimitive {
	return &ParamPrimitive{
		Name:        Name,
		Desctiption: desctiption,
		Type:        paramType,
		Required:    required,
	}
}

func (m *MethodPrimitive) AddParam(
	Name, desctiption string, paramType BaseTypePrimitive, required bool) *MethodPrimitive {
	param := &ParamPrimitive{
		Name:        Name,
		Desctiption: desctiption,
		Type:        paramType,
		Required:    required,
	}
	m.Parameters = append(m.Parameters, param)
	return m
}

func (m *MethodPrimitive) AddError(e *ErrorPrimitive) *MethodPrimitive {
	m.Errors = append(m.Errors, e)
	return m
}

func (m *MethodPrimitive) AddResponse(r *ResponsePrimitive) *MethodPrimitive {
	m.Responses = append(m.Responses, r)
	return m
}

type ResponsePrimitive struct {
	Name   string
	Object *ObjectPrimitive
}

func NewResponsePrimitive(n string, o *ObjectPrimitive) *ResponsePrimitive {
	return &ResponsePrimitive{
		Name:   n,
		Object: o,
	}
}

type ObjectPrimitive struct {
	Name       string
	Properties []*ParamPrimitive
}

func NewObjectPrimitive(n string) *ObjectPrimitive {
	return &ObjectPrimitive{Name: n}
}

func (o *ObjectPrimitive) AddProp(p *ParamPrimitive) *ObjectPrimitive {
	o.Properties = append(o.Properties, p)
	return o
}

type Type string

const (
	IntegerType     Type = "integer"
	StringType      Type = "string"
	NumberType      Type = "nubmber"
	BooleanType     Type = "boolean"
	IntegerEnumType Type = "integerEnum"
	StringEnumType  Type = "stringEnum"
	ArrayType       Type = "array"
	ObjectType      Type = "object"
)

type BaseTypePrimitive interface {
	Type() Type
}

type IntegerTypePrimitive struct {
	Name     Type
	MinValue int
	MaxValue int
}

func (t *IntegerTypePrimitive) Type() Type {
	return t.Name
}

func NewIntegerTypePrimitive(minValue, maxValue int) *IntegerTypePrimitive {
	return &IntegerTypePrimitive{
		Name:     IntegerType,
		MinValue: minValue,
		MaxValue: maxValue,
	}
}

func NewIntegerTypePrimitiveMinBound(minValue int) *IntegerTypePrimitive {
	return NewIntegerTypePrimitive(minValue, math.MaxInt32)
}

func NewIntegerTypePrimitiveNoBounds() *IntegerTypePrimitive {
	return NewIntegerTypePrimitive(math.MinInt32, math.MaxInt32)
}

type StringTypePrimitive struct {
	Name Type
}

func (t *StringTypePrimitive) Type() Type {
	return t.Name
}

func NewStringTypePrimitive() *StringTypePrimitive {
	return &StringTypePrimitive{Name: StringType}
}

type NumberTypePrimitive struct {
	Name Type
}

func (t *NumberTypePrimitive) Type() Type {
	return t.Name
}

func NewNumberTypePrimitive() *NumberTypePrimitive {
	return &NumberTypePrimitive{Name: NumberType}
}

type BooleanTypePrimitive struct {
	Name Type
}

func (t *BooleanTypePrimitive) Type() Type {
	return t.Name
}

func NewBooleanTypePrimitive() *BooleanTypePrimitive {
	return &BooleanTypePrimitive{Name: BooleanType}
}

type IntegerEnumRecord struct {
	Value int
	Name  string
}

type IntegerEnumTypePrimitive struct {
	Name    Type
	Records []*IntegerEnumRecord
}

func (t *IntegerEnumTypePrimitive) Type() Type {
	return t.Name
}

func NewIntegerEnumTypePrimitive() *IntegerEnumTypePrimitive {
	return &IntegerEnumTypePrimitive{Name: IntegerEnumType}
}

func (t *IntegerEnumTypePrimitive) AddRecord(
	value int, Name string) *IntegerEnumTypePrimitive {
	rec := &IntegerEnumRecord{Value: value, Name: Name}
	t.Records = append(t.Records, rec)
	return t
}

type StringEnumRecord struct {
	Value string
	Name  string
}

type StringEnumTypePrimitive struct {
	Name    Type
	Records []*StringEnumRecord
}

func (t *StringEnumTypePrimitive) Type() Type {
	return t.Name
}

func NewStringEnumTypePrimitive() *StringEnumTypePrimitive {
	return &StringEnumTypePrimitive{Name: StringEnumType}
}

func (t *StringEnumTypePrimitive) AddRecord(
	value string, Name string) *StringEnumTypePrimitive {
	rec := &StringEnumRecord{Value: value, Name: Name}
	t.Records = append(t.Records, rec)
	return t
}

type ArrayTypePrimitive struct {
	Name      Type
	ItemsType BaseTypePrimitive
	MinItems  int
	MaxItems  int
}

func (t *ArrayTypePrimitive) Type() Type {
	return t.Name
}

func NewArrayTypePrimitive(
	itemsType BaseTypePrimitive, minItems, maxItems int) *ArrayTypePrimitive {
	return &ArrayTypePrimitive{
		Name:      ArrayType,
		ItemsType: itemsType,
		MinItems:  minItems,
		MaxItems:  maxItems,
	}
}

func NewArrayTypePrimitiveNoLimits(
	itemsType BaseTypePrimitive) *ArrayTypePrimitive {
	return NewArrayTypePrimitive(itemsType, math.MinInt32, math.MaxInt32)
}

type ObjectTypePrimitive struct {
	Name   Type
	Object *ObjectPrimitive
}

func (t *ObjectTypePrimitive) Type() Type {
	return t.Name
}

func NewObjectTypePrimitive(object *ObjectPrimitive) *ObjectTypePrimitive {
	return &ObjectTypePrimitive{Name: ObjectType, Object: object}
}
