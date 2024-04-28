package main

import (
	"fmt"

	"github.com/ericlys/go-expert/7-packaging/1/math"
)

func main() {
	m := math.NewMath(2, 3)
	fmt.Println(m.Add())
}
