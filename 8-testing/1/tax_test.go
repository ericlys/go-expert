package tax

import (
	"testing"
)

func TestCalculateTax(t *testing.T) {
	amount := 500.0
	expected := 5.0

	result := CalculateTax(amount)

	if result != expected {
		t.Errorf("Expected %f but got %f", expected, result)
	}
}

func TestCalculateTaxBatch(t *testing.T) {
	type calcTax struct {
		amount, expect float64
	}

	table := []calcTax{
		{500.0, 5.0},
		{1000.0, 10.0},
		{1500.0, 10.0},
		{0.0, 0.0},
	}

	for _, item := range table {
		result := CalculateTax(item.amount)
		if result != item.expect {
			t.Errorf("Expected %f but got %f", item.expect, result)
		}
	}
}

func BenchmarkCalculateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax(1000.0)
	}
}

func BenchmarkCalculateTax2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax2(1000.0)
	}
}

func FuzzCalculateTax(f *testing.F) {
	seed := []float64{-1, -2, -2.5, 500.0, 1000.0, 1501.0}
	for _, amount := range seed {
		f.Add(amount)
	}
	f.Fuzz(func(t *testing.T, a float64) {
		result := CalculateTax(a)
		if a <= 0 && result != 0 {
			t.Errorf("Reveived %f but expected 0", result)
		}
		if a > 20000 && result != 20 {
			t.Errorf("Reveived %f but expected 20", result)
		}
	})

}

//comands
//go test -coverprofile=coverage.out
// go tool cover -html=coverage.out
//go test -bench=.
//go test -bench=. -run=^# -count=10 -benchtime=3s
//-benchmen  -alocacao de memoria
