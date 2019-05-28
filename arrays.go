package main

import "fmt"

func main() {
	/*Declaring Array*/
	var a [7]string

	fmt.Println("Trying to print array a", a)
	// Assign value to array index 4
	a[4] = "Fourth"
	//Print all the value of array
	fmt.Println("Getting value of array a", a)
	//Get particular index of array
	fmt.Printf("Getting Fourth index value of a %s\n", a[4])
	//Print length of array
	fmt.Printf("Getting length of a %d\n", len(a))

	//Declaring and intialize a array
	b := [3]int{3, 4, 5}

	//iterating on array index
	for i := 0; i <= len(b)-1; i++ {
		fmt.Println(b[i])
	}

	//Multi-dimension array
	var c [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			c[i][j] = i + j
		}
	}
	fmt.Println(c)
}
