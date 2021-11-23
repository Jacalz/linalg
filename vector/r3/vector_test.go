package r3

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	u := Vec3{5, 6, 2}
	v := Vec3{15, -2, 12}

	assert.Equal(t, Vec3{20, 4, 14}, Add(u, v))
	assert.Equal(t, Add(u, v), Add(v, u))

	assert.Equal(t, ScalarMult(u, 2), Add(u, u))
	assert.Equal(t, ScalarMult(v, 2), Add(v, v))
}

func TestSub(t *testing.T) {
	u := Vec3{5, 6, 2}
	v := Vec3{15, -2, 12}

	assert.Equal(t, Vec3{-10, 8, -10}, Sub(u, v))
	assert.Equal(t, Vec3{10, -8, 10}, Sub(v, u))
}

func TestScalarMult(t *testing.T) {
	u := Vec3{0, 0, 0}
	v := Vec3{15, -2, 1.5}
	w := Vec3{-10, 6, 5}

	assert.Equal(t, u, ScalarMult(u, 29.5))
	assert.Equal(t, Vec3{30, -4, 3}, ScalarMult(v, 2))
	assert.Equal(t, Vec3{-15, 9, 7.5}, ScalarMult(w, 1.5))
}

func TestLength(t *testing.T) {
	u := Vec3{3, 4, -2}
	v := Vec3{15, -2, 0}

	assert.Equal(t, math.Sqrt(u.X*u.X+u.Y*u.Y+u.Z*u.Z), Length(u))
	assert.Equal(t, math.Sqrt(v.X*v.X+v.Y*v.Y+v.Z*v.Z), Length(v))

	assert.Equal(t, 1.0, Length(UnitVector(u)))
	assert.Equal(t, 1.0, Length(UnitVector(v)))

	u = Vec3{0, 5, 0}
	v = Vec3{0, 0, 5}
	assert.Equal(t, Length(v), Length(u))
}

func TestLengthSquared(t *testing.T) {
	u := Vec3{3, 4, -2}
	v := Vec3{15, -2, 0}

	assert.Equal(t, u.X*u.X+u.Y*u.Y+u.Z*u.Z, LengthSquared(u))
	assert.Equal(t, v.X*v.X+v.Y*v.Y+v.Z*v.Z, LengthSquared(v))

	assert.Equal(t, math.Round(Length(u)*Length(u)), LengthSquared(u))
	assert.Equal(t, math.Round(Length(v)*Length(v)), LengthSquared(v))

	assert.Equal(t, 1.0, LengthSquared(UnitVector(u)))
	assert.Equal(t, 1.0, LengthSquared(UnitVector(v)))

	u = Vec3{0, 5, 0}
	v = Vec3{0, 0, 5}
	assert.Equal(t, LengthSquared(v), LengthSquared(u))
}
