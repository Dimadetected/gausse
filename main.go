package main

import (
	"fmt"
	"math"
)

var (
	f string = "tx^3+x^2"
)

func u(x, t float64) float64 {
	return t*math.Pow(x, 3) + math.Pow(x, 2)
}
func psi(x []float64) []float64 {
	f := []float64{}
	for i := range x {
		f = append(f, math.Pow(x[i], 2))
	}
	return f
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
	U := make([][]float64, 0)
	for i := 0; i < len(t); i++ {
		u := make([]float64, len(x))
		U = append(U, u)
	}
	U[0] = psi(x)

	for n := 1; n < n0+1; n++ {
		for m := m0 - n0 + n; m < m0+n0+1-n; m++ {
			U[n][m] = r*(U[n-1][m-1]+U[n-1][m+1]) + (1-2*r)*U[n-1][m] + tau*phi(x[m], t[n-1])
		}
	}
	fmt.Println("Функция:", f)

}
