// Package constants provides predefined matricies for specific use cases.
package constants

import (
	"math"

	"github.com/Jacalz/linalg/matrix"
)

// RotateX3D returns a three dimensional rotation matrix
// that can be used rotate rad radians around the X-axis.
func RotateX3D(rad float64) matrix.Matrix {
	return matrix.Matrix{
		{1, 0, 0},
		{0, math.Cos(rad), -math.Sin(rad)},
		{0, math.Sin(rad), math.Cos(rad)},
	}
}

// RotateY3D returns a three dimensional rotation matrix
// that can be used rotate rad radians around the Y-axis.
func RotateY3D(rad float64) matrix.Matrix {
	return matrix.Matrix{
		{math.Cos(rad), 0, math.Sin(rad)},
		{0, 1, 0},
		{-math.Sin(rad), 0, math.Cos(rad)},
	}
}

// RotateZ3D returns a three dimensional rotation matrix
// that can be used rotate rad radians around the Z-axis.
func RotateZ3D(rad float64) matrix.Matrix {
	return matrix.Matrix{
		{math.Cos(rad), -math.Sin(rad), 0},
		{math.Sin(rad), math.Cos(rad), 0},
		{0, 0, 1},
	}
}

// Rotate2D returns a two dimensional rotation matrix
// that can be used rotate rad radians around the Z-axis.
func Rotate2D(rad float64) matrix.Matrix {
	return matrix.Matrix{
		{math.Cos(rad), -math.Sin(rad)},
		{math.Sin(rad), math.Cos(rad)},
	}
}
