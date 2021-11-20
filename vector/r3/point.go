// Package r2 provides vectors and operations optimized for three dimensions.
package r3

// Point3 is a point within a three dimensional room.
type Point3 struct {
	X, Y, Z float64
}

// Vec3BetweenPoints returns a vector from point a to point b.
func Vec3BetweenPoints(p, q Point3) Vec3 {
	return Vec3{q.X - p.X, q.Y - p.Y, q.Z - p.Z}
}
