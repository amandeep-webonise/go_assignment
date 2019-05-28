package main

import "fmt"

const a int = 5

func main() {
	fmt.Println("The value of constant a define outside of main function is ", a)

	const s string = "'Constant of string type'"
	fmt.Println("The value of constant s is ", s)
}
