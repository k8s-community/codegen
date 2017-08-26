package utils

import (
	"fmt"
	"io"
	"os"
)

// CopyFile copies file from srcFilePath to destFilePath with current filemode
func CopyFile(srcFilePath, destFilePath string) (err error) {
	srcFile, err := os.Open(srcFilePath)
	if err != nil {
		return fmt.Errorf("cannot open src file: %s", err)
	}
	defer srcFile.Close()

	destfile, err := os.Create(destFilePath)
	if err != nil {
		return fmt.Errorf("cannot create dest file: %s", err)
	}
	defer destfile.Close()

	_, err = io.Copy(destfile, srcFile)
	if err != nil {
		return fmt.Errorf("cannot copy file: %s", err)
	}

	sourceinfo, err := os.Stat(srcFilePath)
	if err != nil {
		return fmt.Errorf("cannot get info of src file: %s", err)
	}

	err = os.Chmod(destFilePath, sourceinfo.Mode())
	if err != nil {
		return fmt.Errorf("cannot copy file mode: %s", err)
	}

	return nil
}

// CopyDir copies dir from srcDirPath to destDirPath with current filemodes
func CopyDir(srcDirPath, destDirPath string) error {
	dirInfo, err := os.Stat(srcDirPath)
	if err != nil {
		return fmt.Errorf("cannot get info of src dir: %s", err)
	}

	directory, err := os.Open(srcDirPath)
	if err != nil {
		return fmt.Errorf("cannot open src dir: %s", err)
	}
	defer directory.Close()

	objects, err := directory.Readdir(-1)
	if err != nil {
		return fmt.Errorf("cannot read src dir: %s", err)
	}

	err = os.MkdirAll(destDirPath, dirInfo.Mode())
	if err != nil {
		return fmt.Errorf("cannot create dest dir: %s", err)
	}

	for _, obj := range objects {
		srcObjectPath := srcDirPath + "/" + obj.Name()
		destObjectPath := destDirPath + "/" + obj.Name()

		if obj.IsDir() {
			err = CopyDir(srcObjectPath, destObjectPath)
			if err != nil {
				return err
			}
		} else {
			err = CopyFile(srcObjectPath, destObjectPath)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
