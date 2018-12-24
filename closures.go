package main

import "fmt"

func retunAnanymousFucntion() func(msg string) {
	return func(message string) {
		fmt.Println(message)
	}
}
func testMsg(msg string) string {
	return msg
}

func nextFunc() func(a int) int {
	return func(a int) int {
		a++
		return a
	}
}

func nextFunc1() func() int {
	a := 1
	return func() int {
		a++
		return a
	}
}
func main() {

	fmt.Println(testMsg("Hello World"))

	// anonymous function

	func(msg string) {
		fmt.Println(msg)
	}("I am Anonymous Fucntion")

	anonymousReturn := retunAnanymousFucntion()
	anonymousReturn("I am Returned from Ananymous Fucntion")

	incInt := nextFunc()
	incInt1 := nextFunc1()
	fmt.Println(incInt(8))
	fmt.Println(incInt(8))

	fmt.Println(incInt1())
	fmt.Println(incInt1())

	incInt2 := nextFunc1()
	fmt.Println(incInt2())
	fmt.Println(incInt2())

}
