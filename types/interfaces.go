package main

import (
	"fmt"
)

// basically, it is like abstract class. collections of method signatures 

type shape interface {
	area() float64
	perimeter() float64
}

//uses interface (acting as an abstract class object basically, if we compare to C++)
func calculateShapeArea (sh shape) float64 {
	return sh.area()
}

type rect struct {
	width, height float64
}

func (rct rect) area() float64{
	return rct.width*rct.height
}

func (rct rect) perimeter() float64 {
    return 2 * (rct.width + rct.height)
}