package r3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVec2BetweenPoints(t *testing.T) {
	p1 := Point3{10, 5, -9}
	p2 := Point3{-7, 23, 2}

	assert.Equal(t, Vec3{p2.X - p1.X, p2.Y - p1.Y, p2.Z - p1.Z}, Vec3BetweenPoints(p1, p2))
	assert.Equal(t, Vec3{p1.X - p2.X, p1.Y - p2.Y, p1.Z - p2.Z}, Vec3BetweenPoints(p2, p1))
}
