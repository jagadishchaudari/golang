package main

import "fmt"

func main() {

	var intArr [3]int
	fmt.Println(intArr)

	var bolArr []bool
	fmt.Println(bolArr)

	var strArr []string
	fmt.Println(strArr)

	// Accessing array elements using indices
	intArr[0] = 10
	intArr[2] = 20

	fmt.Println(intArr)

	// Initialising the one dimensional array
	bigArr := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(bigArr)

	fmt.Println(len(intArr))
	fmt.Println(len(bolArr))
	fmt.Println(len(strArr))
	fmt.Println(len(bigArr))

	//Mutli Dimensional Arrays

	mulArr := [2][3]int{{1, 2, 3},
		{4, 5, 6}}

	fmt.Println(mulArr)

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Println(mulArr[i][j])
		}
	}
}
