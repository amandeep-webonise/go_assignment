package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{"john", 23})

	//other way of call struct
	fmt.Println(person{"mack", 25})

	//assign one parameter to val
	fmt.Println(&person{"jambo", 55})

	//Acess struct with dot operator
	s := person{"mark", 29}
	fmt.Println(s.name)

	//Struct are mutable
	s.age = 90
	fmt.Println(s.age)

}
