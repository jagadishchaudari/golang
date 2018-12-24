package main

import "fmt"

func addSub(a, b int) (int, int) {
	return a + b, a - b
}
func main() {

	add, sub := addSub(5, 3)
	add1, _ := addSub(4, 3)
	fmt.Println(add1)
	fmt.Println(add, sub)
	fmt.Println(addSub(5, 4))

}
