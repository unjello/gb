package core

import (
	"bytes"
	"html/template"
)

func ExecuteTemplate(templateName string, templateData string, info interface{}) (string, error) {
	t := template.Must(template.New(templateName).Parse(templateData))
	buffer := new(bytes.Buffer)
	err := t.Execute(buffer, info)
	if err != nil {
		return "", err
	}
	ninjaFile := buffer.String()
	return ninjaFile, nil
}
