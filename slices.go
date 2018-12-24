package main

import "fmt"

func main() {

	s := make([]int, 3)

	// are similar to Arrays
	fmt.Println(s)
	count := 1
	for i := 0; i < 3; i++ {
		s[i] = count
		fmt.Println(s[i])
		count++
	}

	fmt.Println(len(s))

	//append functions can be used in slices

	s = append(s, 4, 5, 6, 7)
	fmt.Println(s)

	//slice syntax
	fmt.Println(s[0:3])
	fmt.Println(s[1:3])
	fmt.Println(s[:5])
	fmt.Println(s[4:])

	//concise slice definition
	s1 := []int{100, 200, 300}
	fmt.Println(s1)

	//t1 := s
	//fmt.Println(t1)
	//t1[0] = 200
	//fmt.Println(t1)
	//fmt.Println(s)

	//use copy to avoid changing th orginal slice S

	t2 := make([]int, len(s))
	copy(t2, s)
	fmt.Println(t2)
	t2[0] = 200
	fmt.Println(t2)
	fmt.Println(s)

	//multi dimensional slices

	ss := make([][]int, 3)

	fmt.Println(ss)

	//U can keep the size of the inner slices varaible
	for i := 0; i < 3; i++ {
		//innLen := i + 1
		ss[i] = make([]int, 4)
		for j := 0; j < 4; j++ {
			ss[i][j] = i + j

		}
	}

	fmt.Println(ss)

}
