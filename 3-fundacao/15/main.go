package main

import "fmt"

func main() {
	var va1 interface{} = "Ericlys Moreira"
	println(va1.(string))
	res, ok := va1.(int)
	fmt.Printf("O valor de res é %v e o resultado ok é %v", res, ok)

	res2, ok2 := va1.(string)
	if ok == true {
		fmt.Printf("O valor de res2 é %v /n", res2)
	} else {
		fmt.Printf("Erro ao converter va1 para string: %v", ok2)
	}

}
