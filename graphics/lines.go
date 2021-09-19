package main

import (
	"math"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"github.com/Jacalz/linalg/matrix"
	"github.com/Jacalz/linalg/rn"
)

type LineDrawer struct {
	widget.BaseWidget
	lines []fyne.CanvasObject
}

func NewLineDrawer() *LineDrawer {
	// These are actually points, but use vectors for simplicity.
	p1 := rn.VecN{0, 0, 0}
	p2 := rn.VecN{1, 0, 0}
	p3 := rn.VecN{1, 1, 0}
	p4 := rn.VecN{0, 1, 0}

	M := matrix.NewFromVec(p1, p2, p3, p4)

	// Scale around the set x, y and z factors.
	// Also multiply by 500 so that we get something resonably sized.
	x, y, z := 1.5*500, 0.5*500, 1.0*500
	T := matrix.Matrix{
		{x, 0, 0},
		{0, y, 0},
		{0, 0, z},
	}
	M, _ = matrix.Mult(T, M)

	// Translate the vector in the direction of u.
	u := rn.VecN{-2.0, -1.0, 3.0}
	M, _ = matrix.AddVec(M, u)

	// Transpose -30 degrees around z-axis.
	a := -math.Pi / 6
	T = matrix.Matrix{
		{math.Cos(a), -math.Sin(a), 0},
		{math.Sin(a), math.Cos(a), 0},
		{0, 0, 1},
	}
	M, _ = matrix.Mult(T, M)

	line1 := &canvas.Line{Position1: fyne.NewPos(float32(M[0][0]), float32(M[1][0])), Position2: fyne.NewPos(float32(M[0][1]), float32(M[1][1])), StrokeColor: theme.ErrorColor(), StrokeWidth: 2}
	line2 := &canvas.Line{Position1: fyne.NewPos(float32(M[0][1]), float32(M[1][1])), Position2: fyne.NewPos(float32(M[0][2]), float32(M[1][2])), StrokeColor: theme.ErrorColor(), StrokeWidth: 2}
	line3 := &canvas.Line{Position1: fyne.NewPos(float32(M[0][2]), float32(M[1][2])), Position2: fyne.NewPos(float32(M[0][3]), float32(M[1][3])), StrokeColor: theme.ErrorColor(), StrokeWidth: 2}
	line4 := &canvas.Line{Position1: fyne.NewPos(float32(M[0][3]), float32(M[1][3])), Position2: fyne.NewPos(float32(M[0][0]), float32(M[1][0])), StrokeColor: theme.ErrorColor(), StrokeWidth: 2}

	return &LineDrawer{lines: []fyne.CanvasObject{line1, line2, line3, line4}}
}

func (l *LineDrawer) CreateRenderer() fyne.WidgetRenderer {
	l.ExtendBaseWidget(l)
	return &lineRenderer{lineDrawer: l}
}

type lineRenderer struct {
	lineDrawer *LineDrawer
}

func (lr *lineRenderer) Destroy() {
}

func (lr *lineRenderer) Layout(s fyne.Size) {
}

func (lr *lineRenderer) MinSize() fyne.Size {
	return fyne.NewSize(theme.IconInlineSize(), theme.IconInlineSize())
}

func (lr *lineRenderer) Objects() []fyne.CanvasObject {
	return lr.lineDrawer.lines
}

func (lr *lineRenderer) Refresh() {
	canvas.Refresh(lr.lineDrawer)
}

func NewLineFromVecs(a, b rn.VecN) *canvas.Line {
	return nil
}
