package main

import (
	"fmt"
)

func main() {
	adj := [2]string{"big", "tasty"}
	fruits := [3]string{"apple", "orange", "banana"}
	for i := 0; i < len(adj); i++ {
		for j := 0; j < len(fruits); j++ {
			fmt.Println(adj[i], fruits[j])
		}
	}
	for idx, val := range fruits {
		fmt.Printf("%v\t%v\n", idx, val)
	}
	for _, val := range fruits {
		fmt.Printf("%v\n", val)
	}
	for idx, _ := range fruits {
		fmt.Printf("%v\n", idx)
	}
}
