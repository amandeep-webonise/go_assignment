package main

import "fmt"

//Interface
type geometry interface {
	area() float64
	perim() float64
}

//struct
type rect struct {
	width, height float64
}

//struc
/*type circle struct {
	radius
}*/

//implement interface
func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}
func main() {

	r := rect{5, 9}
	measure(r)

}
