// Package matrix provides a generic matrix abstraction.
package matrix

import "errors"

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

// Mult multiplies the matrixes toegether to form anew matrix.
// The new matrix will have the rows as u and columns as b.
func Mult(u, v Matrix) (Matrix, error) {
	if u.Cols() != v.Rows() {
		return nil, errors.New("The columns of u must be equal to the rows of v")
	}

	data := make([][]float64, v.Rows())
	for r := range data {
		data[r] = make([]float64, v.Cols())
	}

	for i := 0; i < u.Rows(); i++ {
		for j := 0; j < v.Cols(); j++ {
			for k := 0; k < v.Rows(); k++ {
				data[i][j] += u[i][k] * v[k][j]
			}
		}
	}

	return data, nil
}
