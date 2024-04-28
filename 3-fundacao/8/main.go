package main

import (
	"fmt"
)

func main() {

	fmt.Println(sum(1, 2, 435, 436, 22))
}

func sum(numbers ...int) int {
	total := 0
	for _, n := range numbers {
		total += n
	}
	return total
}
