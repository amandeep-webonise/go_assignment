package main

import "fmt"

func main() {
	/*For loop with condition*/
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	/*For loop with intitial/condition/increment*/
	for j := 4; j <= 10; j++ {
		fmt.Println(j)
	}

	/*For with break*/
	for k := 1; k <= 5; k++ {

		if k%2 == 0 {
			break
		}
		fmt.Println(k)
	}

	/*For with continue*/
	for l := 1; l <= 5; l++ {

		if l%2 == 0 {
			continue
		}
		fmt.Println(l)
	}
}
