package main

import "fmt"

func add(a int, b int) int {
	return a + b
}
func main() {

	sum := add(1, 2)
	fmt.Printf("Result sum of two number using function %d\n", sum)

}
