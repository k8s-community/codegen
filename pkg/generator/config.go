package generator

import (
	// "text/template"
	"fmt"
	"regexp"

	codeTemplate "github.com/k8s-community/codegen/pkg/template"
)

// Config stores app configs
type Config struct {
	AppName string

	codeTemplate.Config

	SrcPath  string
	DestPath string

	ReplacePaths map[string]string
}

const (
	maxAppNameLength = 32
)

// Validate fields from environment
func (e Config) Validate() error {
	err := e.validateServiceName()
	if err != nil {
		return err
	}

	return nil
}

// validateServiceName validate app name regarding k8s requirements
func (e Config) validateServiceName() error {
	// TODO: move regexp in config
	r, err := regexp.Compile("^([a-z]{1}[a-z-0-9]+)$")
	if err != nil {
		return err
	}

	if !r.MatchString(e.AppName) {
		return fmt.Errorf("serviceName `%s` must be started with letter and consisted of lowercase letters (a-z), numbers and '-' symbol", e.AppName)
	}

	if len(e.AppName) >= maxAppNameLength {
		return fmt.Errorf("length of serviceName `%s` must be <64 symbols", e.AppName)
	}

	return nil
}
