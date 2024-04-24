package service

import (
	_ "embed"
	"flag"
	"generators/generator"
)

type ServiceGenerator struct {
	PackageName string
	RootPath    string
}

//go:embed templates/service.txt
var service string

func NewServiceGenerator() *ServiceGenerator {
	var packageName string
	var rootPath string

	flag.StringVar(&packageName, "package", "", "package name")
	flag.StringVar(&rootPath, "root", "", "root path")

	return &ServiceGenerator{
		PackageName: packageName,
		RootPath:    rootPath,
	}
}

func (s *ServiceGenerator) Generate() {

	data := struct {
		PackageName string
	}{
		PackageName: s.PackageName,
	}

	generator.WriteFile(s.RootPath+"/service.go", service, data)
	// generator.WriteFile(s.RootPath+"/controller.go", controller, data)
	// generator.WriteFile(s.RootPath+"/entities/model.go", repository, data)
	// generator.WriteFile(s.RootPath+"/dtos/create.go", createDto, data)
	// generator.WriteFile(s.RootPath+"/dtos/update.go", updateDto, data)
}
