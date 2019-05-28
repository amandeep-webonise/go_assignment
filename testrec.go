package main

import (
	"fmt"
	//"os"
)

func recursive(num int, div int) int {
	if div == 1 {
		return 1
	}
	switch num % div {
	case 0:
		return recursive(num, div-1)
	default:
		return -1
	}
	return num
}

func main() {
	for i := 120; ; i += 20 {
		if recursive(i, 19) == 1 {
			fmt.Printf("finished with %v\n", i)
			break
		}
	}
}
