package template

import (
	"reflect"
	"testing"
)

func TestExecuteFromMap(t *testing.T) {
	originalMap := make(map[string]string)
	originalMap["foo/bar.txt"] = "foo/{[(.name)]}.txt" // must be changed
	originalMap["foo"] = "bar"                         // no changes
	originalMap["bar"] = "{{name}}"                    // no changes because delimeters are other

	data := make(map[string]string)
	data["name"] = "tttt"

	expectedResult := make(map[string]string)
	expectedResult["foo/bar.txt"] = "foo/tttt.txt"
	expectedResult["foo"] = "bar"
	expectedResult["bar"] = "{{name}}"

	realResult, err := ExecuteFromMap(originalMap, "{[(", ")]}", data)
	if err != nil {
		t.Fatalf("Cannot execute template on map of string: %s", err)
	}

	eq := reflect.DeepEqual(expectedResult, realResult)
	if !eq {
		t.Fatalf("Result is incorrect. Expected: %v, real: %v", expectedResult, realResult)
	}
}
