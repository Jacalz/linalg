package r2

import "math"

// Vec2Zero is a zero vector in the plane.
// This vector is paralell to all other vectors.
var Vec2Zero = Vec2{0, 0}

// Vec2 is a two dimensional vector in the plane.
type Vec2 struct {
	X, Y float64
}

// Add adds the two vectors u and v together.
func Add(u, v Vec2) Vec2 {
	return Vec2{u.X + v.X, u.Y + v.Y}
}

// Sub subtracts the vector v from u.
func Sub(u, v Vec2) Vec2 {
	return Vec2{u.X - v.X, u.Y - v.Y}
}

// ScalarMult multiplies the vector u with the scalar s.
func ScalarMult(u Vec2, s float64) Vec2 {
	return Vec2{u.X * s, u.Y * s}
}

// Abs returns the absolute value (length) of the vector u.
func Abs(u Vec2) float64 {
	return math.Sqrt(u.X*u.X + u.Y*u.Y)
}

// ScalarProduct returns the scalar product of the vectors u and v.
func ScalarProduct(u, v Vec2) float64 {
	return u.X*v.X + u.Y*v.Y
}

// UnitVector returns a unit vector (length 1) from u.
func UnitVector(u Vec2) Vec2 {
	return ScalarMult(u, 1/Abs(u))
}

// OrthoProject projects the vector u orthogonallY on the vector v.
func OrthoProject(u, v Vec2) Vec2 {
	e := UnitVector(v)
	return ScalarMult(e, ScalarProduct(u, e))
}

// Parallel returns true if the vectors u and v are parallell.
func Parallel(u, v Vec2) bool {
	return u.X/v.X == u.Y/v.Y
}

// Orthogonal returns true if the vectors u and v are orthogonal.
func Orthogonal(u, v Vec2) bool {
	return ScalarProduct(u, v) == 0
}

// AngleBetween returns the angle, in radians, between the vetors u and v.
func AngleBetween(u, v Vec2) float64 {
	return math.Acos(ScalarProduct(u, v) / (Abs(u) * Abs(v)))
}

// TODO: Implement calculations for distance between vectors.
// func DistanceBetween(u, v Vec2) float64 {
// 	return 0
// }
