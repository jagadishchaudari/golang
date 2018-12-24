package main

import "fmt"

type employee struct {
	firstName string
	lastName  string
	id        int
}

func main() {

	fmt.Println(employee{firstName: "jagadish", lastName: "Chaudari", id: 1})
	fmt.Println(employee{"jagad", "Chaud", 2})
	fmt.Println(employee{firstName: "Sur", lastName: "Chaud"})

	emp := employee{firstName: "Lavanya", lastName: "Chaudari"}
	fmt.Println(emp)

	empPtr := &emp
	fmt.Println(empPtr.firstName)
	empPtr.id = 3

	fmt.Println(*empPtr)

}
