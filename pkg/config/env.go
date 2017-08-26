package config

import (
	"fmt"
	"regexp"
)

const (
	maxAppNameLength = 32
)

// Env stores configs for app like host, port, env, ...
type Env struct {
	AppName  string `envconfig:"app_name"`
	AppDesc  string `envconfig:"app_desc"`
	Language string `envconfig:"lang"`

	OwnerName  string `envconfig:"owner_name"`
	OwnerEmail string `envconfig:"owner_email"`

	SrcPath  string `envconfig:"src_path"`
	DestPath string `envconfig:"dest_path"`

	RepositoryURL string `envconfig:"repos_url"`
	ProjectURL    string `envconfig:"project_url"`

	ImageRegistryURL         string `envconfig:"registry_url"`
	ServicesURL              string `envconfig:"services_url"`
	AppURL                   string `envconfig:"app_url"`
	PullSecretName           string `envconfig:"pull_secret_name"`
	TLSSecretNameForRegistry string `envconfig:"tls_secret_name_for_registry"`
	TLSSecretNameForApp      string `envconfig:"tls_secret_name_for_app"`

	DevImageRegistryURL         string `envconfig:"dev_registry_url"`
	DevServicesURL              string `envconfig:"dev_services_url"`
	DevAppURL                   string `envconfig:"dev_app_url"`
	DevPullSecretName           string `envconfig:"dev_pull_secret_name"`
	DevTLSSecretNameForRegistry string `envconfig:"dev_tls_secret_name_for_registry"`
	DevTLSSecretNameForApp      string `envconfig:"dev_tls_secret_name_for_app"`

	EnvPrefix        string
	ServiceEnvParams map[string]string `envconfig:"app_env_params"`

	ExternalServicesEnvParams map[string]string `envconfig:"ext_env_params"`

	LeftDelim  string `envconfig:"left_delim"`
	RightDelim string `envconfig:"right_delim"`

	SkipPaths     []string          `envconfig:"skip_paths"`
	TemplatePaths map[string]string `envconfig:"template_paths"`
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
