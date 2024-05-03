package main

import (
	"fmt"
)

func main() {
	a := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}

	var b []string // defining the order
	b = append(b, "one", "two", "three", "four")

	for k, v := range a { // loop with no order
		fmt.Printf("%v : %v, ", k, v)
	}

	fmt.Println()

	for _, element := range b { // loop with the defined order
		fmt.Printf("%v : %v, ", element, a[element])
	}

	d := map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964", "day": ""}

	val1, ok1 := d["brand"] // Checking for existing key and its value
	val2, ok2 := d["color"] // Checking for non-existing key and its value
	val3, ok3 := d["day"]   // Checking for existing key and its value
	_, ok4 := d["model"]    // Only checking for existing key and not its value

	fmt.Println(val1, ok1)
	fmt.Println(val2, ok2)
	fmt.Println(val3, ok3)
	fmt.Println(ok4)

	var c = make(map[string]string)
	c["brand"] = "Ford"
	c["model"] = "Mustang"
	c["year"] = "1964"

	fmt.Println(c)

	delete(c, "year")

	fmt.Println(c)
}
