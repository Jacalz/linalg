// Package matrix provides a generic matrix abstraction.
package matrix

import (
	"errors"
	"math"
)

var (
	errorInvalidMultiplication = errors.New("the columns of u must be equal to the rows of v")
	errorDifferentSize         = errors.New("the matrix sizes are not equal")
)

// Matrix is an extension of a 2d slice if float64.
type Matrix [][]float64

// Rows returns the amount of rows in the matrix.
func (m Matrix) Rows() int {
	return len(m)
}

// Cols returns the amount of columns in the matrix.
func (m Matrix) Cols() int {
	return len(m[0])
}

// New allocates a new matrix with the set ammount of rows and columns.
func New(rows, cols int) Matrix {
	data := make([][]float64, rows)
	for r := range data {
		data[r] = make([]float64, cols)
	}
	return data
}

// Mult multiplies the matrices together to form a new matrix.
// The new matrix will have the rows as u and columns as b.
func Mult(u, v Matrix) (Matrix, error) {
	if u.Cols() != v.Rows() {
		return nil, errorInvalidMultiplication
	}

	rows, cols := u.Rows(), v.Cols()
	data := New(rows, cols)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			for k := 0; k < rows; k++ {
				data[i][j] = math.FMA(u[i][k], v[k][j], data[i][j])
			}
		}
	}

	return data, nil
}

// Add adds the matrices together.
// The matrices must have the same ammount of rows and columns.
func Add(u, v Matrix) (Matrix, error) {
	rows, cols := u.Rows(), u.Cols()
	if rows != v.Rows() || cols != v.Cols() {
		return nil, errorDifferentSize
	}

	data := New(rows, cols)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			data[i][j] = u[i][j] + v[i][j]
		}
	}

	return data, nil
}

// Sub subtracts the matrix v from u.
// The matrices must have the same ammount of rows and columns.
func Sub(u, v Matrix) (Matrix, error) {
	rows, cols := u.Rows(), u.Cols()
	if rows != v.Rows() || cols != v.Cols() {
		return nil, errorDifferentSize
	}

	data := New(rows, cols)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			data[i][j] = u[i][j] - v[i][j]
		}
	}

	return data, nil
}
