package main

import "fmt"

func add_multiply(a int, b int) (int, int) {
	sum := a + b
	multiply := a * b
	return sum, multiply
}
func main() {
	sum, multi := add_multiply(5, 7)
	fmt.Printf("result of return multiple statement sum -->%d, multiplication -->%d\n", sum, multi)
}
