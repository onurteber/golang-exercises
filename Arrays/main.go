package main

func main() {

	myArray1 := [3]int{}
	myArray1[0] = 320
	myArray1[1] = 230
	myArray1[2] = 540
	println(myArray1[0])
	println(len(myArray1))

	var colors = [3]string{"Red", "Green", "Blue"}
	println(colors[0])

	var myArray2 = [...]int{3, 5, 6, 8}
	println(myArray2[3])

	for i, value := range colors {
		println("i= ", i, " & value= ", value)
	}

}
