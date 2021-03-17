package matrix

package main

import "testing"

func BenchmarkMult(b *testing.B) {
	a := [][]float64{
		{1, +0, 3},
		{2, +2, 8},
		{3, -2, 6},
	}

	for i := 0; i < b.N; i++ {
		Mult(a, a)
	}
}
