package main

import (
	"errors"
	"fmt"
)

func main() {
	// salarios := map[string]int{"eric": 9000, "vitor": 8500}
	// salarios["ericka"] = 8000

	// for name, value := range salarios {
	// 	println(name, value)
	// }
	valor, err := sum(5, 10)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(valor)
}

func sum(a, b int) (int, error) {
	if a+b >= 50 {
		return 0, errors.New("soma moior que 50")
	}
	return a + b, nil
}
