package main

import "fmt"

func main() {
	//Slice declare
	s := make([]string, 3)
	fmt.Println("This is format", s)

	//Assign value to slice
	s[0] = "first"
	s[1] = "second"
	s[2] = "third"

	fmt.Println("This is array with element ", s)

	//Append element to slice
	s = append(s, "four", "five", "six")
	fmt.Println("Append result to slice", s)

	//Copy a slice
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("Result of copy", c)

	//Get slice means a part from array
	l := s[2:5]
	fmt.Println("Get slice of s[2:5] ", l)

	//Other way to get value from slice
	l = s[:5]
	fmt.Println("Get value of s[:5] ", l)

	//other way to get value from slice
	l = s[2:]
	fmt.Println("Get value of s[2:]", l)

	//one way to creating slice
	n := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("New way to declare/define slice", n)

}
