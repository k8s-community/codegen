package template

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

// RecursiveExecutor process templates in selected directory
type RecursiveExecutor struct {
	SrcPath string
	Data    interface{}

	LeftDelim  string
	RightDelim string
	FuncMap    template.FuncMap

	SkipPaths []string
}

// NewRecursiveExecutor inits RecursiveExecutor instance
func NewRecursiveExecutor(srcPath string, data interface{}, leftDelim string, rightDelim string, funcMap template.FuncMap, skipPaths []string) *RecursiveExecutor {
	return &RecursiveExecutor{
		SrcPath:    srcPath,
		Data:       data,
		LeftDelim:  leftDelim,
		RightDelim: rightDelim,
		FuncMap:    funcMap,
		SkipPaths:  skipPaths,
	}
}

// Process all templates from source path
func (e RecursiveExecutor) Process() error {
	// check if the source dir exist
	src, err := os.Stat(e.SrcPath)
	if err != nil {
		return err
	}

	if src.IsDir() {
		return e.processTemplatesDir(e.SrcPath)
	}

	return e.processTemplateFile(e.SrcPath)
}

// ProcessTemplatesDir process templates by path
func (e RecursiveExecutor) processTemplatesDir(path string) error {
	directory, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("cannot open dir: %s", err)
	}
	objects, err := directory.Readdir(-1)
	if err != nil {
		return fmt.Errorf("cannot read dir: %s", err)
	}

	for _, obj := range objects {
		objectPath := path + "/" + obj.Name()
		if e.skipObject(objectPath) {
			continue
		}

		if obj.IsDir() {
			err = e.processTemplatesDir(objectPath)
			if err != nil {
				return fmt.Errorf("cannot process dir %s: %s", objectPath, err)
			}
		} else {
			err = e.processTemplateFile(objectPath)
			if err != nil {
				return fmt.Errorf("cannot process file %s: %s", objectPath, err)
			}
		}
	}

	return nil
}

// processTemplateFile process template file by path
func (e RecursiveExecutor) processTemplateFile(path string) error {
	fileName := filepath.Base(path)

	tpl, err := template.New(fileName).Funcs(e.FuncMap).Delims(e.LeftDelim, e.RightDelim).ParseFiles(path)
	if err != nil {
		return fmt.Errorf("cannot parse file: %s", err)
	}

	f, err := os.Create(path)
	defer f.Close()
	if err != nil {
		return fmt.Errorf("cannot create file: %s", err)
	}

	err = tpl.Execute(f, e.Data)
	if err != nil {
		return fmt.Errorf("cannot execute template: %s", err)
	}

	return nil
}

// skipObject checks by list that path can be skipped
func (e RecursiveExecutor) skipObject(path string) bool {
	for _, skipPath := range e.SkipPaths {
		if path == e.SrcPath+"/"+skipPath {
			return true
		}
	}

	return false
}
