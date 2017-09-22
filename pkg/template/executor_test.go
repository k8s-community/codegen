package template

import (
	"io/ioutil"
	"os"
	"testing"
)

// TestProcess checks that template executor can process files and dirs with files
func TestProcess(t *testing.T) {
	srcDir := "/tmp/dat1/"

	fileName1 := "1.txt"
	fileName2 := "2.txt"
	childDir := "/child"
	filePath1 := srcDir + childDir + "/" + fileName1
	filePath2 := srcDir + "/" + fileName2

	err := os.MkdirAll(srcDir+childDir, os.ModePerm)
	if err != nil {
		t.Fatalf("Сannot create src dir: %s", err)
	}

	expectedText1 := "Hello, {[(.worldTo)]}"
	err = ioutil.WriteFile(filePath1, []byte(expectedText1), 0644)
	if err != nil {
		t.Fatalf("Сannot create file 1: %s", err)
	}

	expectedText2 := "Goodbye, {[(.worldFrom)]}"
	err = ioutil.WriteFile(filePath2, []byte(expectedText2), 0644)
	if err != nil {
		t.Fatalf("Сannot create file 2: %s", err)
	}

	data := make(map[string]string)
	data["worldTo"] = "Nova"
	data["worldFrom"] = "Earth"

	config := Config{
		SrcPath:    srcDir,
		Data:       data,
		LeftDelim:  "{[(",
		RightDelim: ")]}",
	}
	executor := NewRecursiveExecutor(config)

	err = executor.Process()
	defer func() {
		// clean data after test
		os.RemoveAll(srcDir)
	}()
	if err != nil {
		t.Fatalf("Cannot execute template dir: %s", err)
	}

	bytes, err := ioutil.ReadFile(filePath1)
	if err != nil {
		t.Fatalf("Сannot read file 1: %s", err)
	}
	if "Hello, Nova" != string(bytes) {
		t.Fatalf("Cannot execute file 1. Real result: %s", string(bytes))
	}

	bytes, err = ioutil.ReadFile(filePath2)
	if err != nil {
		t.Fatalf("Сannot read file 2: %s", err)
	}
	if "Goodbye, Earth" != string(bytes) {
		t.Fatalf("Cannot execute file 2. Real result: %s", string(bytes))
	}
}
