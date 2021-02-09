// Package rn provides a general, but less optimized, library for up to n dimensions.
// Use packages r2 and r3 for calculations within the plane and room respectivly.
package rn

import "errors"

// PointN is a point within a three dimensional room.
type PointN struct {
	Coord []float64 // The point coordinates.
}

// Dim returns what dimension the point lies within.
func (p *PointN) Dim() int {
	return len(p.Coord)
}

// VecNBetweenPoints returns a vector from point a to point b.
// The points must be in the same dimension.
func VecNBetweenPoints(p, q *PointN) (*VecN, error) {
	dimension := p.Dim()
	if dimension != q.Dim() {
		return nil, errors.New("The points do not have the same dimension")
	}

	new := make([]float64, dimension)
	for i, val := range q.Coord {
		new[i] = val - p.Coord[i]
	}

	return &VecN{Values: new}, nil
}
