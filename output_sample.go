package main

import (
	"fmt"
)

func main() {
	var i, j string = "Hello", "World"

	fmt.Print(i, " ", j)
	fmt.Println(i, j)
	fmt.Printf("i has value: %v and type: %T\n", i, i)
}
