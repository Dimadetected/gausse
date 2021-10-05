package main

import (
	"fmt"
	"math"
	"math/rand"
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
	answer     [][]float64
	cols, rows int

	MatrixFunctions
}

func (m *Matrix) getValue() [][]float64 {
	v := make([][]float64, 0, 0)
	for i := range m.value {
		v1 := make([]float64, 0, 0)
		for _, v2 := range m.value[i] {
			v1 = append(v1, v2)
		}
		v = append(v, v1)
	}
	return v
}
func (m *Matrix) getAnswer() [][]float64 {
	v := make([][]float64, 0, 0)
	for i := range m.answer {
		v1 := make([]float64, 0, 0)
		for _, v2 := range m.answer[i] {
			v1 = append(v1, v2)
		}
		v = append(v, v1)
	}
	return v
}
func (m *Matrix) valueRand() {
	for i := 0; i < m.cols; i++ {
		newArray := make([]float64, m.rows, m.rows)
		for j := 0; j < m.rows; j++ {
			newArray[j] = float64(rand.Intn(15))
		}
		m.value = append(m.value, newArray)

		m.answer = [][]float64{
			{1, 0, 0, 0, 0},
			{0, 1, 0, 0, 0},
			{0, 0, 1, 0, 0},
			{0, 0, 0, 1, 0},
			{0, 0, 0, 0, 1},
		}
	}
}
func (m *Matrix) valueVyrozhd() {
	m.answer = [][]float64{
		{1, 0, 3, 0, 2},
		{1, 2, 3, 4, 0},
		{1, 0, 5, 0, 0},
		{7, 5, 0, 9, 0},
		{4, 0, 0, 0, 0},
	}
	m.value = [][]float64{
		//{1, 0, 0, 0, 0},
		//{0, 1, 0, 0, 0},
		//{0, 0, 1, 0, 0},
		//{0, 0, 0, 1, 0},
		//{0, 0, 0, 0, 1},
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20},
		{27, 29, 31, 33, 35},
	}
	//todo: добавить генерацию массива
}
func (m *Matrix) printMatrix(text string) {

	fmt.Println(text)
	fmt.Println("-------")
	for i := 0; i < m.cols; i++ {
		str := ""
		for j := 0; j < m.rows; j++ {
			str += fmt.Sprintf("%5.2f  ", m.value[i][j])
		}
		str += "|  " + fmt.Sprintf("%5.2f  ", m.answer[i])
		fmt.Println(str)
	}
}
func (m *Matrix) maxColValue(col int) (float64, int) {
	maxIndex := col
	max := m.value[maxIndex][col]
	for j := col; j < m.rows; j++ {
		if max < math.Abs(m.value[j][col]) {
			max = math.Abs(m.value[j][col])
			maxIndex = j
		}
	}
	if math.Abs(max) < 0.0000000000001 {
		panic("Матрица вырожденна...")
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

	for i := 0; i < m.cols; i++ {
		m.value[row][i] = m.value[row][i] / deleter
	}
	for i := range m.answer[row] {
		m.answer[row][i] = m.answer[row][i] / deleter
	}
	//m.printMatrix("Матрица с единицей на главной диагонали: ")

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
	for k := range m.answer[secondRow] {
		m.answer[secondRow][k] = m.answer[secondRow][k] - (m.answer[firstRow][k] * kef)
	}
}

func NewMatrix(cols, rows int) *Matrix {
	return &Matrix{
		cols:   cols,
		rows:   rows,
		value:  make([][]float64, 0, 0),
		answer: make([][]float64, 0, 0),
	}
}

func (m *Matrix) spoil() {
	for k, v := range m.answer {
		for k1 := range v {
			m.answer[k][k1] = m.answer[k][k1] - rand.Float64()*(0.01-0.001)
		}
	}
}
