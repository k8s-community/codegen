package template

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

// Config consists of params for template execution
type Config struct {
	SrcPath string
	Data    interface{}

	LeftDelim  string
	RightDelim string

	FuncMap template.FuncMap

	SkipPaths []string
}

// RecursiveExecutor process templates in selected directory
type RecursiveExecutor struct {
	Config
}

// NewRecursiveExecutor inits RecursiveExecutor instance
func NewRecursiveExecutor(config Config) *RecursiveExecutor {
	fmt.Printf("executor: #%v", config)
	return &RecursiveExecutor{
		Config: config,
	}
}

// Process all templates from source path
func (r RecursiveExecutor) Process() error {
	// check if the source dir exist
	src, err := os.Stat(r.SrcPath)
	if err != nil {
		return err
	}

	if src.IsDir() {
		return r.processTemplatesDir(r.SrcPath)
	}

	return r.processTemplateFile(r.SrcPath)
}

// ProcessTemplatesDir process templates by path
func (r RecursiveExecutor) processTemplatesDir(path string) error {
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
		if r.skipObject(objectPath) {
			continue
		}

		if obj.IsDir() {
			err = r.processTemplatesDir(objectPath)
			if err != nil {
				return fmt.Errorf("cannot process dir %s: %s", objectPath, err)
			}
		} else {
			err = r.processTemplateFile(objectPath)
			if err != nil {
				return fmt.Errorf("cannot process file %s: %s", objectPath, err)
			}
		}
	}

	return nil
}

// processTemplateFile process template file by path
func (r RecursiveExecutor) processTemplateFile(path string) error {
	fileName := filepath.Base(path)

	tpl, err := template.New(fileName).Funcs(r.FuncMap).Delims(r.LeftDelim, r.RightDelim).ParseFiles(path)
	if err != nil {
		return fmt.Errorf("cannot parse file: %s", err)
	}

	f, err := os.Create(path)
	defer f.Close()
	if err != nil {
		return fmt.Errorf("cannot create file: %s", err)
	}

	err = tpl.Execute(f, r.Data)
	if err != nil {
		return fmt.Errorf("cannot execute template: %s", err)
	}

	return nil
}

// skipObject checks by list that path can be skipped
func (r RecursiveExecutor) skipObject(path string) bool {
	for _, skipPath := range r.SkipPaths {
		if path == strings.TrimRight(r.SrcPath+"/"+skipPath, "/") {
			return true
		}
	}

	return false
}
