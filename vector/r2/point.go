// Package r2 provides vectors and operations optimized for two dimensions.
package r2

// Point2 is a point within two dimensions.
type Point2 struct {
	X, Y float64
}

// Vec2BetweenPoints returns a vector from point p1 to point p2.
func Vec2BetweenPoints(p1, p2 Point2) Vec2 {
	return Vec2{p2.X - p1.X, p2.Y - p1.Y}
}
