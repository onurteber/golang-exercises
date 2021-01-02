package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	_, err := os.Open("text.rar")
	checkError(err)
	/*i := -2
	if i < 0 {
		return 0, fmt.Errorf("hata: %g", i)
	}*/

	/*myError := errors.New("this is an error")
	println(myError.Error())*/

}

func checkError(err error) {
	if err != nil {
		fmt.Println("Error: ", err.Error())
		log.Fatal("Error: ", err)
		os.Exit(1)
	}
}
