package main

import (
	"io/ioutil"
)

func main() {
	fileName := "./config.yaml"
	source, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	println(string(source))
}
