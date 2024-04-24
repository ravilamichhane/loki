package generator

import (
	"html/template"
	"os"
	"strings"
)

type Generator interface {
	Generate()
}

func WriteFile(path string, content string, data any) {

	funcMap := template.FuncMap{
		"ToUpper": strings.ToUpper,
		"ToCapitalize": func(s string) string {
			return strings.ToUpper(string(s[0])) + s[1:]
		},
	}
	// Write file

	template, err := template.New("service").Funcs(
		funcMap).Parse(content)

	if err != nil {
		panic(err)
	}

	f, err := os.Create(path)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)

	if err != nil {
		panic(err)
	}

	defer file.Close()
	err = template.Execute(file, data)

	if err != nil {
		panic(err)
	}
}
