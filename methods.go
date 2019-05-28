package main

import "fmt"

//Define struct
type rect struct {
	width, height int
}

//Receiver type of method
func (r *rect) area() int {
	return r.width * r.height
}
func main() {

	r := rect{width: 10, height: 5}
	fmt.Println("area: ", r.area())

}
