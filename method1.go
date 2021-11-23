package main

import (
	"fmt"
	"math"
)

func cos(x float64) float64 {
	return math.Cos(x)
}

func cosL(x []float64, p, q float64) []float64 {
	arr := []float64{}
	for i := range x {
		arr = append(arr, -math.Cos(x[i])-p*math.Sin(x[i])+q*math.Cos(x[i]))
	}
	return arr
}

func startMethod1() {
	p := 3.0
	q := -3.0

	a := 0.0
	b := math.Pi

	alpha := []float64{1, 1}
	betta := []float64{0, 0}
	gamma := []float64{cos(a), cos(b)}

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
	for i := 1; i < N; i++ {
		A[i] = (1 - h/2*p) / (h * h)
		B[i] = (-2 + h*h*q) / (h * h)
		C[i] = (1 + h/2*p) / (h * h)
	}
	B[0] = alpha[0] - betta[0]/h
	C[0] = betta[0] / h
	D[0] = gamma[0]

	A[N] = alpha[1] - betta[1]/h
	B[N] = betta[1] / h
	D[N] = gamma[1]

	y := traditionalSolve(A, B, C, D)
	for i := 0; i < N+1; i++ {
		fmt.Printf("x[%d] = %.10f\t y[%d] = %.10f\t f(x) = %.10f \t|\t%.5f\n", i, x[i], i, y[i], cos(x[i]), math.Abs(cos(x[i])-y[i]))
	}
}
