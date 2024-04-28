package main

import "fmt"

type Address struct {
	city   string
	number int
	state  string
}

type People interface {
	Disable()
}

type Client struct {
	name    string
	year    int
	active  bool
	address Address
}

func (c Client) Disable() {
	c.active = false
	fmt.Println("Cliente", c.name, "desativado com sucesso!")

}

func Disability(p People) {
	p.Disable()
}

func main() {
	p := Client{
		name:   "Eric",
		year:   26,
		active: true,
	}

	Disability(p)
}
