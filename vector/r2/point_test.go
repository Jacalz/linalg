// Package r2 provides vectors and operations optimized for two dimensions.
package r2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVec2BetweenPoints(t *testing.T) {
	p1 := Point2{10, 5}
	p2 := Point2{-7, 23}

	assert.Equal(t, Vec2{p2.X - p1.X, p2.Y - p1.Y}, Vec2BetweenPoints(p1, p2))
	assert.Equal(t, Vec2{p1.X - p2.X, p1.Y - p2.Y}, Vec2BetweenPoints(p2, p1))
}
