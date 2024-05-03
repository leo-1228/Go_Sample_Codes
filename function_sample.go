package main

import (
	"fmt"
)

func familyName(fname string, age int) {
	fmt.Println("Hello", age, "year old", fname, "Refsnes")
}

func myFunction1(x int, y int) (result int) {
	result = x + y
	return
}

func myFunction(x int, y int) int {
	return x + y
}
func myFunction2(x int, y string) (result int, txt1 string) {
	result = x + x
	txt1 = y + " World!"
	return
}

func main() {
	familyName("Liam", 3)
	fmt.Println(myFunction(1, 2))
	total := myFunction(1, 2)
	fmt.Println(total)
	fmt.Println(myFunction1(1, 2))
	fmt.Println(myFunction2(5, "Hello"))
	a, b := myFunction2(5, "Hello")
	fmt.Println(a, b)
	_, c := myFunction2(5, "Hello")
	fmt.Println(c)
}
