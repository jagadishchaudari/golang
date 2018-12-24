package main

import "fmt"

func main() {

	//range is similar to range in python

	strArr := []string{"a", "b", "c", "d"}

	for c, char := range strArr {
		fmt.Println(c, char)
	}

	//can also range over key/value pairs

	m := map[int]string{1: "a", 2: "b"}

	for k, v := range m {
		fmt.Println("Key:", k, "Val:", v)
	}

	//just print the keys of a map
	for k := range m {
		fmt.Println("Key:", k)
	}

}
