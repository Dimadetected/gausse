package main

import "math"

var (
	f string = "tx^3+x^2"
)

func u(x, t float64) float64 {
	return t*math.Pow(x, 3) + math.Pow(x, 2)
}
func psi(x float64) float64 {
	return math.Pow(x, 2)
}
func phi(x, t float64) float64 {
	return math.Pow(x, 3) - 6*t*x - 2
}
func main() {
	h := 0.01
	tau := 0.03

	r := tau / math.Pow(h, 2)

	if r > 0.5 {
		panic("r <= 0.5. Условие устойчивости не выполнено.")
	}

	A := -10.0
	B := 10.0

	T := 0.5

	var x []float64
	var t []float64

	for i := A; i < B; i += h {
		x = append(x, i)
	}

	for i := 0.0; i < T; i += tau {
		t = append(t, i)
	}

	n0 := 5
	m0 := 5
	U :=
}
