package main

import ("fmt"
		"math"
)

type Shape interface {
	area() float64
}

type Circle struct {
	x,y,radius float64
}

type Rect struct {
	width, height float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rect) area() float64 {
	return r.width * r.height
}

func getArea(s Shape) float64{
	return s.area()
}


func main() {
	circle := Circle{0,0,5}
	rectangle := Rect{10,5}

	fmt.Println(getArea(circle))
	fmt.Println(getArea(rectangle))

}