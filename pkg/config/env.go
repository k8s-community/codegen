package config

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
)

const (
	maxAppNameLength = 32
)

// Env stores configs for app like host, port, env, ...
type Env struct {
	AppName string `envconfig:"app_name"`

	SrcPath  string `envconfig:"src_path"`
	DestPath string `envconfig:"dest_path"`

	LeftDelim  string `envconfig:"left_delim"`
	RightDelim string `envconfig:"right_delim"`

	SkipPaths    []string          `envconfig:"skip_paths"`
	ReplacePaths map[string]string `envconfig:"replace_paths"`

	TemplateDataPath string `envconfig:"template_data_path"`
}

// ReadTemplateData gets all data for template
func (e Env) ReadTemplateData() (map[string]interface{}, error) {
	templateData := make(map[string]interface{})

	file, err := os.Open(e.TemplateDataPath)
	if err != nil {
		return nil, fmt.Errorf("cannot open config file: %s", err.Error())
	}

	jsonParser := json.NewDecoder(file)
	if err = jsonParser.Decode(&templateData); err != nil {
		return nil, fmt.Errorf("cannot parse config file: %s", err.Error())
	}

	templateData["appName"] = e.AppName

	return templateData, nil
}

// Validate fields from environment
func (e Env) Validate() error {
	err := e.validateAppName()
	if err != nil {
		return err
	}

	return nil
}

// validateAppName validate app name regarding k8s requirements
func (e Env) validateAppName() error {
	// TODO: move regexp in config
	r, err := regexp.Compile("^([a-z]{1}[a-z-0-9]+)$")
	if err != nil {
		return err
	}

	if !r.MatchString(e.AppName) {
		return fmt.Errorf("appName must be started with letter and consisted of lowercase letters (a-z), numbers and '-' symbol")
	}

	if len(e.AppName) >= maxAppNameLength {
		return fmt.Errorf("length of AppName must be <64 symbols")
	}

	return nil
}
