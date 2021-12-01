package main

import (
	"fmt"
	"github.com/crackcell/gotabulate"
	"math"
)

//const f = "tx^3+x^2"
const f = "tsinx+x^2"

//func u(x, t float64) float64 {
//	return t*math.Pow(x, 3) + math.Pow(x, 2)
//}
func u(x, t float64) float64 {
	return t*math.Sin(x) + math.Pow(x, 2)
}
func psi(x []float64) []float64 {
	var f []float64
	for i := range x {
		f = append(f, math.Pow(x[i], 2))
	}
	return f
}

//func phi(x, t float64) float64 {
//	return math.Pow(x, 3) - 6*t*x - 2
//}
func phi(x, t float64) float64 {
	return math.Sin(x) - (t*-math.Sin(x) + 2)
}

func main() {
	h := 0.1
	tau := 0.005

	r := tau / math.Pow(h, 2)

	A := 0.0
	B := 0.9

	T := 1.0

	var x []float64
	var t []float64

	for i := A; i < B; i += h {
		fmt.Println(A, B, i, h)
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
		d[1] = tau*phi(x[0], t[n]) + U[n-1][0]

		a[len(a)-1] = -r
		b[len(b)-1] = 2*r + 1
		c[len(c)-1] = 0
		d[len(d)-1] = tau*phi(x[len(x)-1], t[n]) + U[n-1][len(U[n-1])-1] + r*(t[n]+1)

		for m := 1; m < len(x)-1; m++ {
			a[m] = -r
			b[m] = 2*r + 1
			c[m] = -r
			d[m] = tau*phi(x[m], t[n]) + U[n-1][m]
		}
		l := traditionalSolve(a[1:], b[1:], c[1:], d[1:])

		for i := range l {
			U[n][1+i] = l[i]
		}

		tabulator := gotabulate.NewTabulator()
		tabulator.SetFirstRowHeader(true)
		tabulator.SetFormat("orgtbl")

		var printArr [][]string
		header := []string{"x", "u(n0, m0)", "u(x, t)", "Δ"}
		printArr = append(printArr, header)

		for j := range U[n] {
			strArr := []string{
				fmt.Sprintf("%.1f", x[j]),
				fmt.Sprintf("%.7f", u(x[j], t[n])),
				fmt.Sprintf("%.7f", U[n][j]),
				fmt.Sprintf("%.7f", math.Abs(u(x[j], t[n])-U[n][j])),
			}

			printArr = append(printArr, strArr)
		}
		fmt.Printf("Слой %d/%d\n", n, n0)
		fmt.Print(
			tabulator.Tabulate(
				printArr,
			))
		fmt.Println()
	}
	fmt.Println("Функция:", f)
	fmt.Println("Полученное занчение в выбранном узле:", U[n0][m0])
	fmt.Println("Точное решение в выбранном узле:", u(x[m0], t[n0]))
	fmt.Println("Погрешность вычислений :", math.Abs(u(x[m0], t[n0])-U[n0][m0]))
}

func traditionalSolve(a, b, c, d []float64) []float64 {
	n := len(a)

	m := make([]float64, n)

	P := make([]float64, n-1)
	Q := make([]float64, n-1)

	P[0] = -c[0] / b[0]
	Q[0] = d[0] / b[0]

	for k := 1; k < n-1; k++ {
		P[k] = -c[k] / (a[k]*P[k-1] + b[k])
		Q[k] = (d[k] - a[k]*Q[k-1]) / (a[k]*P[k-1] + b[k])
	}
	m[n-1] = (d[n-1] - a[n-1]*Q[n-2]) / (a[n-1]*P[n-2] + b[n-1])

	for i := n - 2; i > -1; i-- {
		m[i] = P[i]*m[i+1] + Q[i]
	}
	return m
}
