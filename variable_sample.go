package main

import (
	"fmt"
)

const PI = 3.14

func main() {
	var student1 string = "John" //type is string
	var student2 = "Jane"        //type is inferred
	x := 2                       //type is inferred

	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(x)
	fmt.Println(PI)
}
