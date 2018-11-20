package core

import (
	"bytes"
	"html/template"

	"github.com/unjello/gb/layout"
)

func getTestName(file layout.SourceFile) string {
	return file.BaseName
}

func ExecuteTemplate(templateName string, templateData string, info interface{}) (string, error) {
	t := template.Must(template.New(templateName).Funcs(template.FuncMap{
		"getTestName": getTestName,
	}).Parse(templateData))
	buffer := new(bytes.Buffer)
	err := t.Execute(buffer, info)
	if err != nil {
		return "", err
	}
	ninjaFile := buffer.String()
	return ninjaFile, nil
}
