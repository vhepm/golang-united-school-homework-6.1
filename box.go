package golang_united_school_homework

import (
	"errors"
	"fmt"
)

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if len(b.shapes) >= b.shapesCapacity {
		return errors.New("out of capacity")
	}
	b.shapes = append(b.shapes, shape)

	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if len(b.shapes) <= i {
		return nil, errors.New("index out of range")
	}
	value := b.shapes[i]

	if value == nil {
		return nil, errors.New("shape doesn't exist")
	}

	return value, nil
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	closure := func(index int) []Shape {
		return append(b.shapes[:index], b.shapes[index+1:]...)
	}

	it, err := b.GetByIndex(i)
	if err != nil {
		return nil, err
	}

	closure(i)

	return it, nil
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	it, err := b.GetByIndex(i)
	if err != nil {
		return nil, err
	}

	b.shapes[i] = shape

	return it, nil
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var sum float64

	for _, shape := range b.shapes {
		sum += shape.CalcPerimeter()
	}

	return sum
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var sum float64

	for _, shape := range b.shapes {
		sum += shape.CalcArea()
	}

	return sum

}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	called := false

	closure := func(index int) []Shape {
		called = true
		return append(b.shapes[:index], b.shapes[index+1:]...)
	}

	for index, shape := range b.shapes {
		if fmt.Sprintf("%T", shape) == "Circle" {
			closure(index)
		}
	}

	if !called {
		return errors.New("no circles")
	}
	return nil
}
