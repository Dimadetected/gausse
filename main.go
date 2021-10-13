package main

import (
	"fmt"
	"github.com/rgeoghegan/tabulate"
	"math"
	"strconv"
	"time"
)

type Print struct {
	Step float64
	X    float64
	Y    float64
	Err  float64
}

var (
	eps = 1e-6
	h0  = 0.3
	x   = 0.0
	X   = 1.0
	y   = 0.0
)

func f(x, y float64) float64 {
	return math.Exp(x) + y
}

func step(x, y, h float64) float64 {
	g0 := h * f(x, y)
	g1 := h * f(x+h/2, y+g0/2)
	g2 := h * f(x+h/2, y+g1/2)
	g3 := h * f(x+h, y+g2)
	deltaY := (g0 + 2*g1 + 2*g2 + g3) / 6
	return y + deltaY
}

func jump(x, y float64) (float64, float64) {
	h := h0
	yh := step(x, y, h)
	for {
		yh0 := step(x, y, h/2)
		yh1 := step(x+h/2, yh0, h/2)
		if math.Abs(yh-yh1) <= eps {
			return yh, h
		}
		h = h / 2
		yh = yh0
	}
}
func main() {
	var h float64
	res := make([]float64, 0)
	i := 0.0
	prints := []*Print{}
	for X-x > eps {
		err := math.Abs(math.Exp(x)*x - y)
		res = append(res, i, x, y, err)
		prints = append(prints, &Print{i, x, y, err})
		y, h = jump(x, y)
		x += h
		if X-x < h0 {
			h0 = X - x
		}
		i += 1
	}
	err := math.Abs(math.Exp(x)*x - y)
	res = append(res, i, x, y, err)
	prints = append(prints, &Print{i, x, y, err})

	layout := &tabulate.Layout{Format: tabulate.GridFormat}
	table, erro := tabulate.Tabulate(prints, layout)
	if erro != nil {
		panic(err)
	}
	fmt.Println(table)
	return
	//A := [][]float64{
	//	{1, 0.5, 0.33333333, 0.25, 0.2},
	//	{0.5, 0.33333333, 0.25, 0.2, 0.16666667},
	//	{0.33333333, 0.25, 0.2, 0.16666667, 0.14285714},
	//	{0.25, 0.2, 0.16666667, 0.14285714, 0.125},
	//	{0.2, 0.16666667, 0.14285714, 0.125, 0.11111111},
	//}
	//gradient(A)
	//return

	var a string
	//fmt.Println("Введите \n 1)Если нужно обычную матрицу \n 2)Если нужно вырожденную матрицу")
	//fmt.Scan(&a)
	a = "1"
	if a == "1" {

		m := NewMatrix(5, 5)
		//m.valueRand()
		//fmt.Println("Обычная матрица: ")
		m.value = [][]float64{
			{11, 12, 2, 14, 1},
			{3, 10, 5, 1, 0},
			{14, 1, 12, 14, 13},
			{14, 6, 5, 2, 11},
			{5, 6, 8, 13, 12},
		}
		m.answer = [][]float64{
			{1, 0, 0, 0, 0},
			{0, 1, 0, 0, 0},
			{0, 0, 1, 0, 0},
			{0, 0, 0, 1, 0},
			{0, 0, 0, 0, 1},
		}
		m.printMatrix("Заполнили матрицу")
		A := m.getValue()
		for j := 0; j < m.cols; j++ {
			_, maxRow := m.maxColValue(j)
			m.rowsSwap(j, maxRow)
			//m.printMatrix("матрица после переставления столбцов")
			m.toSingleMatrix(j, j)
			//m.printMatrix("вывод матрицы после прохода столбца")
		}
		m.printMatrix("Вывод итоговой матрицы")
		E := [][]float64{
			{1, 0, 0, 0, 0},
			{0, 1, 0, 0, 0},
			{0, 0, 1, 0, 0},
			{0, 0, 0, 1, 0},
			{0, 0, 0, 0, 1},
		}
		fmt.Println("Погрешность итоговой матрицы:", sumAbs(diff(E, multiply(A, m.getAnswer()))))
		m.spoil()
		m.printMatrix("Вывод испорченной матрицы")

		Xk := m.getAnswer()
		fmt.Println("Погрешность после поломки:", sumAbs(diff(E, multiply(A, Xk))))

		E2 := [][]float64{
			{2, 0, 0, 0, 0},
			{0, 2, 0, 0, 0},
			{0, 0, 2, 0, 0},
			{0, 0, 0, 2, 0},
			{0, 0, 0, 0, 2},
		}

		for i := 1; i < 7; i++ {
			Xk = multiply(Xk, diff(E2, multiply(A, Xk)))
			fmt.Println("Погрешность "+strconv.Itoa(i), sumAbs(diff(multiply(A, Xk), E))) //погрешность
		}
		fmt.Println("X_6: ")
		fmt.Println(Xk[0])
		fmt.Println(Xk[1])
		fmt.Println(Xk[2])
		fmt.Println(Xk[3])
		fmt.Println(Xk[4])

	} else {

		m := NewMatrix(5, 5)
		m.valueVyrozhd()
		fmt.Println("Вырожденная матрица: ")
		m.printMatrix("Заполнили матрицу")
		for j := 0; j < m.cols; j++ {
			_, maxRow := m.maxColValue(j)
			m.rowsSwap(j, maxRow)
			//m.printMatrix("матрица после переставления столбцов")
			m.toSingleMatrix(j, j)
			m.printMatrix("вывод матрицы после прохода столбца")
		}
		m.printMatrix("Вывод итоговой матрицы")
	}
	time.Sleep(time.Second * 10000)
}

