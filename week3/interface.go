package main

import "fmt"

func main() {
	passInterface(&circle{
		radius: 10,
	})
	passInterfaceRectangle(&rectangle{
		length: 20,
		width:  10,
	})
}

type shape interface {
	area() float64
	perimeter() float64
}

type circle struct {
	radius float64
}

type rectangle struct {
	length, width float64
}

func (c *circle) area() float64 {
	return 22 / 7 * (c.radius * c.radius)
}

func (c *circle) perimeter() float64 {
	return 2 * 22 / 7 * c.radius
}

func (r *rectangle) area() float64 {
	return r.length * r.width
}

func (r *rectangle) perimeter() float64 {
	return 2 * (r.length + r.width)
}

func passInterface(sh shape) {
	fmt.Println("Area is circle: ", sh.area())
	fmt.Println("perimeter is circle: ", sh.perimeter())

}
func passInterfaceRectangle(sh shape) {
	fmt.Println("Area is rectangle: ", sh.area())
	fmt.Println("perimeter is rectangle: ", sh.perimeter())

}
