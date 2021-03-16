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
	toSingleMatrix(col, row int)
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
func (m *Matrix) valueVyrozhd() {
	m.answer = make([]float64, m.cols, m.rows)
	//todo: добавить генерацию массива
}
func (m *Matrix) printMatrix(text string) {

	fmt.Println(text)
	fmt.Println("-------")
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
func (m *Matrix) rowsSwap(rowFirst, rowLast int) {
	//Меняем местами строки в матрице
	newArr := make([]float64, len(m.value[rowLast]), len(m.value[rowLast]))
	for i := 0; i < len(m.value[rowLast]); i++ {
		newArr[i] = m.value[rowLast][i]
	}
	m.value[rowLast] = m.value[rowFirst]
	m.value[rowFirst] = newArr

	//Меняем местами строки свободных членов
	newValue := m.answer[rowFirst]
	m.answer[rowFirst] = m.answer[rowLast]
	m.answer[rowLast] = newValue
}

func (m *Matrix) toSingleMatrix(col, row int) {
	deleter := m.value[row][col]

	for q := 0; q < m.rows; q++ {
		if q > row && m.value[q][col] == 0 {
			panic("Матрица вырождена...")
		}
	}

	for i := 0; i < m.cols; i++ {
		m.value[row][i] = math.Round(m.value[row][i]/deleter*100) / 100
	}

	m.answer[row] = math.Round(m.answer[row]/deleter*100) / 100
	m.printMatrix("Матрица с единицей на главной диагонали: ")

	for i := 0; i < m.rows; i++ {
		if i != row {
			m.rowsDifferent(row, i, col)
		}
	}
}

func (m *Matrix) rowsDifferent(firstRow, secondRow, col int) {
	kef := m.value[secondRow][col]
	for i := 0; i < m.cols; i++ {
		m.value[secondRow][i] = m.value[secondRow][i] - (m.value[firstRow][i] * kef)
	}
	m.answer[secondRow] = m.answer[secondRow] - (m.answer[firstRow] * kef)
}

func NewMatrix(cols, rows int) *Matrix {
	return &Matrix{
		cols:   cols,
		rows:   rows,
		value:  make([][]float64, 0, 0),
		answer: make([]float64, 0, 0),
	}
}
