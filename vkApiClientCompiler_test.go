package main

import (
	"io/ioutil"
	"os"
	"path"
	"testing"
)

const (
	testDataDir string = "test_data"
)

func TestCompile(t *testing.T) {
	runEndToEndTest(t, "no_object_refs", "methods.go", "objects.go")
	runEndToEndTest(t, "two_responses", "methods.go", "objects.go")
}

func runEndToEndTest(t *testing.T, testName string, goFilesToMatch ...string) {
	schemas := readTestSchemas(t, testName)
	compilerOutputDir := getCompilerOutputDirName(testName)
	compiler := newVkApiClienCompiler(schemas, compilerOutputDir)
	err := compiler.compile()
	if err != nil {
		t.Fatalf("Cannot compile: %v", err)
	}
	testResultDir := getTestResultDirName(testName)
	for _, fileName := range goFilesToMatch {
		compiled := path.Join(compilerOutputDir, fileName)
		expected := path.Join(testResultDir, fileName)
		assertGoFilesEquivalent(t, expected, compiled)
	}
}

func readTestSchemas(t *testing.T, testName string) map[string]string {
	schemas := make(map[string]string)
	schemas["methods"] = getTestInputSchema(t, testName, "methods.json")
	schemas["objects"] = getTestInputSchema(t, testName, "objects.json")
	schemas["responses"] = getTestInputSchema(t, testName, "responses.json")
	return schemas
}

func getTestInputSchema(t *testing.T, testName, fileName string) string {
	path := path.Join(testDataDir, testName, "inputs", fileName)
	return readFile(t, path)
}

func getCompilerOutputDirName(testName string) string {
	return path.Join(testDataDir, testName, "run")
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

func readFile(t *testing.T, fileName string) string {
	if !fileExist(fileName) {
		t.Fatalf("file '%s' not exist", fileName)
	}
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		t.Fatalf("Cannot get content of file '%s': %v", fileName, err)
	}
	return string(content)
}

func fileExist(fileName string) bool {
	if _, err := os.Stat(fileName); err == nil {
		return true
	}
	return false
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
