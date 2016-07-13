package main

import "fmt"

// Shape interface
type Shape interface {
	Draw()
}

// Rectangle class implementing Shape interface
type Rectangle struct{}

func (r *Rectangle) Draw() {
	fmt.Print("A rect is drawing\n")
}

// Circle class implementing Shape interface
type Circle struct{}

func (c *Circle) Draw() {
	fmt.Print("A circle is drawing\n")
}

// Square class implementing Shape interface
type Square struct{}

func (s *Square) Draw() {
	fmt.Print("A square is drawing\n")
}

// ShapeFactory is a factory to generate class based on given name
type ShapeFactory struct{}

// ShapeType Nay thi comment
type ShapeType int

const (
	CIRCLE ShapeType = iota
	RECTANGLE
	SQUARE
)

func (sf *ShapeFactory) GetShape(shapeType ShapeType) Shape {
	if shapeType == CIRCLE {
		return &Circle{}
	} else if shapeType == RECTANGLE {
		return &Rectangle{}
	} else if shapeType == SQUARE {
		return &Square{}
	} else {
		return nil
	}
}

// Demo program
func main() {
	factory := ShapeFactory{}

	circle := factory.GetShape(CIRCLE)
	circle.Draw()

	rectangle := factory.GetShape(RECTANGLE)
	rectangle.Draw()

	square := factory.GetShape(SQUARE)
	square.Draw()
}
