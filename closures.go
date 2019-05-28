package main

import "fmt"

func intSqe() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	callIntSeq := intSqe()

	fmt.Println(callIntSeq())
	fmt.Println(callIntSeq())
	fmt.Println(callIntSeq())
	fmt.Println(callIntSeq())

}
