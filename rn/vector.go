package rn

import (
	"errors"
	"math"
)

// VecN is a vector in the dimension n.
type VecN struct {
	Values []float64 // The vector cordinates.
}

// Dim returns what dimension the vector lies within.
func (v *VecN) Dim() int {
	return len(v.Values)
}

// NewVecN returns a vector with the coordinates in u and the dimension of the length of u.
func NewVecN(u []float64) *VecN {
	return &VecN{Values: u}
}

// Add returns a new vector that is the result of u and v added together.
// Both vectors must be in the same dimension.
func Add(u, v *VecN) (*VecN, error) {
	dimension := u.Dim()
	if dimension != v.Dim() {
		return nil, errors.New("The vectors do not have the same dimension")
	}

	new := make([]float64, dimension)
	for i, val := range u.Values {
		new[i] = val + v.Values[i]
	}

	return &VecN{Values: new}, nil
}

// Sub returns a new vector that is the result of v subracted from u.
// Both vectors must be in the same dimension.
func Sub(u, v *VecN) (*VecN, error) {
	dimension := u.Dim()
	if dimension != v.Dim() {
		return nil, errors.New("The vectors do not have the same dimension")
	}

	new := make([]float64, dimension)
	for i, val := range u.Values {
		new[i] = val - v.Values[i]
	}

	return &VecN{Values: new}, nil
}

// Mult returns a new vector that is the result of u multiplied with v.
// Both vectors must be in the same dimension.
func Mult(u, v *VecN) (*VecN, error) {
	dimension := u.Dim()
	if dimension != v.Dim() {
		return nil, errors.New("The vectors do not have the same dimension")
	}

	new := make([]float64, dimension)
	for i, val := range u.Values {
		new[i] = val * v.Values[i]
	}

	return &VecN{Values: new}, nil
}

// ScalarMult returns a new vector that is the result of u multiplied with the scalar a.
func ScalarMult(u *VecN, a float64) *VecN {
	new := make([]float64, u.Dim())
	for i, val := range u.Values {
		new[i] = val * a
	}

	return &VecN{Values: new}
}

// Abs returns the absolute value (length) of the vector u.
func Abs(u *VecN) float64 {
	under := float64(0)
	for _, val := range u.Values {
		under += val * val
	}

	return math.Sqrt(under)
}

// ScalarProduct returns the scalar product of the vectors u and v.
// Both vectors must be in the same dimension.
func ScalarProduct(u, v *VecN) (float64, error) {
	dimension := u.Dim()
	if dimension != v.Dim() {
		return 0, errors.New("The vectors do not have the same dimension")
	}

	under := float64(0)
	for i, val := range u.Values {
		under += val * v.Values[i]
	}

	return under, nil
}

// UnitVector returns a unit vector (length 1) from u.
func UnitVector(u *VecN) *VecN {
	return ScalarMult(u, 1/Abs(u))
}

// OrthoProject projects the vector u orthogonallY on the vector v.
// Both vectors must be in the same dimension.
func OrthoProject(u, v *VecN) (*VecN, error) {
	e := UnitVector(v)
	scalar, err := ScalarProduct(u, e)
	if err != nil {
		return nil, err
	}

	return ScalarMult(e, scalar), nil
}

// Orthogonal returns true if the vectors u and v are orthogonal.
// Both vectors must be in the same dimension.
func Orthogonal(u, v *VecN) (bool, error) {
	scalar, err := ScalarProduct(u, v)
	if err != nil {
		return false, err
	}

	return scalar == 0, nil
}

// AngleBetween returns the angle, in radians, between the vetors u and v.
// Both vectors must be in the same dimension.
func AngleBetween(u, v *VecN) (float64, error) {
	scalar, err := ScalarProduct(u, v)
	if err != nil {
		return 0, err
	}

	return math.Acos(scalar / (Abs(u) * Abs(v))), nil
}
