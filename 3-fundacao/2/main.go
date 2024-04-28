package main

import "fmt"

const a = "hello, world"

type ID int

var (
	c string
	b bool
	f ID      = 1
	e float64 = 1.2
)

func main() {
	s := []int{40, 20, 50}

	s = append(s, 110)

	fmt.Printf("len=%d cap=%d %v/n", len(s), cap(s), s)
}
