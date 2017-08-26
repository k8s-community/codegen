package template

import (
	"bytes"
	"fmt"
	"text/template"
)

// ExecuteFromMap executes templates from map and creates new map with results
func ExecuteFromMap(list map[string]string, leftDelim string, rightDelim string, data interface{}) (map[string]string, error) {
	newList := make(map[string]string)

	for srcPath, templatePath := range list {
		var tpl bytes.Buffer
		t, err := template.New("").Delims(leftDelim, rightDelim).Parse(templatePath)

		err = t.Execute(&tpl, data)
		if err != nil {
			return nil, fmt.Errorf("cannot change path for %s", templatePath)
		}

		newList[srcPath] = tpl.String()
	}

	return newList, nil
}
