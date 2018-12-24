package main

import "fmt"

func main() {

	//Maps are similar to Dictinaries and has table

	m := make(map[int]string)
	m[1] = "jagadish"
	m[2] = "chaudari"

	// printing the values of map
	fmt.Println(m)
	fmt.Println(m[1])
	fmt.Println(len(m)) // len of the map
	delete(m, 1)
	fmt.Println(m)
	m[1] = "JAGADISH"
	fmt.Println(m)
	m[2] = "JAGADISH"
	fmt.Println(m)

	m[1] = "jagadish"
	m[2] = "chaudari"
	fmt.Println(m)

	//Another way to initialise the map

	m2 := map[int]string{1: "Tanishka", 2: "Tanvi"}
	fmt.Println(m2)

	val, isPresent := m[3]
	if isPresent {
		fmt.Println(val)
	} else {
		fmt.Printf("Key is not present in the map\n")
	}

	//if we dont want to use val then _(underscore)
	_, isPresent1 := m[3]
	fmt.Println(isPresent1)
}
