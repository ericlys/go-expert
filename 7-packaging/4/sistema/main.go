package main

import (
	"fmt"

	"github.com/ericlys/go-expert/7-packaging/3/math"
	"github.com/google/uuid"
)

func main() {
	m := math.NewMath(2, 3)
	fmt.Println(uuid.New().String())
	fmt.Println(m.Add())
}

//go mod tidy -e --ignorar os pacotes n publicados
