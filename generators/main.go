package main

import (
	"flag"
	"fmt"
	"generators/generator"
	"generators/service"
)

type defaultGenerator struct {
}

func (d defaultGenerator) Generate() {
	fmt.Println("No generator found")
}

func main() {

	var generatorToRun generator.Generator

	genType := flag.String("t", "Kevin", "The name of the user")
	flag.Parse()

	switch *genType {
	case "service":
		generatorToRun = service.NewServiceGenerator()
	default:
		generatorToRun = defaultGenerator{}
	}

	generatorToRun.Generate()

}
