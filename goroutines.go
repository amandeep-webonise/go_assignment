package main

import "fmt"

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("Normal Function Call")

	go f("Routine Function Call")

	go func(msg string) {
		fmt.Println("Message:", msg)
	}("call function in go fashion")

	fmt.Scanln()
	fmt.Println("done")
}
