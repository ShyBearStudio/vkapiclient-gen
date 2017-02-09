package main

import (
	"encoding/json"
	"fmt"
)

type feErrorPrimitive struct {
	Name        *string `json:"name"`
	Code        *string `json:"code"`
	Description *string `json:"description"`
}

type feMethodParamTypePrimitive struct {
	Type    *string `json:"type"`
	Minimum *int    `json:"minimum"`
	Ref     *string `json:"$ref"`
}

type feParamPrimitive struct {
	Name        *string                     `json:"name"`
	Description *string                     `json:"description"`
	Type        *string                     `json:"type"`
	Items       *feMethodParamTypePrimitive `json:"items"`
	Minimum     *int                        `json:"minimum"`
	Required    *bool                       `json:"required"`
	Enum        []interface{}               `json:"enum"`
	EnumNames   []string                    `json:"enumNames"`
}

type feMethodPrimitive struct {
	Name        *string            `json:"name"`
	Description *string            `json:"description"`
	Open        *bool              `json:"open"`
	Parameters  []feParamPrimitive `json:"parameters"`
}

type MethodsJson struct {
	Errors  []*feErrorPrimitive  `json:"errors"`
	Methods []*feMethodPrimitive `json:"methods"`
}

type feObjectPrimitive struct {
	Properties []*feParamPrimitive `json:"properties"`
	Required   []string            `json:"required"`
}

type ObjectsJson struct {
	Objects map[string]*feObjectPrimitive `json:"definitions"`
}

type ResponseSubPropPrimitive struct {
	Props map[string]*feParamPrimitive `json:"properties"`
}

type ResponsePropPrimitive struct {
	Props map[string]*ResponseSubPropPrimitive `json:"properties"`
}

type ResponsesJson struct {
	Title     string                            `json:"title"`
	Responses map[string]*ResponsePropPrimitive `json:"definitions"`
}

type feSchemas struct {
	methods   *MethodsJson
	objects   *ObjectsJson
	responses *ResponsesJson
}

func (c *vkApiClientCompiler) frontEnd() (*PrimitivesSystem, error) {
	_, err := parseFeSchemas(c.Schemas)
	if err != nil {
		return nil, fmt.Errorf("Cannot parse front end schemas: %v", err)
	}
	return new(PrimitivesSystem), nil
}

func parseFeSchemas(cSchemas map[string][]byte) (*feSchemas, error) {
	var feSchemas = new(feSchemas)
	if err := json.Unmarshal(cSchemas["methods"], &feSchemas.methods); err != nil {
		return nil, fmt.Errorf("Cannot unmarshal methods: %v", err)
	}
	if err := json.Unmarshal(cSchemas["objects"], &feSchemas.objects); err != nil {
		return nil, fmt.Errorf("Cannot unmarshal methods: %v", err)
	}
	if err := json.Unmarshal(cSchemas["responses"], &feSchemas.responses); err != nil {
		return nil, fmt.Errorf("Cannot unmarshal responses: %v", err)
	}
	return feSchemas, nil
}
