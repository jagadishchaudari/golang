package main

import "fmt"

func varArg(nums ...int) int {
	total := 1
	for _, num := range nums {
		total *= num
	}

	return total
}
func main() {
	//Println is an example of variadic fucntion
	fmt.Print(1, 2, 3, 4, 5, 6)

	fmt.Println(varArg(1, 2, 3))
	fmt.Println(varArg(1, 2, 3, 4))
	fmt.Println(varArg(1, 2, 3, 4, 5))

	//varidic functions to slices

	slicArr := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(varArg(slicArr...))
}
