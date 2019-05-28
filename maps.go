package main

import "fmt"

func main() {

	//Declare/define map
	m := make(map[string]int)
	m["first"] = 1
	m["two"] = 2
	m["three"] = 3
	fmt.Println("Print map ", m)

	//length of map
	l := len(m)
	fmt.Println("Print the length of map", l)

	//delete value of map
	delete(m, "two")
	fmt.Println("map after delete a value", m)

	//check map key exits or not
	_, p := m["first"]
	fmt.Println("Get map values", p)

	//other way of define map
	new_map := map[string]int{"eleven": 11, "tweleve": 12}
	fmt.Println("Get new map", new_map)

}
