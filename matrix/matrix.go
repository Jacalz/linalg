// Package matrix provides matricies and operations for any dimension.
package matrix

import (
	"errors"
	"math"

	"github.com/Jacalz/linalg/vector"
)

var (
	errorInvalidMultiplication = errors.New("the columns of u must be equal to the rows of v")
	errorDifferentSize         = errors.New("the matrix sizes are not equal")
	errorNotQuadratic          = errors.New("the matrix must be quadratic")
)

// Matrix is an extension of a 2d slice if float64.
type Matrix [][]float64

// Rows returns the amount of rows in the matrix.
func (m Matrix) Rows() int {
	return len(m)
}

// Cols returns the amount of columns in the matrix.
func (m Matrix) Cols() int {
	if m.Rows() == 0 {
		return 0
	}

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

// NewFromVec creates a new matrix from a set of vectors.
// The vectors are assummed as column vectors with the same length.
func NewFromVec(vectors ...vector.VecN) Matrix {
	if len(vectors) < 1 {
		return nil
	}

	data := New(len(vectors[0]), len(vectors))
	for r := range data {
		for c := range data[r] {
			data[r][c] = vectors[c][r]
		}
	}

	return data
}

// Mult multiplies the matrices together to form a new matrix.
// The new matrix will have the same rows as u and columns as v.
func Mult(u, v Matrix) (Matrix, error) {
	if u.Cols() != v.Rows() {
		return nil, errorInvalidMultiplication
	}

	// TODO: Use Strassen Algorithm, https://en.wikipedia.org/wiki/Strassen_algorithm

	rows, cols := u.Rows(), v.Cols()
	data := New(rows, cols)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			for k := 0; k < v.Rows(); k++ {
				data[i][j] = math.FMA(u[i][k], v[k][j], data[i][j])
			}
		}
	}

	return data, nil
}

// ScalarMult multiplies the vector u with the scalar s.
func ScalarMult(u Matrix, s float64) Matrix {
	rows, cols := u.Rows(), u.Cols()
	data := New(rows, cols)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			data[i][j] = u[i][j] * s
		}
	}

	return data
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

// AddVec adds the vector to the matrix.
// The vector must be of the same length as the matrix rows.
func AddVec(u Matrix, v vector.VecN) (Matrix, error) {
	rows, cols := u.Rows(), u.Cols()
	if rows != v.Dim() {
		return nil, errorDifferentSize
	}

	data := New(rows, cols)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			data[i][j] = u[i][j] + v[i]
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

// SubVec subtracts the vector from the matrix.
// The vector must be of the same length as the matrix rows.
func SubVec(u Matrix, v vector.VecN) (Matrix, error) {
	rows, cols := u.Rows(), u.Cols()
	if rows != v.Dim() {
		return nil, errorDifferentSize
	}

	data := New(rows, cols)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			data[i][j] = u[i][j] - v[i]
		}
	}

	return data, nil
}

// Transpose returns the transposed matrix of u.
func Transpose(u Matrix) Matrix {
	rows, cols := u.Rows(), u.Cols()
	data := New(cols, rows)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			data[j][i] = u[i][j]
		}
	}

	return data
}

// Orthogonal returns true if the matrix is orthogonal.
// An error will be returned if the matrix is not quadratic.
func Orthogonal(u Matrix) (bool, error) {
	if u.Rows() != u.Cols() {
		return false, errorNotQuadratic
	}

	rows := u.Rows()
	rowsum := float64(0)
	for i := 0; i < rows; i++ { // ON-matrix => A * A^t == I
		for j := 0; j < rows; j++ {
			for k := 0; k < rows; k++ {
				rowsum = math.FMA(u[i][k], u[j][k], rowsum)
			}

			if (i != j && rowsum != 0) || (i == j && rowsum != 1) {
				return false, nil
			}

			rowsum = 0
		}
	}

	return true, nil
}

// ON returns true if the matrix is an ON-matrix.
// It returns an error if the matrix is not quadratic.
func ON(u Matrix) (bool, error) {
	if u.Rows() != u.Cols() {
		return false, errorNotQuadratic
	}

	rows := u.Rows()
	rowsum := float64(0)
	for i := 0; i < rows; i++ {
		if length := vector.Abs(u[i]); length != 1 {
			return false, nil
		}

		// Manual inlining of part of the Orthogonal function.
		for j := 0; j < rows; j++ {
			for k := 0; k < rows; k++ {
				rowsum = math.FMA(u[i][k], u[j][k], rowsum)
			}

			if (i != j && rowsum != 0) || (i == j && rowsum != 1) {
				return false, nil
			}

			rowsum = 0
		}
	}

	return true, nil
}
