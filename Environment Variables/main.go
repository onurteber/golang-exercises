package main

import "os"

func main() {

	uName := os.Getenv("USERNAME")
	uDomain := os.Getenv("USERDOMAIN")
	processorArchitecture := os.Getenv("PROCESSOR_ARCHITECTURE")
	processorIdentifier := os.Getenv("PROCESSOR_IDENTIFIER")
	processorLevel := os.Getenv("PROCESSOR_LEVEL")
	goRoot := os.Getenv("GOROOT")
	goPath := os.Getenv("GOPATH")
	homePath := os.Getenv("HOMEPATH")

	println("Username: " + uName)
	println("Domain: " + uDomain)
	println("İşlemci Mimarisi: " + processorArchitecture)
	println("İşlemci ID: " + processorIdentifier)
	println("İşlemci Seviyesi: " + processorLevel)
	println("HOMEPATH: " + homePath)
	println("GOPATH: " + goPath)
	println("GOROOT: " + goRoot)
	println("HOMEPATH: " + homePath)
}
