package main

import "strconv"

func main() {

	numbers := make([]int, 5, 5)
	numbers[0] = 123
	numbers[1] = 456
	numbers[2] = 789
	numbers[3] = 987
	numbers[4] = 654
	for _, value := range numbers {
		print(strconv.Itoa(value) + " ")
	}

	print("\n")

	numbers = append(numbers, 358)
	for _, value := range numbers {
		print(strconv.Itoa(value) + " ")
	}
	print("\n")
	println(cap(numbers))
	println(len(numbers))

	print("\n")

	var colors = []string{"Red", "Green", "Blue"}
	colors = append(colors, "Yellow")
	for _, value := range colors {
		print(value + " ")
	}

	print("\n")

	colors = append(colors[1:len(colors)])
	for _, value := range colors {
		print(value + " ")
	}

	print("\n")

	colors = append(colors[:len(colors)-1])
	for _, value := range colors {
		print(value + " ")
	}

	/*myArray := [...]int{15, 62, 35}

	mySlice := myArray[2:]
	println(mySlice[0])*/

}
