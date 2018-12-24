package main

import "fmt"

func main() {

	i := 2

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	j := 8

	switch j {

	case 1, 2, 3:
		fmt.Println(" case one, two, or three")
	default:
		fmt.Println("Not one, two, or three")
	}

	switch {
	case j == 11:
		fmt.Println("J is equal to 11")
	case j < 12 && j > 9:
		fmt.Println(" j is less than 12 and greater than 9")
		//default:
		//fmt.Println(" j is less than 5")

	}

}
