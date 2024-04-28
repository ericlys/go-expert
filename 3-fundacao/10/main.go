package main

import "fmt"

type Client struct {
	name   string
	year   int
	active bool
}

func main() {
	p := Client{
		name:   "Eric",
		year:   26,
		active: true,
	}

	p.active = false

	fmt.Println(p.name, p.active)
}
