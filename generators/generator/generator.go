package generator

import (
	"html/template"
	"os"
)

type Generator interface {
	Generate()
}

func WriteFile(path string, content string, data any) {
	// Write file

	template, err := template.New("service").Parse(content)

	if err != nil {
		panic(err)
	}

	f, err := os.Create(path)

	if err != nil {
		panic(err)
	}

	err = template.Execute(f, data)

	if err != nil {
		panic(err)
	}
}
