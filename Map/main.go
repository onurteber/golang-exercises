package main

func main() {

	// 1.
	myMap := make(map[int]string)
	myMap[1] = "Onur"
	myMap[2] = "Teber"
	for _, value := range myMap {
		print(value + " ")
	}
	print("\n")
	// 2.
	states := make(map[string]string)
	states["IST"] = "Istanbul"
	states["ANK"] = "Ankara"
	for _, value := range states {
		print(value + " ")
	}
	print("\n")
	println(states["ANK"])
	print("\n")

	delete(states, "ANK")
	for _, value := range states {
		print(value + " ")
	}
	print("\n")

	states["ERZ"] = "Erzurum"
	for _, value := range states {
		print(value + " ")
	}

	print("\n")

	// Create list, list capasite equals map length
	keys := make([]string, len(states))
	i := 0
	for k := range states {
		keys[i] = k
		i++
	}
	for i := range keys {
		print(states[keys[i]] + " ")
	}

}
