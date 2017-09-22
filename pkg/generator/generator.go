package generator

import (
	"fmt"
	"os"

	"github.com/k8s-community/codegen/pkg/template"
	"github.com/k8s-community/codegen/pkg/utils"
)

// GenerateCode generates code for some language from templates using config
func GenerateCode(config Config) error {
	err := copyTemplatesDir(config)
	if err != nil {
		return err
	}

	err = executeTemplatesDir(config)
	if err != nil {
		return err
	}

	return nil
}

func copyTemplatesDir(config Config) error {
	_, err := os.Stat(config.DestPath)
	if err == nil {
		return fmt.Errorf("dest path %s already exists", config.DestPath)
	}

	err = utils.CopyDir(config.SrcPath, config.DestPath)
	if err != nil {
		return fmt.Errorf("cannot copy dir: %s", err)
	}

	// paste app name and etc. in template paths
	pathsMapForMove, err := template.ExecuteFromMap(config.ReplacePaths, config.LeftDelim, config.RightDelim, config.Data)
	if err != nil {
		return fmt.Errorf("cannot process templates for rename/move paths: %s", err)
	}
	for srcPath, destPath := range pathsMapForMove {
		err = os.Rename(config.DestPath+"/"+srcPath, config.DestPath+"/"+destPath)
		if err != nil {
			return fmt.Errorf("cannot rename/move path %s to %s", config.DestPath+"/"+srcPath, config.DestPath+"/"+destPath)
		}
	}

	return nil
}

func executeTemplatesDir(config Config) error {
	r := template.NewRecursiveExecutor(config.Config)
	err := r.Process()
	if err != nil {
		return fmt.Errorf("cannot process templates: %s", err)
	}

	return nil
}
