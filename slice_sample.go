package main

import (
	"fmt"
)

func main() {
	myslice1 := []int{}
	fmt.Println(len(myslice1))
	fmt.Println(cap(myslice1))
	fmt.Println(myslice1)

	myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
	fmt.Println(len(myslice2))
	fmt.Println(cap(myslice2))
	fmt.Println(myslice2)
	arr1 := []int{10, 11, 12, 13, 14, 15}
	arr1 = append(arr1, 20, 21, 22, 23)
	myslice := arr1[2:4]

	fmt.Printf("myslice = %v\n", myslice)
	fmt.Printf("length = %d\n", len(myslice))
	fmt.Printf("capacity = %d\n", cap(myslice))

	myslice11 := make([]int, 5, 10)
	fmt.Printf("myslice1 = %v\n", myslice11)
	fmt.Printf("length = %d\n", len(myslice11))
	fmt.Printf("capacity = %d\n", cap(myslice11))

	// with omitted capacity
	myslice21 := make([]int, 5)
	fmt.Printf("myslice2 = %v\n", myslice21)
	fmt.Printf("length = %d\n", len(myslice21))
	fmt.Printf("capacity = %d\n", cap(myslice21))

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	neededNumbers := numbers[:len(numbers)-10]
	numbersCopy := make([]int, len(neededNumbers))
	copy(numbersCopy, neededNumbers)

	fmt.Printf("numbersCopy = %v\n", numbersCopy)
}
