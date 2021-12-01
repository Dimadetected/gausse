package main

import (
	"fmt"
	"github.com/crackcell/gotabulate"
	"math"
)

const f = "tx^3+x^2"

func u(x, t float64) float64 {
	return t*math.Pow(x, 3) + math.Pow(x, 2)
}
func psi(x []float64) []float64 {
	var f []float64
	for i := range x {
		f = append(f, math.Pow(x[i], 2))
	}
	return f
}
func phi(x, t float64) float64 {
	return math.Pow(x, 3) - 6*t*x - 2
}

func main() {
	h := 0.1
	tau := 0.003

	r := tau / math.Pow(h, 2)

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

	n0 := 4
	m0 := 5
	U := make([][]float64, 0)
	for i := 0; i < len(t); i++ {
		u := make([]float64, len(x))
		U = append(U, u)
	}
	U[0] = psi(x)

	a := make([]float64, len(x))
	b := make([]float64, len(x))
	c := make([]float64, len(x))
	d := make([]float64, len(x))

	for n := 1; n < n0+1; n++ {
		a[1] = 0
		b[1] = 2*r + 1
		c[1] = -r
		d[1] = tau * phi(x[0],t[n]) + U[n-1][0]

		a[len(a)-1] = -r
		b[len(b)-1] = 2 * r +1
		c[len(c)-1] = 0
		d[len(d)-1] = tau * phi(x[len(x)-1],t[n]) + U[n-1][len(U[n-1])-1] + r * (t[n]+1)

		for m := m0 - n0 + n; m < m0+n0+1-n; m++ {
			a[m] = -r
			b[m] = 2 * r +1
			c[m] = -r
			d[m] = tau * phi(x[m],t[n]) + U[n-1][m]
		}
		l :=
		U[n][1] = r*(U[n-1][m-1]+U[n-1][m+1]) + (1-2*r)*U[n-1][m] + tau*phi(x[m], t[n-1])

	}
	fmt.Println("Функция:", f)

	var printArr [][]string
	var header []string

	for m := m0 - n0; m < m0+n0+1; m++ {
		header = append(header, fmt.Sprintf("m=%d", m))
	}
	printArr = append(printArr, header)
	for i := 0; i < n0+1; i++ {
		var strArr []string
		for j := m0 - n0; j < m0+n0+1; j++ {
			strArr = append(strArr, fmt.Sprintf("%.7f", U[i][j]))
		}
		printArr = append(printArr, strArr)
	}
	tabulator := gotabulate.NewTabulator()
	tabulator.SetFirstRowHeader(true)
	tabulator.SetFormat("orgtbl")
	fmt.Print(
		tabulator.Tabulate(
			printArr,
		))
	fmt.Println("Точное решение в выбранном узле:", u(x[m0], t[n0]))
	fmt.Println("Погрешностьвычислений :", math.Abs(u(x[m0], t[n0]))-U[n0][m0])
}
