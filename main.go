package main

import "fmt"

func main() {
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

	m = NewMatrix(5, 5)
	m.valueVyrozhd()
	fmt.Println("Вырожденная матрица: ")
	m.printMatrix("Заполнили матрицу")
	for j := 0; j < m.cols; j++ {
		_, maxRow := m.maxColValue(j)
		m.rowsSwap(j, maxRow)
		//m.printMatrix("матрица после переставления столбцов")
		m.toSingleMatrix(j, j)
		//m.printMatrix("вывод матрицы после прохода столбца")
	}
	m.printMatrix("Вывод итоговой матрицы")
}
