package r3

import "math"

// Vec3Zero is a Zero vector in the room.
// This vector is paralell to all other vectors.
var Vec3Zero = Vec3{0, 0, 0}

// Vec3 is a three dimensional vector in the room.
type Vec3 struct {
	X, Y, Z float64
}

// Add adds the two vectors u and v together.
func Add(u, v Vec3) Vec3 {
	return Vec3{u.X + v.X, u.Y + v.Y, u.Z + v.Z}
}

// Sub subtracts the vector v from u.
func Sub(u, v Vec3) Vec3 {
	return Vec3{u.X - v.X, u.Y - v.Y, u.Z - v.Z}
}

//Mult multiplies the vector u with the vector v.
func Mult(u, v Vec3) Vec3 {
	return Vec3{u.X * v.X, u.Y * v.Y, u.Z * v.Z}
}

// ScalarMult multiplies the vector u with the scalar s.
func ScalarMult(u Vec3, s float64) Vec3 {
	return Vec3{u.X * s, u.Y * s, u.Z * s}
}

// Abs returns the absolute value (length) of the vector u.
func Abs(u Vec3) float64 {
	return math.Sqrt(u.X*u.X + u.Y*u.Y + u.Z*u.Z)
}

// ScalarProduct returns the scalar product of the vectors u and v.
func ScalarProduct(u, v Vec3) float64 {
	return u.X*v.X + u.Y*v.Y + u.Z*v.Z
}

// CrossProduct returns the cross product of the vectors u and v.
func CrossProduct(u, v Vec3) Vec3 {
	return Vec3{
		X: u.Y*v.Z - u.Z*v.Y,
		Y: u.Z*v.X - v.Z*u.X,
		Z: u.X*v.Y - u.Y*v.X,
	}
}

// UnitVector returns a unit vector (length 1) from u.
func UnitVector(u Vec3) Vec3 {
	return ScalarMult(u, 1/Abs(u))
}

// OrthoProject projects the vector u orthogonallY on the vector v.
func OrthoProject(u, v Vec3) Vec3 {
	e := UnitVector(v)
	return ScalarMult(e, ScalarProduct(u, e))
}

// Parallel returns true if the vectors u and v are parallell.
func Parallel(u, v Vec3) bool {
	return CrossProduct(u, v) == Vec3Zero
}

// Orthogonal returns true if the vectors u and v are orthogonal.
func Orthogonal(u, v Vec3) bool {
	return ScalarProduct(u, v) == 0
}

// InSamePlane returns true if the vectors u, v and w are in the same plane.
func InSamePlane(u, v, w Vec3) bool {
	return ScalarProduct(CrossProduct(u, v), w) == 0
}

// AngleBetween returns the angle, in radians, between the vetors u and v.
func AngleBetween(u, v Vec3) float64 {
	return math.Acos(ScalarProduct(u, v) / (Abs(u) * Abs(v)))
}
