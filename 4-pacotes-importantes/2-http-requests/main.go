package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get("https://www.google.com")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close() //defer - para n esquecer de fechar no final de tudo

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// Print the response body as a string
	fmt.Println(string(body))

}
