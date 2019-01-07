package main

import (
      "fmt"

      "./shapes"
)

func main() {
	c := new(shapes.circle)
	c.setRadius(3)

	r := new(shapes.Rect)
	r.setLength(5)
	r.setBreadth(4)

	s := new(shapes.square) 
	s.setLength(6)

	fmt.Println("Area of the circle:", c.area())
	fmt.Println("Perimeter of the circle:", c.Perimeter())
	fmt.Println("Area of the Rectangle:", r.area())
	fmt.Println("Perimeter of the Rectangle:", r.Perimeter())
	fmt.Println("Area of the Square:", s.area())
	fmt.Println("Perimeter of the Square:", s.Perimeter())
}