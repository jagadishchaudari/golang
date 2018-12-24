package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func add3(a, b, c int) int {
	return a + b + c
}
func main() {
	fmt.Println(add(1, 2))
	fmt.Println(add3(1, 2, 3))

}
