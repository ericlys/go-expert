package main

import (
	"encoding/json"
	"os"
)

type Account struct {
	Number int `json:"n"`
	Total  int `json:"t"`
}

func main() {
	acc1 := Account{Number: 1, Total: 100}
	res, e := json.Marshal(acc1)
	if e != nil {
		panic(e)
	}

	println(string(res))

	//encoder ja entrega o valor para alguem
	encoder := json.NewEncoder(os.Stdout)
	e = encoder.Encode(acc1)
	if e != nil {
		panic(e)
	}

	jsonP := []byte(`{"n":2, "t":200}`)
	var accX Account
	e = json.Unmarshal(jsonP, &accX)
	if e != nil {
		panic(e)
	}

	println(accX.Total)
}
