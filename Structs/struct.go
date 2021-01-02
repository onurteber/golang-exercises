package main

import "strconv"

type Human struct {
	FirstName string
	LastName  string
	Age       int
}

// Default and Empty Constructure Method
func NewHuman() *Human {
	h := new(Human)
	return h
}

func NewHumanWithFullParams(firstName, lastName string, age int) *Human {
	h := new(Human)
	h.FirstName = firstName
	h.LastName = lastName
	h.Age = age
	return h
}

func main() {

	// CREATE OBJECT INSTANCE

	//v1
	//x := Human{FirstName: "Onur", LastName: "Teber", Age: 23}
	//println(x.FirstName)

	//v2
	//x := new(Human)
	//x.FirstName = "Onur"
	//println(x.FirstName)

	//x := NewHuman()
	//x.Age = 23
	//println(x.Age)

	x := NewHumanWithFullParams("Onur", "Teber", 23)

	var data = x.FirstName + " " + x.LastName + " / " + strconv.Itoa(x.Age)

	println(data)

}