func multiply(a, b [][]float64) [][]float64 {
	f := make([][]float64, 0, 0)
	for i := range a {
		k := make([]float64, len(b), len(b))
		for j := range b {
			for l := range b {
				k[j] += a[i][l] * b[l][j]
			}
		}
		f = append(f, k)

	}
	return f
}
func multiplyVector(a [][]float64, b []float64) []float64 {
	f := make([]float64, len(b), len(b))
	for i := range a {
		for j := range b {
			f[i] += a[i][j] * b[j]
		}

	}
	return f
}
func multiplyVectorOnNumber(a []float64, b float64) []float64 {
	f := make([]float64, len(a), len(a))
	for i := range a {
		f[i] = a[i] * b
	}
	return f
}

func diff(a, b [][]float64) [][]float64 {
	f := make([][]float64, 0, 0)
	for i := range a {
		f1 := make([]float64, 0, 0)
		for j := range a[i] {
			f1 = append(f1, a[i][j])
		}
		f = append(f, f1)
	}

	for i := range a {
		for j := range b {
			f[i][j] = a[i][j] - b[i][j]
		}

	}
	return f
}
func diffVector(a, b []float64) []float64 {
	f := make([]float64, 0, 0)
	for i := range a {
		f = append(f, a[i]-b[i])
	}

	return f
}

func sumAbs(a [][]float64) float64 {
	var b float64
	for _, v := range a {
		for _, v1 := range v {
			b += math.Abs(v1)
		}
	}
	return b
}

func gradient(A [][]float64) {
	EPS := math.Pow(10, -6)

	f := make([]float64, len(A), len(A))
	for i, v := range A {
		sum := 0.0
		for _, v1 := range v {
			sum += v1
		}
		f[i] = sum
	}

	fmt.Println("A:", A)
	fmt.Println("f:", f)

	x := make([]float64, len(A))
	omega := diffVector(multiplyVector(A, x), f)
	fmt.Println("x:", x)
	fmt.Println("omega:", omega)

	xArr := make([]float64, len(A))
	copy(xArr, x)

	iter := 0

	for {
		y := multiplyVector(A, omega)
		r := scolMultipl(omega, omega)
		s := scolMultipl(y, omega)

		if s < EPS*EPS {
			break
		}

		t := r / s
		x = diffVector(x, multiplyVectorOnNumber(omega, t))
		copy(xArr, x)
		iter += 1
		omega = diffVector(omega, multiplyVectorOnNumber(y, t))
	}
	fmt.Println("Решение найдено на ", iter, "итерации")
}

func scolMultipl(a, b []float64) float64 {
	f := 0.0
	for i := range a {
		f += a[i] * b[i]
	}
	return f
}
