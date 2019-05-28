package main

import (
	"errors"
	"fmt"
)

//By convention error is last agrument
func f1(arg int) (int, error) {
	if arg == 42 {
		//error.New constructs a basic error value with given error message
		return -1, errors.New("cannot work with 42")
	}
	//A nil value in the error position indicate that there is no error
	return arg + 3, nil
}
func main() {

	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}

}
