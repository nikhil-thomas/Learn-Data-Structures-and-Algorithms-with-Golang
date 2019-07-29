package main

import "fmt"

type IDrawShape interface {
	drawShape(x [5]float32, y [5]float32)
}

type DrawContour struct {
	x [5]float32
	y [5]float32
	shape IDrawShape
	factor int
}

type IContour interface {
	drawContour()
	resizeByFactor(factor int)
}

func (dc DrawContour) drawContour() {
	fmt.Println("Drawing Contour")
	dc.shape.drawShape(dc.x, dc.y)
}

func (dc DrawContour) resizeByFactor(factor int) {
	dc.factor = factor
}

type DrawShape struct {}

func (ds DrawShape) drawShape(x, y [5]float32) {
	fmt.Println("drawing shape")
}

// The has-a relationship is maintained by composition
// Abstraction has a reference of the implementation.

func main() {
	var x = [5]float32{1,2,3,4,5}
	var y = [5]float32{1,2,3,4,5}

	var contour IContour = DrawContour{x, y, DrawShape{}, 2}
	contour.drawContour()
	contour.resizeByFactor(10)
}
