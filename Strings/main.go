package main

import (
	"bytes"
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type Option struct {
	length    int
	upperCase bool
	lowerCase bool
	numbers   bool
	specials  bool
}

var source = rand.NewSource(time.Now().UnixNano())

const _charsetLowercase = "abcdefghijklmnopqrstuvwxyz"
const _charsetUppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const _charsetNumbers = "0123456789"
const _charsetSpecials = "?*-+%&$"

func GeneratePassword(opt Option) (string, error) {
	x := make([]byte, opt.length)
	charset := "."

	if opt.upperCase {
		charset += _charsetUppercase
	}
	if opt.lowerCase {
		charset += _charsetLowercase
	}
	if opt.specials {
		charset += _charsetSpecials
	}
	if opt.numbers {
		charset += _charsetNumbers
	}

	if charset == "." {
		return "NON-GENERATED!", errors.New("En az bir karakter seçimi yapmak zorundasınız!")
	}

	for i := range x {
		x[i] = charset[source.Int63()%int64(len(charset))]
	}

	return string(x), nil

}

func main() {

	// Generate Password
	password, err := GeneratePassword(Option{length: 15, upperCase: false, lowerCase: false, numbers: false, specials: false})
	if err == nil {
		println(password)
	} else {
		println(err.Error())
	}

	// String Builder
	builder := strings.Builder{}

	for i := 0; i < 10; i++ {
		builder.WriteString("Data ")
	}

	result := builder.String()
	fmt.Println(result)

	// String Buffer
	var x bytes.Buffer
	for i := 0; i < 100; i++ {
		pass, err := GeneratePassword(Option{length: 4, upperCase: true, lowerCase: true, numbers: false, specials: true})
		if err == nil {
			x.WriteString(pass)
		} else {
			println(err.Error())
		}
	}
	fmt.Println(x.String())

}
