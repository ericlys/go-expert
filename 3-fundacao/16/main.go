package main

type MyNumber int

type Number interface {
	~int | ~float64
}

func Soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func Compara[T Number](a T, b T) bool {
	if a > b {
		return true
	}
	return false
}

func main() {
	m := map[string]int{"Eric": 9000, "Ericka": 9000}
	m2 := map[string]float64{"Eric": 9000.99, "Ericka": 9000.3}

	m3 := map[string]MyNumber{"Eric": 90009, "Ericka": 90003}

	println(Soma(m))
	println(Soma(m2))
	println(Soma(m3))
	println(Compara(11, 10.8))
}
