// Package r2 provides easy-to-use abstractions optimized for two dimensions.
package r2

// Point2 is a point within a two dimensional plane.
type Point2 struct {
	X, Y float64
}

// Vec3BetweenPoints returns a vector from point a to point b.
func Vec3BetweenPoints(p, q *Point2) *Vec2 {
	return &Vec2{q.X - p.X, q.Y - p.Y}
}
