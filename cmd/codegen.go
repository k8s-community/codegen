package main

import (
	"log"
	"strings"
	"text/template"

	"github.com/k8s-community/codegen/pkg/config"
	"github.com/k8s-community/codegen/pkg/generator"
	"github.com/takama/envconfig"
)

func main() {
	config := getConfigForCodeGeneration()

	err := generator.GenerateCode(config)
	if err != nil {
		log.Fatalf("Cannot generate code: %s", err)
	}
}

func getConfigForCodeGeneration() generator.Config {
	var initConfig config.Env
	err := envconfig.Process("codegen", &initConfig)
	if err != nil {
		log.Fatalf("Couldn't get service config: %s", err)
	}

	err = initConfig.Validate()
	if err != nil {
		log.Fatalf("Config validation error: %s", err)
	}

	templateData, err := initConfig.ReadTemplateData()

	config := generator.Config{
		SrcPath:  initConfig.SrcPath,
		DestPath: initConfig.DestPath,

		LeftDelim:  initConfig.LeftDelim,
		RightDelim: initConfig.RightDelim,

		FuncMap: template.FuncMap{
			"ToUpper": strings.ToUpper,
			"ToLower": strings.ToLower,
		},

		SkipPaths:    initConfig.SkipPaths,
		ReplacePaths: initConfig.ReplacePaths,

		TemplateData: templateData,
	}

	return config
}
