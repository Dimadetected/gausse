package main

import (
	"fmt"
	"time"
)

func main() {
	startMethod3()
	return
	var a string
	fmt.Println("Введите \n 1)Если нужно обычную матрицу \n 2)Если нужно вырожденную матрицу")
	fmt.Scan(&a)
	if a == "1" {

		m := NewMatrix(5, 5)
		m.valueRand()
		fmt.Println("Обычная матрица: ")
		m.printMatrix("Заполнили матрицу")
		for j := 0; j < m.cols; j++ {
			_, maxRow := m.maxColValue(j)
			m.rowsSwap(j, maxRow)
			//m.printMatrix("матрица после переставления столбцов")
			m.toSingleMatrix(j, j)
			//m.printMatrix("вывод матрицы после прохода столбца")
		}
		m.printMatrix("Вывод итоговой матрицы")

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
