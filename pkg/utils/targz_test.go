package utils

import (
	"io/ioutil"
	"os"
	"os/exec"
	"testing"
)

// TestCreateTarGzFile checks that a file will be archived in destFile from srcFile
func TestCreateTarGzFile(t *testing.T) {
	fileName := "dat1.txt"
	srcFile := "/tmp/" + fileName
	destDir := "/tmp/"
	destFile := destDir + "dat2.tar.gz"

	expectedText := "test\ncopy\nfile\n"
	err := ioutil.WriteFile(srcFile, []byte(expectedText), 0644)
	if err != nil {
		t.Fatalf("Сannot create src file: %s", err)
	}

	err = CreateTarGzArchive(srcFile, destFile)
	defer func() {
		// clean data after test
		os.Remove(srcFile)
		os.Remove(destFile)
		// unpacked file
		os.Remove(destDir + fileName)
	}()
	if err != nil {
		t.Fatalf("Cannot archive file: %s", err)
	}

	cmd := exec.Command("tar", "-zvxf", destFile, "-C", destDir)
	err = cmd.Run()
	if err != nil {
		t.Fatalf("Cannot untar archive: %s", err)
	}

	bytes, err := ioutil.ReadFile(destDir + fileName)
	if err != nil {
		t.Fatalf("Сannot read dest file: %s", err)
	}

	if expectedText != string(bytes) {
		t.Fatal("Expected file content does not equal real")
	}
}

// TestCreateTarGzDir checks that a directory will be archived in destDir from srcDir
func TestCreateTarGzDir(t *testing.T) {
	srcDir := "/tmp/dat3/"
	destDir := "/tmp/"
	destFile := destDir + "dat4.tar.gz"

	fileName := "data.txt"
	childDir := "/child"
	filePath := childDir + "/" + fileName

	err := os.MkdirAll(srcDir+childDir, os.ModePerm)
	if err != nil {
		t.Fatalf("Сannot create src dir: %s", err)
	}

	expectedText := "test\ncopy\ndir\n"
	err = ioutil.WriteFile(srcDir+filePath, []byte(expectedText), 0644)
	if err != nil {
		t.Fatalf("Сannot create src file: %s", err)
	}

	err = CreateTarGzArchive(srcDir, destFile)
	defer func() {
		// clean data after test
		os.RemoveAll(srcDir)
		os.RemoveAll(destFile)
		// unpacked file
		os.RemoveAll(destDir + childDir)
	}()
	if err != nil {
		t.Fatalf("Cannot archive dir: %s", err)
	}

	cmd := exec.Command("tar", "-zvxf", destFile, "-C", destDir)
	err = cmd.Run()
	if err != nil {
		t.Fatalf("Cannot untar archive: %s", err)
	}

	bytes, err := ioutil.ReadFile(destDir + filePath)
	if err != nil {
		t.Fatalf("Сannot read dest file: %s", err)
	}

	if expectedText != string(bytes) {
		t.Fatal("Expected file content does not equal real")
	}
}
