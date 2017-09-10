package generator

import "text/template"

// Config stores app configs
type Config struct {
	SrcPath  string
	DestPath string

	LeftDelim  string
	RightDelim string

	FuncMap template.FuncMap

	SkipPaths    []string
	ReplacePaths map[string]string

	TemplateData map[string]interface{}
}
