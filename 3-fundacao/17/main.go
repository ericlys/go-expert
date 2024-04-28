package main

import (
	"fmt"

	"github.com/ericlys/curso-go/matt"
)

func main() {
	soma := matt.Sum(10, 20)
	fmt.Println("Resultado:", soma)
	fmt.Println(matt.A)

	car := matt.Carro{Marca: "vw"}
	fmt.Println(car.Marca)
}
