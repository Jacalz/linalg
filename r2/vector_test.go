package r2

import (
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
