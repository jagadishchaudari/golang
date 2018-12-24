package main

import "fmt"

func main() {

	i := 5

	if i%2 == 0 {
		fmt.Println(i, "is even")
	} else {
		fmt.Println(i, "is odd")
	}

	k := 10
	if k == 10 {
		fmt.Println("K value is 10")
	}

	j := 5

	if j < 50 {
		fmt.Println("j value is less than 50")
	} else if j > 50 {
		fmt.Println("j value is greater than 50")
	} else {
		fmt.Println("K value is equal to 50")
	}

	//Ternary Operators are not allowed in Go
	// a ? b: c
}
