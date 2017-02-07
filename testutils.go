package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
	"testing"
)

const (
	testDataDir string = "test_data"
)

func newTestCompiler(t *testing.T, testName string) *vkApiClientCompiler {
	schemas := readTestSchemas(t, testName)
	compilerOutputDir := getCompilerOutputDirName(testName)
	compiler := newVkApiClienCompiler(schemas, compilerOutputDir)
	return compiler
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

func assertStructsEqual(t *testing.T, expected, actual interface{}) {
	expectedJson := toJson(t, expected)
	actualJson := toJson(t, actual)
	if expectedJson != actualJson {
		t.Fatalf("Structs ain't equal\n\texpected: '%v'\n\tactual: '%v'\n", expectedJson, actualJson)
	}
}

func toJson(t *testing.T, data interface{}) string {
	bytes, err := json.Marshal(data)
	if err != nil {
		t.Fatalf("Cannot get json for  '%#v': %v\n", err)
	}
	return string(bytes)
}
