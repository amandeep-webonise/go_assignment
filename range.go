package main

import "fmt"

func main() {

	//Range on array
	l := []int{1, 2, 3, 4, 5}
	sum := 0

	for _, num := range l {
		sum = sum + num
	}

	fmt.Println("Print value num ", sum)

	//Range on map
	m := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}

	for key, val := range m {
		fmt.Printf("%s ---> %d\n", key, val)
	}

	//iterate over only key
	for key := range m {
		fmt.Printf("Key only %s\n", key)
	}

	//Range on string

	for i, c := range "AMANDEEP" {
		fmt.Printf("%d -----> %d\n", i, c)
	}
}
