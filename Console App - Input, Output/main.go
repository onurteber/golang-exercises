package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	print("Enter text: ")
	str, _ := reader.ReadString('\n')
	println(str)

	print("Enter number: ")
	str, _ = reader.ReadString('\n')
	f, err := strconv.ParseFloat(strings.TrimSpace(str), 64)
	if err != nil {
		println(err)
	} else {
		println("Number: ", f)
	}

}
