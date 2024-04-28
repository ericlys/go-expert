package main

import "fmt"

type Client struct {
	name  string
	coins int
}

func (c *Client) wolk() {
	c.name = "Ericlys Moreira"
	fmt.Println("O client", c.name, "andou")
}

func newClient(name string) *Client {
	return &Client{name: name, coins: 0}
}

func main() {
	p1 := Client{
		name: "Ericlys",
	}
	p1.wolk()
	fmt.Println(p1.name)

	p2 := newClient("Ericka")
	println(p2.name)
}
