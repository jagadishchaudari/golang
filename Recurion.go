package main

import "fmt"

func fact(number int) int {
	if number <= 1 {
		return 1
	}
	return number * fact(number-1)
}

func main() {
	fmt.Println(fact(5))

}
