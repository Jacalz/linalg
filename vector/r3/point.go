// Package r2 provides vectors and operations optimized for three dimensions.
package r3

// Point3 is a point within three dimensions.
type Point3 struct {
	X, Y, Z float64
}

// Vec3BetweenPoints returns a vector from point p1 to point p2.
func Vec3BetweenPoints(p1, p2 Point3) Vec3 {
	return Vec3{p2.X - p1.X, p2.Y - p1.Y, p2.Z - p1.Z}
}
