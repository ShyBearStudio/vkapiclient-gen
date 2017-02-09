package main

import (
	"fmt"
)

type compiler interface {
	compile()
}

type vkApiClientCompiler struct {
	Schemas       map[string][]byte
	OutputDirName string
}

func newVkApiClienCompiler(
	schemas map[string][]byte, outputDirName string) *vkApiClientCompiler {
	return &vkApiClientCompiler{
		Schemas:       schemas,
		OutputDirName: outputDirName,
	}
}

func (c *vkApiClientCompiler) compile() error {
	if err := c.validate(); err != nil {
		return err
	}
	primitives, err := c.frontEnd()
	if err != nil {
		return fmt.Errorf("Cannot fulfill front-end phase: '%v'", err)
	}
	err = c.backEnd(primitives)
	if err != nil {
		return fmt.Errorf("Cannot fulfill back-end phase: '%v'", err)
	}
	return nil
}

// Validates vk API client compiler, valid if satisfies:
//  - there should be reference with name 'methods';
//  - output directory name should not be empty;
func (c *vkApiClientCompiler) validate() error {
	_, ok := c.Schemas["methods"]
	if ok != true {
		return fmt.Errorf("There is no reference for 'methods' schema.")
	}
	if len(c.OutputDirName) == 0 {
		return fmt.Errorf("Output directory name is not valid: '%s'", c.OutputDirName)
	}
	return nil
}
