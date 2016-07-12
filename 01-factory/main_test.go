package main_test

import (
  "testing"
  "reflect"
  "github.com/huytd/go-design-patterns/01-factory"
)

func TestShapeFactoryReturnCircle(t *testing.T) {
  factory := main.ShapeFactory{ }
  circle := factory.GetShape("CIRCLE")
  if reflect.TypeOf(circle).String() != "*main.Circle" {
    t.Fail()
  }
}

func TestShapeFactoryReturnRectangle(t *testing.T) {
  factory := main.ShapeFactory{ }
  rectangle := factory.GetShape("RECTANGLE")
  if reflect.TypeOf(rectangle).String() != "*main.Rectangle" {
    t.Fail()
  }
}

func TestShapeFactoryReturnSquare(t *testing.T) {
  factory := main.ShapeFactory{ }
  square := factory.GetShape("SQUARE")
  if reflect.TypeOf(square).String() != "*main.Square" {
    t.Fail()
  }
}

func TestCircleImplementsShape(t *testing.T) {
  shapeInterface := reflect.TypeOf((*main.Shape)(nil)).Elem()
  circle := main.Circle{}
  if !reflect.TypeOf(&circle).Implements(shapeInterface) {
    t.Fail()
  }
}

func TestRectangleImplementsShape(t *testing.T) {
  shapeInterface := reflect.TypeOf((*main.Shape)(nil)).Elem()
  rectangle := main.Rectangle{}
  if !reflect.TypeOf(&rectangle).Implements(shapeInterface) {
    t.Fail()
  }
}

func TestSquareImplementsShape(t *testing.T) {
  shapeInterface := reflect.TypeOf((*main.Shape)(nil)).Elem()
  square := main.Square{}
  if !reflect.TypeOf(&square).Implements(shapeInterface) {
    t.Fail()
  }
}
