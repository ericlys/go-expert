package main

func soma(a, b *int) int {
	*a = 3
	*b = 8000

	return *a + *b
}

func main() {
	va1 := 10
	va2 := 20

	soma(&va1, &va2)
	println(va1, va2)

}
