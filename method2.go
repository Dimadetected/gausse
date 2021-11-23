package main

import (
	"fmt"
	"math"
)

func startMethod2() {
	p := 3.0
	q := -5.0

	a := 0.0
	b := math.Pi

	N := 32
	h := (b - a) / float64(N)

	x := make([]float64, 0)
	for i := a; i < b; i += 0.09817477 {
		x = append(x, i)
	}

	A := make([]float64, N+1)
	B := make([]float64, N+1)
	C := make([]float64, N+1)

	D := cosL(x, p, q)
	for i := range D {
		D[i] *= h * h
	}

	for i := 1; i < N; i++ {
		A[i] = 1 - h/2*p
		B[i] = -2 + math.Pow(h, 2)*q
		C[i] = 1 + h/2*p
	}
	B[0] = 1
	C[0] = -1
	D[0] = 0 * h

	A[N] = 1
	B[N] = -1
	D[N] = 0 * h

	y := traditionalSolve(A, B, C, D)

	for i := 0; i < N+1; i++ {
		fmt.Printf("x[%d] = %.10f\t y[%d] = %.10f\t f(x) = %.10f \t|\t%.5f\n", i, x[i], i, y[i], cos(x[i]), math.Abs(cos(x[i])-y[i]))
	}
}
