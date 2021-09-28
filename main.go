package main

import (
	"fmt"
	"time"
)

func main() {
	var a string
	fmt.Println("Введите \n 1)Если нужно обычную матрицу \n 2)Если нужно вырожденную матрицу")
	//fmt.Scan(&a)
	a = "1"
	if a == "1" {

		m := NewMatrix(5, 5)
		m.valueRand()
		fmt.Println("Обычная матрица: ")
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
		m.spoil()
		m.printMatrix("Вывод испорченной матрицы")

		Xk := m.getAnswer()
		E := [][]float64{
			{1, 0, 0, 0, 0},
			{0, 1, 0, 0, 0},
			{0, 0, 1, 0, 0},
			{0, 0, 0, 1, 0},
			{0, 0, 0, 0, 1},
		}
		E2 := [][]float64{
			{2, 0, 0, 0, 0},
			{0, 2, 0, 0, 0},
			{0, 0, 2, 0, 0},
			{0, 0, 0, 2, 0},
			{0, 0, 0, 0, 2},
		}

		for i := 1; i < 7; i++ {
			Xk = multiply(Xk, diff(E2, multiply(A, Xk)))
			fmt.Println(i, diff(multiply(A, Xk), E))
			time.Sleep(1 * time.Second)
		}

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
