package main

import "fmt"

func main() {

	first()
	//call function using defer to make execution last of main
	defer second()
	third()

}

func first() {
	fmt.Println("First function call")
}
func second() {
	fmt.Println("Second function call")
}
func third() {
	fmt.Println("Third function call")
}
