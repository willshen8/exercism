package matrix

import (
	"errors"
	"strconv"
	"strings"
)

// Matrix is a 2 dimensional int array
type Matrix [][]int

// New is a constructor that sets a matrix based on data from input string
func New(input string) (Matrix, error) {
	parsedRows := strings.Split(input, "\n")
	var newMatrix Matrix
	for _, row := range parsedRows {
		trimmedInts := strings.Split(strings.Trim(row, " "), " ")
		ints := make([]int, len(trimmedInts))
		for i, v := range trimmedInts {
			parsedInt, err := strconv.Atoi(string(v))
			if err != nil {
				return nil, err
			}
			ints[i] = parsedInt
		}
		newMatrix = append(newMatrix, ints)
	}

	rowLen := len(newMatrix[0])
	for _, v := range newMatrix {
		if len(v) != rowLen {
			return nil, errors.New("Uneven rows")
		}
	}
	return newMatrix, nil
}

// Rows returns the rows of the matrix m
func (m Matrix) Rows() [][]int {
	var rows [][]int
	for _, v := range m {
		var copiedRow = make([]int, len(v))
		copy(copiedRow, v)
		rows = append(rows, copiedRow)
	}
	return rows
}

// Cols returns the columns of the matrix m
func (m Matrix) Cols() [][]int {
	var cols [][]int
	for i := 0; i < len(m[0]); i++ {
		var columns []int
		for _, v := range m {
			columns = append(columns, v[i])
		}
		cols = append(cols, columns)
	}
	return cols
}

// Set allocate a matrix element with value located at row, col
func (m *Matrix) Set(row, col, val int) bool {
	if len(*m) == 0 || row < 0 || col < 0 || row > len((*m)[0])-1 || col > len((*m))-1 {
		return false
	}
	(*m)[row][col] = val
	return true
}
