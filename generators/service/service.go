package service

import (
	_ "embed"
	"generators/generator"
	"log"
	"os"
)

type ServiceGenerator struct {
	PackageName string
	RootPath    string
	RootPackage string
}

//go:embed templates/controller.txt
var controller string

//go:embed templates/service.txt
var service string

//go:embed templates/entities/model.txt
var repository string

//go:embed templates/dtos/create.txt
var createDto string

//go:embed templates/dtos/update.txt
var updateDto string

func NewServiceGenerator(packageName string, rootPath string) *ServiceGenerator {
	return &ServiceGenerator{
		PackageName: packageName,
		RootPath:    rootPath,
		RootPackage: rootPath,
	}
}

func (s *ServiceGenerator) Generate() {

	a, _ := os.Getwd()

	if s.PackageName == "" {
		panic("Package name is required")
	}

	if s.RootPackage == "" {
		s.RootPackage = s.PackageName
		log.Println(1)

	} else {
		s.RootPackage = s.RootPackage + "/" + s.PackageName
		log.Println(2)

	}

	if s.RootPath == "" {
		log.Println(3)
		s.RootPath = a + "/" + s.PackageName
	} else {
		log.Println(4)
		s.RootPath = a + "/" + s.RootPath + "/" + s.PackageName
	}

	log.Println(s.PackageName, s.RootPackage)

	os.Mkdir(s.RootPath, os.ModePerm)

	generator.WriteFile(s.RootPath+"/service.go", service, s)
	generator.WriteFile(s.RootPath+"/controller.go", controller, s)

	os.Mkdir(s.RootPath+"/entities", os.ModePerm)
	os.Mkdir(s.RootPath+"/dtos", os.ModePerm)
	generator.WriteFile(s.RootPath+"/entities/model.go", repository, s)
	generator.WriteFile(s.RootPath+"/dtos/create.go", createDto, s)
	generator.WriteFile(s.RootPath+"/dtos/update.go", updateDto, s)

}
