package r3

import (
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
