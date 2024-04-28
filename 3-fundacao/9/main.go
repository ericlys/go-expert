package main

import (
	"fmt"
)

func main() {
	total := func() int {
		return sum(234, 523, 1434525, 24, 2, 6, 2423, 4, 3) * 2
	}()

	fmt.Println(total)
}

func sum(numbers ...int) int {
	total := 0
	for _, n := range numbers {
		total += n
	}
	return total
}
