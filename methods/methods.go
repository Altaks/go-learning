package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.height * r.width
}

func (r *rect) perimeter() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{
		width:  10,
		height: 5,
	}

	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perimeter())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perimeter())
}
