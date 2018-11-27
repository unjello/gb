package core

import (
	"bytes"
	"strings"
	"text/template"

	"github.com/unjello/gb/layout"
)

func getTestName(file layout.SourceFile) string {
	slices := strings.Split(strings.TrimRight(file.RelPath, file.Extension), "/")
	filtered := slices[:0]
	for _, v := range slices {
		if v != ".." {
			filtered = append(filtered, v)
		}
	}
	return strings.Join(filtered, "_")
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
