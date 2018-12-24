package main

import "fmt"

func main() {

	i := 0
	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 0; j <= 5; j++ {
		fmt.Println(j)
	}

	for k := 0; k <= 10; k++ {

		if k%2 == 0 {
			continue
		}
		fmt.Println(k)
	}

	fmt.Println(i)

	i = 0

	//Infinite Loop
	for {

		if i == 7 {
			break
		}
		i++
		fmt.Println(i)
	}

	//Only For loop is allowed in Go
}
