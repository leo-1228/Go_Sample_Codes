package main

import (
	"fmt"
)

func main() {
	var (
		sum1 = 100 + 50    // 150 (100 + 50)
		sum2 = sum1 + 250  // 400 (150 + 250)
		sum3 = sum2 + sum2 // 800 (400 + 400)
	)
	fmt.Println(sum3)
	var x = 10
	x += 5
	fmt.Println(x)
	var y = 3
	fmt.Println(x > y) // returns 1 (true) because 5 is greater than 3
}
