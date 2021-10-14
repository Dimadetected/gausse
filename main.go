package main

import (
	"fmt"
	"github.com/rgeoghegan/tabulate"
	"math"
)

type Print struct {
	Iter   float64
	Step   float64
	Value  float64
	U      float64
	UValue float64
	V      float64
	VValue float64
}

var (
//eps = 1e-6
)

//dichotomy
func dichotomy(f func(ksi float64) float64, a, b, eps float64) float64 {
	//if f(a)*f(b) < 0 {
	//	panic("dichotomy f(a) * f(b) < 0")
	//}
	//if eps > 0 {
	//	panic("dichotomy eps > 0")
	//}

	if a >= b {
		a, b = b, a
	}

	k := 0
	for (b-a)/2 > eps {
		c := (a + b) / 2
		fmt.Println(a, b, f(a), f(c))
		if f(a)*f(c) <= 0 {
			b = c
		} else {
			a = c
		}
		k += 1
	}
	return (a + b) / 2
}

//functions
func fFunc(x, u, v float64) float64 {
	return u - v + math.Exp(x)*(1+x*x)
}
func gFunc(x, u, v float64) float64 {
	return u + v + x + math.Exp(x)
}
func uFunc(x float64) float64 {
	return x * math.Exp(x)
}
func vFunc(x float64) float64 {
	return x * x * math.Exp(x)
}

//функция находит приближенное значение решения диф задачи в точке х
func step(x, u, v, h float64, fFunc, gFunc func(float64, float64, float64) float64) (float64, float64) {

	phi0 := h * fFunc(x, u, v)
	psi0 := h * gFunc(x, u, v)
	fmt.Println("phi,psi", fFunc(x, u, v), gFunc(x, u, v))
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
func jump(x, u, v, h0 float64, fFunc, gFunc func(float64, float64, float64) float64, eps float64) (float64, float64, float64) {
	h := h0
	sU1, sV1 := step(x, u, v, h, fFunc, gFunc)
	sU05, sV05 := step(x, u, v, h/2, fFunc, gFunc)
	sU2, sV2 := step(x+h/2, sU05, sV05, h/2, fFunc, gFunc)
	for math.Abs(sU1-sU2) > eps && math.Abs(sV1-sV2) > eps {
		h = h / 2
		sU1 = sU05
		sV1 = sV05

		sU05, sV05 = step(x, u, v, h/2, fFunc, gFunc)
		sU2, sV2 = step(x+(h/2), sU05, sV05, h/2, fFunc, gFunc)
	}
	return h, sU1, sV1
}
func rungeKutta(a, b, u, v, h0 float64, fFunc, gFunc func(float64, float64, float64) float64, eps float64, debug bool) (float64, float64) {
	iter := 1.0
	x := a
	answ := []*Print{}
	var h float64
	for x < b {
		h, u, v = jump(x, u, v, h0, fFunc, gFunc, eps)
		fmt.Println("huv:", h, u, v)
		x = x + h
		answ = append(answ, &Print{
			Iter:   iter,
			Step:   h,
			Value:  x,
			U:      u,
			UValue: uFunc(x) - u,
			V:      v,
			VValue: vFunc(x) - v,
		})
		iter += 1
		if b-x < h0 {
			h0 = b - x
		}
	}

	layout := &tabulate.Layout{Format: tabulate.GridFormat}
	table, err := tabulate.Tabulate(answ, layout)
	if err != nil {
		panic(err)
	}
	if debug {
		fmt.Println(table)
	}
	return u, v
}

func F(ksi float64, D []float64, h float64, alpha, betta, gamma []float64, f, g func(float64, float64, float64) float64) float64 {
	v := (gamma[0] - alpha[0]*ksi) / betta[0]
	u, v := rungeKutta(D[0], D[1], ksi, v, h, f, g, 1e-5, false)
	fmt.Println("u,v: ", u, v)
	return alpha[0]*u + betta[1]*v - gamma[1]
}
func main() {
	eps := 1e-9
	D := []float64{0, 1}
	h := 0.3

	alpha := []float64{0, 1}
	betta := []float64{1, 2}
	gamma := []float64{alpha[0]*uFunc(D[0]) + betta[0]*vFunc(D[0]), alpha[1]*uFunc(D[1]) + betta[0]*vFunc(D[1])}
	ksiInput := func(ksi float64) (float64, float64) {
		fKsi := F(ksi, D, h, alpha, betta, gamma, fFunc, gFunc)
		fmt.Printf("F(ξ: %.2f) = %.2f\n", ksi, fKsi)
		return ksi, fKsi
	}
	ksi1, fKsi1 := ksiInput(1)
	ksiInputIter := -1.0
	ksi2, fKsi2 := 0.0, 0.0
	for {
		ksi2, fKsi2 = ksiInput(ksiInputIter)
		if fKsi1*fKsi2 < 0 {
			break
		}
		ksi1, fKsi1 = ksi2, fKsi2
		ksiInputIter += 1
	}
	FDihotStr = FDih{D, h, alpha, betta, gamma, fFunc, gFunc}
	f := func(ksi float64) float64 {
		return F(ksi, FDihotStr.D, FDihotStr.h, FDihotStr.alpha, FDihotStr.betta, FDihotStr.gamma, FDihotStr.f, FDihotStr.g)
	}
	ksi := dichotomy(f, ksi1, ksi2, eps/1e+1)
	fmt.Printf("Поиск корня на отрезке [%.2f,%.2f]\n", math.Min(ksi1, ksi2), math.Max(ksi1, ksi2))
	fmt.Printf("ξ = %f \t\t F(ξ) = %f", ksi, f(ksi))
	fmt.Printf("Запуск автоматической стрельбы...")
	u := ksi
	v := (gamma[0] - alpha[0]*ksi) / betta[0]
	u, v = rungeKutta(D[0], D[1], u, v, h, fFunc, gFunc, eps, false)
	fmt.Println(u, v)
}
func FDihot(ksi float64) float64 {
	return F(ksi, FDihotStr.D, FDihotStr.h, FDihotStr.alpha, FDihotStr.betta, FDihotStr.gamma, FDihotStr.g, FDihotStr.f)
}

var FDihotStr FDih

type FDih struct {
	D     []float64
	h     float64
	alpha []float64
	betta []float64
	gamma []float64
	g     func(float64, float64, float64) float64
	f     func(float64, float64, float64) float64
}
