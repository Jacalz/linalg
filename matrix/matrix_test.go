package matrix

import (
	"testing"

	"github.com/Jacalz/linalg/vector"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	expected := Matrix{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	assert.Equal(t, expected, New(3, 4))
}

func TestNewFromVec(t *testing.T) {
	a := vector.VecN{2, -3}
	b := vector.VecN{5, 7}
	c := vector.VecN{-5, 13}

	expected := Matrix{
		{2, 5, -5},
		{-3, 7, 13},
	}

	assert.Equal(t, expected, NewFromVec(a, b, c))
}

func TestMult(t *testing.T) {
	u := Matrix{
		{1, 2, 3},
		{4, 5, 6},
	}

	v := Matrix{
		{7, 8},
		{9, 10},
		{11, 12},
	}

	actual, err := Mult(u, v)
	if err != nil {
		t.Error(err)
	}

	expected := Matrix{
		{58, 64},
		{139, 154},
	}

	assert.Equal(t, expected, actual)
}

func TestScalarMult(t *testing.T) {
	u := Matrix{
		{1, 2, 3},
		{4, 5, 6},
	}

	expected := Matrix{
		{5, 10, 15},
		{20, 25, 30},
	}

	assert.Equal(t, expected, ScalarMult(u, 5))
}

func TestAdd(t *testing.T) {
	u := Matrix{
		{1, 2, 3},
		{4, 5, 6},
	}

	v := Matrix{
		{9, 8, 7},
		{6, 5, 4},
	}

	actual, err := Add(u, v)
	if err != nil {
		t.Error(err)
	}

	expected := Matrix{
		{10, 10, 10},
		{10, 10, 10},
	}

	assert.Equal(t, expected, actual)
}

func TestAddVec(t *testing.T) {
	u := Matrix{
		{1, 2, 3},
		{4, 5, 6},
	}

	v := vector.VecN{1, 2}

	actual, err := AddVec(u, v)
	if err != nil {
		t.Error(err)
	}

	expected := Matrix{
		{2, 3, 4},
		{6, 7, 8},
	}

	assert.Equal(t, expected, actual)
}

func TestSub(t *testing.T) {
	u := Matrix{
		{1, 2, 3},
		{4, 5, 6},
	}

	actual, err := Sub(u, u)
	if err != nil {
		t.Error(err)
	}

	expected := Matrix{
		{0, 0, 0},
		{0, 0, 0},
	}

	assert.Equal(t, expected, actual)
}

func TestSubVec(t *testing.T) {
	u := Matrix{
		{1, 2, 3},
		{4, 5, 6},
	}

	v := vector.VecN{1, 2}

	actual, err := SubVec(u, v)
	if err != nil {
		t.Error(err)
	}

	expected := Matrix{
		{0, 1, 2},
		{2, 3, 4},
	}

	assert.Equal(t, expected, actual)
}
