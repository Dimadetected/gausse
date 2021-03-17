package main

import (
	"fmt"
	"time"
)

func main() {
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
