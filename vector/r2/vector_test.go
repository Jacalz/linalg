package r2

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	u := Vec2{5, 6}
	v := Vec2{15, -2}

	assert.Equal(t, Vec2{20, 4}, Add(u, v))
	assert.Equal(t, Add(u, v), Add(v, u))

	assert.Equal(t, ScalarMult(u, 2), Add(u, u))
	assert.Equal(t, ScalarMult(v, 2), Add(v, v))
}

func TestSub(t *testing.T) {
	u := Vec2{5, 6}
	v := Vec2{15, -2}

	assert.Equal(t, Vec2{-10, 8}, Sub(u, v))
	assert.Equal(t, Vec2{10, -8}, Sub(v, u))
}

func TestScalarMult(t *testing.T) {
	u := Vec2{0, 0}
	v := Vec2{15, -2}
	w := Vec2{-10, 6}

	assert.Equal(t, u, ScalarMult(u, 29.5))
	assert.Equal(t, Vec2{30, -4}, ScalarMult(v, 2))
	assert.Equal(t, Vec2{-15, 9}, ScalarMult(w, 1.5))
}

func TestLength(t *testing.T) {
	u := Vec2{3, 4}
	v := Vec2{15, -2}

	assert.Equal(t, 5.0, Length(u))
	assert.Equal(t, math.Sqrt(v.X*v.X+v.Y*v.Y), Length(v))

	assert.Equal(t, 1.0, Length(UnitVector(u)))
	assert.Equal(t, 1.0, Length(UnitVector(v)))

	v = Vec2{0, 5}
	assert.Equal(t, Length(v), Length(u))
}

func TestLengthSquared(t *testing.T) {
	u := Vec2{3, 4}
	v := Vec2{15, -2}

	assert.Equal(t, u.X*u.X+u.Y*u.Y, LengthSquared(u))
	assert.Equal(t, v.X*v.X+v.Y*v.Y, LengthSquared(v))

	assert.Equal(t, math.Round(Length(u)*Length(u)), LengthSquared(u))
	assert.Equal(t, math.Round(Length(v)*Length(v)), LengthSquared(v))

	assert.Equal(t, 1.0, math.Round(LengthSquared(UnitVector(u))))
	assert.Equal(t, 1.0, math.Round(LengthSquared(UnitVector(v))))

	v = Vec2{0, 5}
	assert.Equal(t, LengthSquared(v), LengthSquared(u))
}

func TestParalell(t *testing.T) {
	u := Vec2{3, 4}
	v := Vec2{15, -2}

	assert.Equal(t, Parallel(u, v), Parallel(v, u))
	assert.True(t, Parallel(u, ScalarMult(u, 2.5)))
	assert.True(t, Parallel(u, UnitVector(u)))
	assert.False(t, Parallel(u, v))
}
