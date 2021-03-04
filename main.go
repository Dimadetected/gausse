package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
)

type MatrixFunctions interface {
	maxColValue(col int) (float64, row int)
	valueRand()
	printMatrix()
	colsSwap(colFirst, colLast int)
}

type Matrix struct {
	value      [][]float64
	answer     []float64
	cols, rows int

	MatrixFunctions
}

func (m *Matrix) valueRand() {
	for i := 0; i < m.cols; i++ {
		newArray := make([]float64, m.rows, m.rows)
		for j := 0; j < m.rows; j++ {
			newArray[j] = float64(rand.Intn(15))
		}
		m.value = append(m.value, newArray)
		m.answer = append(m.answer, float64(rand.Intn(15)))
	}
}
func (m *Matrix) printMatrix() {
	for i := 0; i < m.cols; i++ {
		str := ""
		for j := 0; j < m.rows; j++ {
			str += strconv.FormatFloat(m.value[i][j], 'f', 2, 64) + "  "
		}
		str += "|  " + strconv.FormatFloat(m.answer[i], 'f', 2, 64)
		fmt.Println(str)
	}
}
func (m *Matrix) maxColValue(col int) (float64, int) {
	maxIndex := 0
	max := m.value[maxIndex][col]
	for j := 0; j < m.rows; j++ {
		if max < math.Abs(m.value[j][col]) {
			max = math.Abs(m.value[j][col])
			maxIndex = j
		}
	}
	return max, maxIndex
}
func (m *Matrix) colsSwap(colFirst, colLast int) {
	//Меняем местами строки в матрице
	newArr := make([]float64, len(m.value[colLast]), len(m.value[colLast]))
	for i := 0; i < len(m.value[colLast]); i++ {
		newArr[i] = m.value[colLast][i]
	}
	m.value[colLast] = m.value[colFirst]
	m.value[colFirst] = newArr

	//Меняем местами строки свободных членов
	newValue := m.answer[colFirst]
	m.answer[colFirst] = m.answer[colLast]
	m.answer[colLast] = newValue
}

func NewMatrix(cols, rows int) *Matrix {
	return &Matrix{
		cols:   cols,
		rows:   rows,
		value:  make([][]float64, 0, 0),
		answer: make([]float64, 0, 0),
	}
}

func main() {
	m := NewMatrix(5, 5)
	m.valueRand()

	m.printMatrix()
	fmt.Println("-------")
	//fmt.Println(m.maxColValue(0))
	m.colsSwap(1, 2)
	m.printMatrix()
	//fmt.Println(m.value)
	//fmt.Println(m.answer)
}
