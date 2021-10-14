package main

import (
	"fmt"
	"github.com/rgeoghegan/tabulate"
	"math"
)

type Print struct {
	Step   float64
	Value  float64
	U      float64
	UValue float64
	V      float64
	VValue float64
}

var (
	eps = 1e-5
)

func fFunc(x, u, v float64) float64 {
	return u - v + math.Exp(x)*(1+x*x)
}
func gFunc(x, u, v float64) float64 {
	return u + v + x*math.Exp(x)
}
func uFunc(x float64) float64 {
	return x * math.Exp(x)
}
func vFunc(x float64) float64 {
	return x * x * math.Exp(x)
}

//функция находит приближенное значение решения диф задачи в точке х
func step(x, u, v, h float64) (float64, float64) {

	phi0 := h * fFunc(x, u, v)
	psi0 := h * gFunc(x, u, v)

	phi1 := h * fFunc(x+h/2, u+phi0/2, v+psi0/2)
	psi1 := h * gFunc(x+h/2, u+phi0/2, v+psi0/2)

	phi2 := h * fFunc(x+h/2, u+phi1/2, v+psi1/2)
	psi2 := h * gFunc(x+h/2, u+phi1/2, v+psi1/2)

	phi3 := h * fFunc(x+h, u+phi2, v+psi2)
	psi3 := h * gFunc(x+h, u+phi2, v+psi2)
	u = u + (phi0+2*phi1+2*phi2+phi3)/6
	v = v + (psi0+2*psi1+2*psi2+psi3)/6
	return u, v
}

/*
   Функция формирует сетку и в узлах этой сетки находит приближенные значения
    (находит шаг и приближ значение с точностью eps)
*/
func jump(x, u, v, h0 float64) (float64, float64, float64) {
	h := h0
	sU1, sV1 := step(x, u, v, h)
	sU05, sV05 := step(x, u, v, h/2)
	sU2, sV2 := step(x+h/2, sU05, sV05, h/2)
	fmt.Println(h, math.Abs(sU1-sU2) > eps, math.Abs(sV1-sV2) > eps, math.Abs(sV1-sV2), sV1, sV2)
	for math.Abs(sU1-sU2) > eps && math.Abs(sV1-sV2) > eps {
		fmt.Println(h, math.Abs(sU1-sU2) > eps, math.Abs(sV1-sV2) > eps)
		h = h / 2
		sU1 = sU05
		sV1 = sV05

		sU05, sV05 = step(x, u, v, h/2)
		sU2, sV2 = step(x+(h/2), sU05, sV05, h/2)
	}
	return h, sU1, sV1
}
func main() {
	h0 := 0.3

	x0 := 0.0
	xn := 1.0

	u0 := 0.0
	v0 := 0.0

	x := x0
	u := u0
	v := v0

	answ := []*Print{}
	for x < xn {
		h, uN, vN := jump(x, u, v, h0)
		x = x + h
		u = uN
		v = vN
		answ = append(answ, &Print{
			Step:   h,
			Value:  x,
			U:      u,
			UValue: uFunc(x) - u,
			V:      v,
			VValue: vFunc(x) - v,
		})

		if xn-x < h0 {
			h0 = xn - x
		}
	}

	layout := &tabulate.Layout{Format: tabulate.GridFormat}
	table, err := tabulate.Tabulate(answ, layout)
	if err != nil {
		panic(err)
	}
	fmt.Println(table)
}
