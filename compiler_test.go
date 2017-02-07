package main

import (
	"path"
	"testing"
)

func TestCompile(t *testing.T) {
	runEndToEndTest(t, "no_object_refs", "methods.go", "objects.go")
	runEndToEndTest(t, "two_responses", "methods.go", "objects.go")
}

func runEndToEndTest(t *testing.T, testName string, goFilesToMatch ...string) {
	compiler := newTestCompiler(t, testName)
	err := compiler.compile()
	if err != nil {
		t.Fatalf("Cannot compile: %v", err)
	}
	testResultDir := getTestResultDirName(testName)
	for _, fileName := range goFilesToMatch {
		compiled := path.Join(compiler.OutputDirName, fileName)
		expected := path.Join(testResultDir, fileName)
		assertGoFilesEquivalent(t, expected, compiled)
	}
}

func getTestResultDirName(testName string) string {
	return path.Join(testDataDir, testName, "results")
}

func assertGoFilesEquivalent(t *testing.T, expected, compiled string) {
	expectedContent := readFile(t, expected)
	compiledContent := readFile(t, compiled)
	if expectedContent != compiledContent {
		t.Fatalf("Expected Go file '%s' is not equivalent to compiled file '%s'.", expected, compiled)
	}
}

func TestValidateValidData(t *testing.T) {
	schemas := map[string]string{
		"objects":   "objects JSON schema",
		"methods":   "methods JSON schema",
		"responses": "responses JSON schema",
	}
	outputDirName := "c:/temp/compilerOutput"
	compiler := newVkApiClienCompiler(schemas, outputDirName)
	err := compiler.validate()
	if err != nil {
		t.Errorf("Schemas should be valid")
	}
}

func TestValidateInvalidSchemas(t *testing.T) {
	schemas := map[string]string{
		"objects":   "objects JSON schema",
		"responses": "responses JSON schema",
	}
	outputDirName := "c:/temp/compilerOutput"
	compiler := newVkApiClienCompiler(schemas, outputDirName)
	err := compiler.validate()
	if err == nil {
		t.Errorf("Schemas should not be valid")
	}
}

func TestValidateInvalidOutputPath(t *testing.T) {
	schemas := map[string]string{
		"objects":   "objects JSON schema",
		"methods":   "methods JSON schema",
		"responses": "responses JSON schema",
	}
	outputDirName := ""
	compiler := newVkApiClienCompiler(schemas, outputDirName)
	err := compiler.validate()
	if err == nil {
		t.Errorf("Schemas should not be valid")
	}
}
