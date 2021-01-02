package main

import (
	"fmt"

	"github.com/go-yaml/yaml"
)

func main() {

	p := Person{"Onur", "Teber", 23}

	y, err := yaml.Marshal(p)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(y))

}

type Person struct {
	FirstName string `yaml:"first_name"`
	LastName  string `yaml:"last_name"`
	Age       int    `yaml:"age"`
}
